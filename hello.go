package main

import (
	"golang.org/x/crypto/nacl/secretbox"
	"crypto/rand"
	"fmt"
)

func main() {
	var key [32]byte
	var nonce1 [24]byte
	var nonce2 [24]byte
	rand.Reader.Read(key[:])
	rand.Reader.Read(nonce1[:])
	rand.Reader.Read(nonce2[:])

	message1 := []byte("FirstPart ")
	message2 := []byte("SecondPart")

	var box1, box2, opened []byte
	box1 = secretbox.Seal([]byte{}, message1, &nonce1, &key)
	box2 = secretbox.Seal([]byte{}, message2, &nonce2, &key)

	opened, _ = secretbox.Open(opened, box1, &nonce1, &key)
	fmt.Println(string(opened))

	opened, _ = secretbox.Open(opened, box2, &nonce2, &key)
	fmt.Println(string(opened))
}
