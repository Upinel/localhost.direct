> [!IMPORTANT]
> non-SSL (HTTP): Running normally  
> SSL (HTTPS): Seems revoked again, issue information : https://github.com/Upinel/localhost.direct/issues/18 , I think I will develop a portal for user signup and distrubute different key pair. During the development, please email me to receive updated certificate bundle, I will send it manaually so if the volume is high, it can take some timne as I have my day work.
> Addionally, *In fact, highly recommanded* you can self-sign your own localhost.direct certificate and trust the certificate in your organisation, it can make sure public CA revoke does nothing to your developing enviroment, and you can still enjoy public supported sub-domain development testing.

> [!WARNING]
> Never Put the .key file in any public accessible place. If founded, the cert will revoke. I don't want to require user registration in the future. issue: https://github.com/Upinel/localhost.direct/issues/18

# *.localhost.direct - Wildcard Publicly Signed SSL Certificate with Subdomain Support
One Sunday morning, I found myself tackling the usual trifecta of local development issues:

1. Using Fully Qualified Domain Names (FQDNs) in local testing environments
2. Dealing with SSL certificates in a local setting, where self-signed certificates are a nuisance
3. Implementing sub-domains in local development environments

To address these, I registered the domain localhost.direct and obtained a wildcard SSL certificate. I configured localhost.direct and *.localhost.direct to point to 127.0.0.1. Now, happy coding!

Eureka! I realized that I could share the private key and SSL certificate with developers globally at no extra cost. Thus, the *.localhost.direct project was born.

A project portal is available at https://get.localhost.direct/ for developers to download the most up-to-date SSL certificate bundle. This becomes the sole reserved subdomain. Updates to the wildcard SSL certificate will be posted, and your feedback is greatly appreciated.

Cheers!

## Usage:  
### For non-SSL user

localhost.direct works immediately without configuration, functioning just like the traditional localhost, with added support for subdomain.localhost.direct.


###  For user would like to use HTTPS (SSL) in their localhost development environment

Download or clone the .key and .crt files, then deploy them to your local web server to set up an SSL-enabled local development environment.

## Limitation:
**get.localhost.direct** is reserved and it is the only subdomain that you cannot use.

## Download
https://aka.re/localhost
> [!IMPORTANT]
> Due to the Issue: https://github.com/Upinel/localhost.direct/issues/18, meanwhile you might need to send email to get@localhost.direct to get the cert bundle. Donater will have higher priority. Please support our new portal development.

## Password for cert file:  
**IWillNotPutKeyFileInPublicAccessiblePlace**

## Last update Log. 
- 2024-Noc-11 User keep leaking key, new policy need to apply before new portal unfortunately.
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

*Donations* are welcomed at [paypal.me/Upinel](https://paypal.me/Upinel) and are deeply appreciated. 

## A heartfelt thank you to the following sponsors via [paypal.me/Upinel](https://paypal.me/Upinel) :)
- SkyArk Inc (UK)
- Jackson Peak LLC
- the Lancelot Limited
- Peter Jong  
- cagnulein
- Klijn Engineering

I love you all <3
