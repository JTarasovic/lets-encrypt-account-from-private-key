# lets-encrypt-account-from-private-key
Get Let's Encrypt account details with only a private key

I had access to a private key (from cert-manager secret) but not the account number.
Threw this together to get the missing details.

- https://letsencrypt.org/docs/account-id/
- https://tools.ietf.org/html/rfc8555#section-7.3

Uses [lego](github.com/go-acme/lego). Should work for other ACME protocol implementations but I haven't tried any others. 
