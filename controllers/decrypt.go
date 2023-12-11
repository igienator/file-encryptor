package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"os"
)

func Decrypt(filename string, output string, keyfile string) error {
	encryptedFile, err := os.ReadFile(filename)
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

	nonce := encryptedFile[:gcm.NonceSize()]

	cipherText := encryptedFile[gcm.NonceSize():]

	plainFile, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return errors.Join(errors.New("can't decrypt file: "), err)
	}

	err = os.WriteFile(output, plainFile, 0777)
	if err != nil {
		return errors.Join(errors.New("can't write output file: "), err)
	}

	return nil
}
