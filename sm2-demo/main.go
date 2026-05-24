package main

import (
	"crypto/rand"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"
)

func main() {
	// Generate a new key pair
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	// Plain text
	msg := []byte("Hello, SM2 encryption!")

	// Encrypt using the public key
	cipher, err := sm2.EncryptAsn1(&priv.PublicKey, msg, rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext (hex): %x\n", cipher)

	// Decrypt using the private key
	plain, err := sm2.DecryptAsn1(priv, cipher, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted message: %s\n", plain)
}
