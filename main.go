package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"

	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

// User implements the registration.User interface
type User struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}

func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

func main() {
	// TODO(jdt): use a flag or something not hardcoded
	key, err := os.ReadFile("./key")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := parseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		Email: "",
		Key:   privateKey,
	}

	client, err := lego.NewClient(lego.NewConfig(&user))
	if err != nil {
		log.Fatal(err)
	}

	reg, err := client.Registration.ResolveAccountByKey()
	if err != nil {
		log.Fatal(err)
	}
	user.Registration = reg

	log.Printf("%+v\n", reg)
}

func parseRSAPrivateKeyFromPEM(privPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}
