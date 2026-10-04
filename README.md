# *.localhost.direct - Wildcard Publicly Signed SSL Certificate with Subdomain Support
> [!NOTE]  
> New DNS Discussion is currently ongoing: https://github.com/Upinel/localhost.direct/issues/21  

One someday morning, I found myself tackling the usual trifecta of local development issues:

1. Using Fully Qualified Domain Names (FQDNs) in local testing environments
2. Dealing with SSL certificates in a local setting, where self-signed certificates are a nuisance
3. Implementing sub-domains in local development environments

To address these, I registered the domain localhost.direct and obtained a wildcard SSL certificate. I configured localhost.direct and *.localhost.direct to point to 127.0.0.1. Now, happy coding!

Eureka! I realized that I could share the private key and SSL certificate with developers globally at no extra cost. Thus, the *.localhost.direct project was born.

A project portal is available at https://get.localhost.direct/ for developers to download the most up-to-date SSL certificate bundle. This becomes the sole reserved subdomain. Updates to the wildcard SSL certificate will be posted, and your feedback is greatly appreciated.

Cheers!

## Important Informations
> [!IMPORTANT]
> non-SSL (HTTP): Running normally  
> SSL (HTTPS): Suggest use Private CA CertBundle and Trust in local environment. Due to the possible key leak causing cert revoke again https://github.com/Upinel/localhost.direct/issues/18, we are currently providing two-tier Cert Bundle.

> [!TIP]
> Addionally, *In fact, highly recommanded* you can self-sign your own *.localhost.direct certificate and trust the certificate in your organisation, it can make sure public CA revoke does nothing to your developing enviroment, and you can still enjoy public supported sub-domain development testing. (See Download - Section D)

> [!WARNING]
> Never Put the .key file in any public accessible place INCLUDING GITHUB projects. If founded, the cert will revoke. I don't want to require user registration in the future. issue: https://github.com/Upinel/localhost.direct/issues/18
> The Key files MUST always be password protected in a zip file. The cert bundle is zipped and password protected with a reason. CA is always keep scanning the internet to see if any key file leaked. Thanks.

## EULA
*To better comply with the guidelines, we’ve made some changes to our user agreement, and they’ll be effective immediately:*  
**By using this service,** you (localhost.direct developer) and LHD (localhost.direct) have agreed that you will act as LHD’s developer. However, LHD will never know what you’ve developed and will never claim ownership or copyright to your work. At the same time, LHD will not pay you any incentive or compensation and will not be responsible for any expenses or cost or damage incurred during your development.  
With this developer agreement, LHD shall point their developing environment domain (localhost.direct and *.localhost.direct) to your development server IP address (in this case, 127.0.0.1). LHD will also grant you the right to use LHD’s SSL certificate for internal development purposes only. Which will be provided to you separately.  
It is important to note that the Certificate Bundle(s) remains LHD's properties and you must never leak, share, or sublet this Certificate Bundle with any other parties. If you have more than one developer, they will all be under the same developer agreement with LHD.

## Usage:  
### For non-SSL user  
localhost.direct works immediately without configuration, functioning just like the traditional localhost, with added support for subdomain.localhost.direct.

