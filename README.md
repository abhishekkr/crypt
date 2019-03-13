## crypt

[![Go Report Card](https://goreportcard.com/badge/github.com/abhishekkr/crypt)](https://goreportcard.com/report/github.com/abhishekkr/crypt)

This is supposed to manage daily secrets 'better'.

[latest release with binaries for v0.2.0](https://github.com/abhishekkr/crypt/releases/latest)

> primarly created it because to manage secrets I don't trust a 3rd Party util
>
> if you don't as well, read the [small-ish single file code here](./crypt.go) and build yourself

---

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

* for help `crypt -help`

* generate secret

```
± % crypt  ### or ± % crypt -axn gen
here to generate secret, new one
Oz2FL=A0Z=U3w6mIP+RYj7s9aQ17rVVP
```

> * all commands other than `gen`, need a secret store file
>
> * would assume secret store `.cyfr.secrets` file to be present at the directory of run
>
> * to provide a custom file, use switch `-path $FILEPATH`

* creating secret; can also update a secret (let's you overwrite with permission)

```
± % crypt -axn create
Enter your passphrase when none is looking: look
here to create secret, carry on
Secret Topic/Domain (eg. gmail.com): gmail
Secret Key (eg. password, username): user
Secret Value (secret for user): me

± %# crypt -axn create -path ~/.cyfr ## to store in custom file
```

* list secrets

```
± % crypt -axn list
Enter your passphrase when none is looking: look
here to list all topics in this secret, find yours
gmail

± %# crypt -axn list -path ~/.cyfr ## to list from custom file
```

* reading secret

```
± % crypt -axn read
Enter your passphrase when none is looking: look
here to read secret, stay hidden
Secret Topic/Domain (eg. gmail): gmail
user : me

± %# crypt -axn read -path ~/.cyfr ## to read from custom file
```

---
