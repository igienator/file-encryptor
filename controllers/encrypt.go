package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

func Encrypt(filename string, output string, keyfile string) error {
	plainFile, err := os.ReadFile(filename)
	if err != nil {
		return errors.Join(errors.New("can't open file: "), err)
	}

	key, err := os.ReadFile(keyfile)
	if err != nil {
		return errors.Join(errors.New("can't open key file: "), err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return errors.Join(errors.New("can't create block: "), err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return errors.Join(errors.New("can't create GCM: "), err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return errors.Join(errors.New("can't create nonce: "), err)
	}

	cipherText := gcm.Seal(nonce, nonce, plainFile, nil)

	err = os.WriteFile(output, cipherText, 0777)
	if err != nil {
		return errors.Join(errors.New("can't write output file: "), err)
	}
	return nil
}
