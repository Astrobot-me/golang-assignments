package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println(ssh.CertAlgoSKED25519v01)
}
