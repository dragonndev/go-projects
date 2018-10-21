package crypto

import (
	"crypto/rand"
	"log"
)

type encryptor interface {
	Encrypt() []byte
	Decrypt() []byte
	GetKey() []byte
}

type aesEncryptor struct {
	plainTextBytes []byte
	encryptedBytes []byte
	encKey         []byte
	iv             []byte
}

func (aes aesEncryptor) Encrypt() []byte {
	return nil
}

func (aes aesEncryptor) Decrypt() []byte {
	return nil
}

func (aes aesEncryptor) GetKey() []byte {
	return getRandomBytes(32)
}
func (aes aesEncryptor) GetIV() []byte {
	return getRandomBytes(16)
}

func getRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatalf("Failed to generate IV %s", err)
	}
	return bytes
}
