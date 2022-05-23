# *.localhost.direct Wildcard public CA signed SSL cert with sub-domain support
This is a story in a Sunday morning, while I am solving a traditional pain of three questions:
1. using the FQDN domain in a local testing environment
2. using SSL in a local environment but self-signed cert can always annoying 
3. using sub-domain in a local development environment  
So I registered the domain name localhost.direct and applied for a wildcard SSL. Point localhost.direct and *.localhost.direct to 127.0.0.1. Happy coding.  
And soon I found that there is no extra cost if I share the private key and the SSL cert to all the developers in the world. As so I did this.

I set up a project website https://get.localhost.direct/ for the entry point to let developers download the latest SSL cert bundle. And it becomes the only subdomain (get.localhost.direct) reserved. 
I will keep updating the wildcard SSL cert on https://get.localhost.direct/ and suggestions are extremely welcome.

Cheers!

## Usage:  
### For non-SSL user

localhost.direct with work directly without any setting, this FQDN function exactly the same as your traditional localhost  but with subdomain.localhost.direct support natively.


###  For user would like to use HTTPS (SSL) in their localhost development environment

Simply download or clone the key and crt file and deploy to your local development web server, you will soon enjoy the SSL development environment!

## Limitation:
**get.localhost.direct** is reserved and it is the only subdomain that you cannot use.

## Download
https://aka.re/localhost

## Password for cert file:  
**localhost**

## Last update Log. 
- 2022-Mar-29 SSL Renewal  
- 2021-Mar-02 SSL Renewal  
- 2020-Feb-26 SSL Renewal  
- 2019-Feb-24 SSL project inital  

## Credit: 
This project is entirely our own spending and sharing to the world with no charge. We also will not know how you use the localhost.direct and do not able to collect any information from you.  
Staring this project will be our energy to keep operating and maintenance the programme.  
If you wish to donate us, please donate to paypal.me/Upinel, it will be really lovely.