### For HTTPS (SSL) in your local development environment  
Download the cert bundle from [Download](#download) below, then configure your web server:

#### nginx
```nginx
server {
    listen 443 ssl http2;
    server_name mysite.localhost.direct;

    ssl_certificate     /path/to/localhost.direct.crt;
    ssl_certificate_key /path/to/localhost.direct.key;

    location / {
        proxy_pass http://127.0.0.1:3000;
    }
}
```

#### Caddy (auto-HTTPS alternative — no cert files needed)
Caddy can use the localhost.direct cert directly:
```
mysite.localhost.direct {
    tls /path/to/localhost.direct.crt /path/to/localhost.direct.key
    reverse_proxy localhost:3000
}
```

#### Apache
```apache
<VirtualHost *:443>
    ServerName mysite.localhost.direct
    SSLEngine on
    SSLCertificateFile /path/to/localhost.direct.crt
    SSLCertificateKeyFile /path/to/localhost.direct.key
    ProxyPass / http://127.0.0.1:3000/
</VirtualHost>
```

#### Node.js (Express)
```javascript
const https = require('https');
const fs = require('fs');
const express = require('express');

const app = express();
https.createServer({
    key: fs.readFileSync('/path/to/localhost.direct.key'),
    cert: fs.readFileSync('/path/to/localhost.direct.crt')
}, app).listen(443);
```

#### Go (net/http)
Use the [HTTPS forwarding script](https://github.com/Upinel/localhost.direct/pull/16) for zero-config HTTPS→HTTP proxying.

## Known Issues

### ISP / DNS blocking of 127.0.0.1 resolution (#15)
Some ISPs and DNS resolvers (OpenDNS, corporate networks) block DNS responses that resolve to `127.0.0.1` as a DNS rebinding protection measure. If `localhost.direct` doesn't resolve for you:

1. **Use /etc/hosts** (simplest fix):
   ```
   127.0.0.1 localhost.direct
   127.0.0.1 mysite.localhost.direct
   ```
2. **Switch DNS** to 8.8.8.8 (Google) or 1.1.1.1 (Cloudflare)
3. **Run a local DNS resolver** (dnsmasq, unbound)

### Certificate revocation (#18)
The public CA certificate may be revoked if the private key is found in public repositories. For maximum stability, use the **Self-Signed Certificate** (Option A below) — it's immune to CA revocation and lasts 10 years.

### Temporarily forward HTTPS requests to your local HTTP service
```
go get ./...
go build
# Forward requests from https://<subdomain>.localhost.direct:3000 to http://192.168.9.113:3000
./localhost_direct -p 3000 -rh 192.168.9.113 -rp 3000
```

## Limitation:
**get.localhost.direct** is reserved and it is the only subdomain that you cannot use.

## Download
We recommend two approaches. **Option A (Self-Signed) is the stable choice.** Option B (Public CA) is a convenience option that may be revoked if keys leak.

### <ins>A. Self-Signed Certificate</ins> ⭐ Recommended
**Immune to CA revocation. 10-year validity. The most stable choice.**

This pre-generated self-signed certificate is the most reliable option. Download, install, and trust it once in your OS/browser. Since it's trusted locally, it's immune to CA revocations and lasts 10 years.

Download: [https://aka.re/localhost-ss](https://aka.re/localhost-ss)  
Password: **localhost**

**Trust the certificate:**
- **macOS:** Double-click the `.crt` file → Keychain Access → drag to "System" → double-click → Trust → "Always Trust"
- **Windows:** Double-click `.crt` → "Install Certificate" → "Local Machine" → "Trusted Root Certification Authorities"
- **Linux (system-wide):** `sudo cp localhost.direct.crt /usr/local/share/ca-certificates/ && sudo update-ca-certificates`
- **Firefox:** Settings → Privacy & Security → Certificates → View Certificates → Import
- **Node.js:** `export NODE_EXTRA_CA_CERTS=/path/to/localhost.direct.crt`

### <ins>B. Public CA Certificate</ins> ⚠️ May be revoked
> [!IMPORTANT]
> The public CA certificate may be revoked if any user leaks the key publicly. It has lower maintenance priority and relies on user reports for re-issuance. See [#18](https://github.com/Upinel/localhost.direct/issues/18).

Download: [https://aka.re/localhost](https://aka.re/localhost)  
Password: **IWillNotPutKeyFileInPublicAccessiblePlace.X1YKK**

### <ins>C. Cert Bundle by Request</ins> (Free — minimal registration)
Email get@localhost.direct with your Email, GitHub ID, and Project Name. Limited distribution reduces (but doesn't eliminate) revocation risk. Processing may take time.

### <ins>D. Cert Bundle by Sponsorship</ins>
Monthly $5+ or one-time $40+ sponsors receive a 1-year dedicated Cert Bundle. Email get@localhost.direct with your sponsorship ID. [Sponsor here](https://github.com/sponsors/Upinel).

## Last update Log. 
- 2026-Sep-26 Backend system update, Cert by Request and Cert by Sponsorship will receive new certificate every 60 days.
- 2024-Dec-01 Suggest User to use and trusting Private CA certificate
- 2024-Nov-19 Reissuing the General Certificate Bundle.
- 2024-Nov-11 User keep leaking key, new policy need to apply before new portal unfortunately.
- 2024-Nov-01 Short Term Cert Issue, Expire 30 Jan 2025
- 2024-Apr-20 SSL Intermediate Chain update with the help of @mundry, Expire keep 15 May 2025
- 2024-Apr-17 SSL Renewal, Expire 15 May 2025
- 2023-Jun-15 SSL Intermediate Chain update, Expire 31 Mar 2024
- 2023-Apr-01 SSL Renewal, Expire 31 Mar 2024
- 2022-Aug-29 Reupload the SSL file of 2022-Mar-29 (Expire 30 Apr 23)
- 2022-Mar-29 SSL Renewal  
- 2021-Mar-02 SSL Renewal  
- 2020-Feb-26 SSL Renewal  
- 2019-Feb-24 SSL project 

## Credit: 
This project is self-funded and shared freely with the community. We respect your privacy; your usage of localhost.direct is anonymous to us.  
Giving this project a star fuels our commitment to maintaining and improving it.

*Donations* are welcomed at [Github Sponsor](https://github.com/sponsors/Upinel) or [paypal.me/Upinel](https://paypal.me/Upinel) and are deeply appreciated. 

## A heartfelt thank you to the following sponsors :)
- SkyArk Inc (UK)
- Jackson Peak LLC
- the Lancelot Limited
- Peter Jong  
- cagnulein
- Klijn Engineering
- Saks Industries LLC

I love you all <3
