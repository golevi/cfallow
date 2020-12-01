# cfallow

## Usage

You must authenticate with a Cloudflare API Token.

To use an API Token, set the `CF_ALLOW` environment variable:

```bash
export CF_ALLOW=Abc123Xyz
```

Example of how it works. It creates an md5 hash of your computer's hostname,
that is your rule hash. We then get your IP address. Next, it looks to see
which accounts you have access to and loops through them, removing your
previous entry and adding a new entry with your latest IP address.

```bash
Starting...
Rule hash 11111c15372208abfb4265ebe0123456
Your IP 260f:5c5:5555:fff:0000:0000:0000:ffff
=========================================================
Client 1 Account
Deleted rule 11112222333344445555666677778888
IP Added...
=========================================================
Client 2 Account
Deleted rule a1112222333344445555666677778888
IP Added...
=========================================================
Client 3 Account
Deleted rule b1112222333344445555666677778888
IP Added...
=========================================================
Client 4 Account
Deleted rule c1112222333344445555666677778888
IP Added...
=========================================================
Client 5 Account
Deleted rule d1112222333344445555666677778888
IP Added...
=========================================================
Client 6 Account
Deleted rule e1112222333344445555666677778888
IP Added...
=========================================================
Done!
```

## API token permissions

Your token will need the following tokens.

### All accounts

* Account Filter Lists:Read
* Account Firewall Access Rules:Edit
* DNS Firewall:Edit
* Account Settings:Read

### All zones

* Firewall Services:Edit

### All users

* User Details:Read
