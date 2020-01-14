package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	text := []byte("My name is Assakra Radhouen")
	key := []byte("the-key-has-to-be-32-bytes-long!")

	ciphertext, err := encrypt(text, key)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Printf("%s => %x\n", text, ciphertext)

	plaintext, err := decrypt(ciphertext, key)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Printf("%x => %s\n", ciphertext, plaintext)
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
//Calling the above function aes.NewCipher (whose []byte key parameter must be 16, 24 or 32, corresponding to the AES-128, AES-192 or AES-256 algorithms, respectively), returns a cipher.Block Interface that implements three functions:
type Block interface {
	// BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
	// Dst and src may point at the same memory.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src may point at the same memory.
	Decrypt(dst, src []byte)
}
