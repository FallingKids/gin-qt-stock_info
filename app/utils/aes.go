package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Encrypt(data []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// PKCS7Padding
	data = pkcs7Padding(data, block.BlockSize())

	ciphertext := make([]byte, len(data))

	// CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, data)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(data string, key []byte, iv []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))

	// CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// PKCS7Padding
	plaintext = pkcs7UnPadding(plaintext)

	return plaintext, nil
}

// PKCS7Padding 填充函数
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7Padding 去除函数
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
