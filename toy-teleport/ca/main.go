package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

const keypath = "./ca_key.pem"

func main() {

	priv, err := loadAndSaveCA()

	if err != nil {
		panic(err)
	}

	caSigner, err := ssh.NewSignerFromKey(priv)

	if err != nil {
		panic(err)
	}

	if err := mintHostCert(caSigner); err != nil {
		panic(err)
	}

}

func generateAndSaveCA() (ed25519.PrivateKey, error) {

	_, priv, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return nil, err
	}

	der, err := x509.MarshalPKCS8PrivateKey(priv)

	if err != nil {
		return nil, err
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})

	if err := os.WriteFile(keypath, pemBytes, 0600); err != nil {
		return nil, err
	}
	return priv, nil

}

func loadAndSaveCA() (ed25519.PrivateKey, error) {
	pemBytes, err := os.ReadFile(keypath)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Fprintln(os.Stderr, "no key found, generating a new one")
		return generateAndSaveCA()
	}

	if err != nil {
		return nil, err
	}

	fmt.Printf("loading exisitng key from disk")

	block, _ := pem.Decode(pemBytes)

	if block == nil {
		return nil, errors.New("file is not a valid PEM")

	}

	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	priv, ok := parsed.(ed25519.PrivateKey)

	if !ok {
		return nil, errors.New("key is not ed25519")

		// fmt.Printf()
	}

	return priv, nil

}

func mintHostCert(caSigner ssh.Signer) error {

	pub, priv, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {
		return err
	}

	if err := writeKEYPem(priv); err != nil {
		return err
	}

	sshPub, err := ssh.NewPublicKey(pub)

	if err != nil {
		return err
	}

	now := time.Now()
	hostCert := &ssh.Certificate{
		Key:             sshPub,
		CertType:        ssh.HostCert,
		KeyId:           "toy-server",
		ValidPrincipals: []string{"localhost", "127.0.0.1"},
		ValidAfter:      uint64(now.Add(-1 * time.Minute).Unix()),
		ValidBefore:     uint64(now.Add(1 * time.Minute).Unix()),
	}

	if err := hostCert.SignCert(rand.Reader, caSigner); err != nil {
		return err
	}

	if err := os.WriteFile("host_cert.pub", ssh.MarshalAuthorizedKey(hostCert), 0644); err != nil {
		return err
	}

	// Inspect the fields (this is the Milestone 2 verify).
	fmt.Printf("signed host cert for principals: %v\n", hostCert.ValidPrincipals)
	fmt.Printf("valid after:  %s\n", time.Unix(int64(hostCert.ValidAfter), 0))
	fmt.Printf("valid before: %s\n", time.Unix(int64(hostCert.ValidBefore), 0))

	return nil
}

func writeKEYPem(priv ed25519.PrivateKey) error {
	der, err := x509.MarshalPKCS8PrivateKey(priv)

	if err != nil {
		return err
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})

	return os.WriteFile(keypath, pemBytes, 0600)

}
