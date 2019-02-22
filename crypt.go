package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	golcrypt "github.com/abhishekkr/gol/golcrypt"
	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
	golrandom "github.com/abhishekkr/gol/golrandom"
)

var (
	CyfrAxn    = flag.String("axn", "gen", "gen/list/create/read/update/delete")
	CyfrPath   = flag.String("path", ".cyfr.secrets", "path to secret store")
	Passphrase string
)

type Secret struct {
	Key   string            `json:"key"`
	Crypt golcrypt.AESBlock `json:"crypt"`
}

type SecretSafe struct {
	Secrets map[string]Secret `json:"secret"`
}

func main() {
	flag.Parse()
	if *CyfrAxn != "gen" {
		fmt.Print("Enter your passphrase when none is looking: ")
		fmt.Scanln(&Passphrase)
	}

	secretSafe, err := readCyfr()

	if err != nil && (*CyfrAxn != "create" && *CyfrAxn != "gen") {
		log.Println("can provide a custom path to cyfr-path with flag '-path <secret-file>'")
		log.Fatalf("can not perform %s on %s, secret file doesn't exist", *CyfrAxn, *CyfrPath)
	} else if err != nil && *CyfrAxn == "create" {
		secretSafe.Secrets = make(map[string]Secret)
	}

	if *CyfrAxn == "gen" {
		fmt.Println("here to generate secret, new one")
		fmt.Println(golrandom.Token(32))
	} else if *CyfrAxn == "create" {
		fmt.Println("here to create secret, carry on")
		secretSafe.create()
	} else if *CyfrAxn == "read" {
		fmt.Println("here to read secret, stay hidden")
		secretSafe.read()
	} else if *CyfrAxn == "list" {
		fmt.Println("here to list all topics in this secret, find yours")
		secretSafe.list()
	} else if *CyfrAxn == "update" {
		fmt.Println("[WIP] here to update secret, good job")
	} else if *CyfrAxn == "delete" {
		fmt.Println("[WIP] here to update secret, keep calm")
	}
}

func (secretSafe SecretSafe) create() {
	var topic, key string
	var crypt golcrypt.AESBlock
	fmt.Print("Secret Topic/Domain (eg. gmail.com): ")
	fmt.Scanln(&topic)
	fmt.Print("Secret Key (eg. password, username): ")
	fmt.Scanln(&key)
	fmt.Printf("Secret Value (secret for %s): ", key)
	fmt.Scanln(&crypt.DataBlob)
	secretKey := golrandom.Token(1024)
	crypt.Key = golcrypt.KeyForAES([]byte(secretKey))
	crypt.Encrypt()
	crypt.DataBlob = golcrypt.DataBlob([]byte{})

	secretSafe.Secrets[topic] = Secret{Key: key, Crypt: crypt}
	writeCyfr(secretSafe)
}

func (secretSafe SecretSafe) read() {
	var topic string
	fmt.Print("Secret Topic/Domain (eg. gmail): ")
	fmt.Scanln(&topic)

	secret := secretSafe.Secrets[topic]
	secret.Crypt.Decrypt()

	fmt.Println(secret.Key, ":", string(secret.Crypt.DataBlob))
}

func (secretSafe SecretSafe) list() {
	for topic, _ := range secretSafe.Secrets {
		fmt.Println(topic)
	}
}

func writeCyfr(secretSafe SecretSafe) error {
	secretsJson, err := json.Marshal(secretSafe)
	if err != nil {
		panic("failed to serialize new secrets data")
	}

	aesBlock := golcrypt.AESBlock{
		DataBlob: []byte(secretsJson),
		Key:      golcrypt.KeyForAES([]byte(Passphrase)),
		Cipher:   nil,
	}
	aesBlock.Encrypt()

	return ioutil.WriteFile(*CyfrPath, aesBlock.Cipher, 0644)
}

func readCyfr() (SecretSafe, error) {
	var secretSafe SecretSafe
	if !golfilesystem.PathExists(*CyfrPath) {
		golfilesystem.MkDir(path.Dir(*CyfrPath))
		return secretSafe, fmt.Errorf("%s doesn't exist to be read", *CyfrPath)
	}
	jsonCyfr, err := ioutil.ReadFile(*CyfrPath)
	if err != nil {
		return secretSafe, err
	}

	aesBlock := golcrypt.AESBlock{
		DataBlob: nil,
		Key:      golcrypt.KeyForAES([]byte(Passphrase)),
		Cipher:   []byte(jsonCyfr),
	}
	aesBlock.Decrypt()

	json.Unmarshal(aesBlock.DataBlob, &secretSafe)
	return secretSafe, nil
}
