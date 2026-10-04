package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/yeka/zip"
)

func main() {
	port := flag.Int("p", 443, "Local HTTPS server port")
	remoteHost := flag.String("rh", "127.0.0.1", "Remote HTTP server host")
	remotePort := flag.Int("rp", 0, "Remote HTTP server port")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nAccess the server using https://a.localhost.direct:<port>\n")
	}
	flag.Parse()

	if !isValidIP(*remoteHost) {
		fmt.Printf("Error: Invalid IP address for remote host: %s\n", *remoteHost)
		flag.Usage()
		os.Exit(1)
	}

	if *remotePort == 0 {
		*remotePort = *port
	}

	// Download and extract SSL certificate
	certFile, keyFile, err := downloadAndExtractCert()
	if err != nil {
		fmt.Printf("Error downloading certificate: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(certFile)
	defer os.Remove(keyFile)

	// Set up reverse proxy
	target, _ := url.Parse(fmt.Sprintf("http://%s:%d", *remoteHost, *remotePort))
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Configure TLS
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Printf("Error loading certificate: %v\n", err)
		os.Exit(1)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", *port),
		TLSConfig: tlsConfig,
		Handler:   proxy,
	}

	fmt.Printf("Starting HTTPS server on port %d, forwarding to %s:%d\n", *port, *remoteHost, *remotePort)
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func downloadAndExtractCert() (string, string, error) {
	resp, err := http.Get("https://aka.re/localhost")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return "", "", err
	}

	var certFile, keyFile string
	for _, file := range zipReader.File {
		if file.IsEncrypted() {
			file.SetPassword("localhost")
		}

		outFile, err := os.CreateTemp("", filepath.Base(file.Name))
		if err != nil {
			return "", "", err
		}
		defer outFile.Close()

		fileReader, err := file.Open()
		if err != nil {
			return "", "", err
		}
		defer fileReader.Close()

		_, err = io.Copy(outFile, fileReader)
		if err != nil {
			return "", "", err
		}

		if filepath.Ext(file.Name) == ".crt" {
			certFile = outFile.Name()
		} else if filepath.Ext(file.Name) == ".key" {
			keyFile = outFile.Name()
		}
	}

	return certFile, keyFile, nil
}
