> [!IMPORTANT]
> non-SSL (HTTP): Running normally  
> SSL (HTTPS): Running normally. Due to the possible key leak causing cert revoke again https://github.com/Upinel/localhost.direct/issues/18, we are currently providing two-tier Cert Bundle.

> [!TIP]
> Addionally, *In fact, highly recommanded* you can self-sign your own *.localhost.direct certificate and trust the certificate in your organisation, it can make sure public CA revoke does nothing to your developing enviroment, and you can still enjoy public supported sub-domain development testing.

> [!WARNING]
> Never Put the .key file in any public accessible place. If founded, the cert will revoke. I don't want to require user registration in the future. issue: https://github.com/Upinel/localhost.direct/issues/18

# *.localhost.direct - Wildcard Publicly Signed SSL Certificate with Subdomain Support
One someday morning, I found myself tackling the usual trifecta of local development issues:

1. Using Fully Qualified Domain Names (FQDNs) in local testing environments
2. Dealing with SSL certificates in a local setting, where self-signed certificates are a nuisance
3. Implementing sub-domains in local development environments

To address these, I registered the domain localhost.direct and obtained a wildcard SSL certificate. I configured localhost.direct and *.localhost.direct to point to 127.0.0.1. Now, happy coding!

Eureka! I realized that I could share the private key and SSL certificate with developers globally at no extra cost. Thus, the *.localhost.direct project was born.

A project portal is available at https://get.localhost.direct/ for developers to download the most up-to-date SSL certificate bundle. This becomes the sole reserved subdomain. Updates to the wildcard SSL certificate will be posted, and your feedback is greatly appreciated.

Cheers!

## EULA
Thanks to [@BenBE](https://github.com/BenBE) for helping us summarize the [CA guideline](https://github.com/Upinel/localhost.direct/issues/19). To better comply with the guidelines, we’ve made some changes to our user agreement, and they’ll be effective immediately:
> **By using this service,** you (localhost.direct developer) and LHD (localhost.direct) have agreed that you will act as LHD’s developer. However, LHD will never know what you’ve developed and will never claim ownership or copyright to your work. At the same time, LHD will not pay you any incentive or compensation and will not be responsible for any expenses or cost or damage incurred during your development.  
> With this developer agreement, LHD shall point their developing environment domain (localhost.direct and *.localhost.direct) to your development server IP address (in this case, 127.0.0.1). LHD will also grant you the right to use LHD’s SSL certificate for internal development purposes only. Which will be provided to you separately.  
> It is important to note that the Certificate Bundle(s) remains LHD's properties and you must never leak, share, or sublet this Certificate Bundle with any other parties. If you have more than one developer, they will all be under the same developer agreement with LHD.
  
**Important Disclaimer**: [@BenBE](https://github.com/BenBE) is NOT endorsing this project in any way, express or implied. FYI.

## Usage:  
### For non-SSL user

localhost.direct works immediately without configuration, functioning just like the traditional localhost, with added support for subdomain.localhost.direct.

###  For user would like to use HTTPS (SSL) in their localhost development environment

Download or clone the .key and .crt files, then deploy them to your local web server to set up an SSL-enabled local development environment.

## Limitation:
**get.localhost.direct** is reserved and it is the only subdomain that you cannot use.

## Download

We now have two-tier cert bundle, for user want to have fully anonymous, please use General Cert Bundle.
We also provide Cert Bundle by Request and Cert Bundle by Sponsorship, you can request it by email.

### General Cert Bundle (fully Anonymous)
> [!IMPORTANT]
> > Using General Cert Bundle might have the risk of cert revocation again and again if any user misuses it. The General Cert might also have very low priority of maintenance and require user reports to re-issue. https://github.com/Upinel/localhost.direct/issues/18        
Download: https://aka.re/localhost
**Password for General Cert Bundle file:**  
**IWillNotPutKeyFileInPublicAccessiblePlace.X1YKK**

### Cert Bundle by Request (Free)
You may send a request to get@localhost.direct to obtain a certificate bundle that is only shared with limited users if you want to reduce the risk of being revoked due to other users’ bad behaviour (not 100%).  
The email must include the following information:  
```
Your Email: (Email must not be temporary email)
Your GitHub ID:
Your Project Name: 
```
**The Cert Bundle by Request might take sometime to process**

### Cert Bundle by Sponsorship
To thank sponsors, sponsors with $5/monthly or >$40 one-time sponsor can request a 1-year version of the cert bundle. You may send a request to get@localhost.direct to obtain the yearly long certificate with your sponsorship github ID / Paypal ID. All sponsors will help us develop our own auto-issuing portal as soon as possible.

### 
> [!IMPORTANT]
> Due to the Issue: https://github.com/Upinel/localhost.direct/issues/18, meanwhile you might need to send email to get@localhost.direct to get the cert bundle. Donater will have higher priority. Please support our new portal development.

## Last update Log. 
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

*Donations* are welcomed at [paypal.me/Upinel](https://paypal.me/Upinel) and are deeply appreciated. 

## A heartfelt thank you to the following sponsors via [paypal.me/Upinel](https://paypal.me/Upinel) :)
- SkyArk Inc (UK)
- Jackson Peak LLC
- the Lancelot Limited
- Peter Jong  
- cagnulein
- Klijn Engineering

I love you all <3
