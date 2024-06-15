package Utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func EncryptString(plainText string, key string) (cipherText string, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plainText)
	blockSize := block.BlockSize()
	plainTextBytes = PKCS5P(plainTextBytes, blockSize)

	cipherTextBytes := make([]byte, blockSize+len(plainTextBytes))
	iv := cipherTextBytes[:blockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherTextBytes[blockSize:], plainTextBytes)

	return base64.URLEncoding.EncodeToString(cipherTextBytes), nil
}

func DecryptString(cipherText string, key string) (plainText string, err error) {
	cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(cipherTextBytes) < block.BlockSize() {
		return "", errors.New("cipherText too short")
	}

	iv := cipherTextBytes[:block.BlockSize()]
	cipherTextBytes = cipherTextBytes[block.BlockSize():]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherTextBytes, cipherTextBytes)

	plainTextBytes := PKCS5U(cipherTextBytes)
	if plainTextBytes == nil {
		return "", errors.New("PKCS5Unpadding error")
	}

	return string(plainTextBytes), nil
}

func PKCS5P(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	text := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, text...)
}

func PKCS5U(plaintext []byte) []byte {
	length := len(plaintext)
	padding := int(plaintext[length-1])
	return plaintext[:(length - padding)]
}
