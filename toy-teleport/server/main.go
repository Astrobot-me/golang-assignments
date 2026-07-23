package main

// generated though claude
import (
	"bytes"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// 1. Load the CA's PUBLIC key. This is the entire basis of trust:
	//    "I will accept anyone whose cert was signed by this key."
	caPubBytes, err := os.ReadFile("host_cert.pub")
	if err != nil {
		panic(err)
	}
	caPub, _, _, _, err := ssh.ParseAuthorizedKey(caPubBytes)
	if err != nil {
		panic(err)
	}

	// 2. Build the signer the server presents to clients: host key + host cert.
	hostCertSigner, err := loadHostCertSigner()
	if err != nil {
		panic(err)
	}

	// 3. A CertChecker answers one question: "was this cert signed by an authority I trust?"
	checker := &ssh.CertChecker{
		IsUserAuthority: func(auth ssh.PublicKey) bool {
			// Compare the cert's signing key against OUR ca's public key.
			return bytes.Equal(auth.Marshal(), caPub.Marshal())
		},
	}

	config := &ssh.ServerConfig{
		// This fires for every client that tries public-key auth.
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			// (a) the presented key MUST be a certificate...
			cert, ok := key.(*ssh.Certificate)
			if !ok {
				return nil, fmt.Errorf("rejected %q: plain key, not a certificate", conn.User())
			}
			// (b) ...and the CertChecker must be happy (signed by our CA,
			//     user cert type, not expired, principal matches the username).
			if err := checker.CheckCert(conn.User(), cert); err != nil {
				return nil, fmt.Errorf("rejected %q: %w", conn.User(), err)
			}
			fmt.Printf("ACCEPTED user=%q certID=%q\n", conn.User(), cert.KeyId)
			return &ssh.Permissions{}, nil
		},
	}
	// The server presents the certificate (not the bare host key) to clients.
	config.AddHostKey(hostCertSigner)

	// 4. Listen and accept connections forever.
	ln, err := net.Listen("tcp", "127.0.0.1:2022")
	if err != nil {
		panic(err)
	}
	fmt.Println("listening on 127.0.0.1:2022 — Ctrl-C to stop")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go handleConn(conn, config) // one goroutine per connection
	}
}

// loadHostCertSigner reads host_key.pem + host_cert.pub and wraps them together.
func loadHostCertSigner() (ssh.Signer, error) {
	keyBytes, err := os.ReadFile("ca_key.pem")
	if err != nil {
		return nil, err
	}
	hostSigner, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	certBytes, err := os.ReadFile("host_cert.pub")
	if err != nil {
		return nil, err
	}
	pub, _, _, _, err := ssh.ParseAuthorizedKey(certBytes)
	if err != nil {
		return nil, err
	}
	cert, ok := pub.(*ssh.Certificate)
	if !ok {
		return nil, fmt.Errorf("host_cert.pub is not a certificate")
	}
	return ssh.NewCertSigner(cert, hostSigner)
}

func handleConn(nConn net.Conn, config *ssh.ServerConfig) {
	// The handshake runs auth. If auth fails, this returns an error — that's
	// the EXPECTED path when someone connects with a plain key.
	conn, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		fmt.Println("connection rejected:", err)
		nConn.Close()
		return
	}
	defer conn.Close()

	go ssh.DiscardRequests(reqs)
	for newChan := range chans {
		// We'll make sessions actually DO something in Milestone 4.
		// For now, accept the channel and close it so nothing hangs.
		if newChan.ChannelType() != "session" {
			newChan.Reject(ssh.UnknownChannelType, "only sessions supported")
			continue
		}
		ch, chReqs, err := newChan.Accept()
		if err != nil {
			continue
		}
		go ssh.DiscardRequests(chReqs)
		ch.Close()
	}
}
