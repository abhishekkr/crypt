## crypt

This is supposed to manage daily secrets 'better'.

[latest release with binaries for v0.2.0](https://github.com/abhishekkr/crypt/releases/latest)

### Does

> for every task, your `passphrase` to be used is asked and never cached

* Generates a new secret.

* Create secrets to a locally managed file. It asks for `topic`, `key` and `value`.

* Read a previously made safe secret. It asks for `topic`.

* List all `topics` of previously saved secrets.


#### What are these constructs

* `topic` (has to be unique, otherwise overwrites) is any identity you want to file your secret detail under. Could be anything like name of site whose password you wanna save or bank name for which you wanna save account details.

* `key` is kind of secret this is. Could be 'password' if saving credentials for any website or `account-number` if a bank detail.

* `value` is the real secret to be kept safe.


#### Todo

* Update and Delete locally managed secrets.

* Create, read, update, delete and list secrets on popular online storages.

---

### Usage

* generate secret

```
± % crypt  ### or ± % crypt -axn gen
here to generate secret, new one
Oz2FL=A0Z=U3w6mIP+RYj7s9aQ17rVVP
```

* creating secret

```
± % crypt -axn create
Enter your passphrase when none is looking: look
here to create secret, carry on
Secret Topic/Domain (eg. gmail.com): gmail
Secret Key (eg. password, username): user
Secret Value (secret for user): me
```

* list secrets

```
± % crypt -axn list
Enter your passphrase when none is looking: look
here to list all topics in this secret, find yours
gmail
```

* reading secret

```
± % crypt -axn read
Enter your passphrase when none is looking: look
here to read secret, stay hidden
Secret Topic/Domain (eg. gmail): gmail
user : me
```

---

> primarly created it because to manage secrets I don't trust a 3rd Party util

---
