package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"os/user"
)

type DataCrypto struct {
	key []byte
}

func NewDataCrypto() *DataCrypto {
	key := getNodeID()
	return &DataCrypto{key: key}
}

func getNodeID() []byte {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return []byte(usr.Uid)
}

func (d *DataCrypto) createAESKey() []byte {
	hash := sha256.New()
	hash.Write(d.key)
	return hash.Sum(nil)[:32]
}

func (d *DataCrypto) createAESIv() []byte {
	hash := sha256.New()
	hash.Write(d.key)
	return hash.Sum(nil)[16:32]
}

func (d *DataCrypto) Encrypt(message string) string {
	block, err := aes.NewCipher(d.createAESKey())
	if err != nil {
		log.Fatalf("Error creating cipher: %v", err)
	}

	iv := d.createAESIv()

	plaintext := pad([]byte(message), aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	copy(ciphertext[:aes.BlockSize], iv)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func (d *DataCrypto) Decrypt(ciphertextBase64 string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		log.Fatalf("Error decoding base64: %v", err)
	}

	block, err := aes.NewCipher(d.createAESKey())
	if err != nil {
		log.Fatalf("Error creating cipher: %v", err)
	}

	if len(ciphertext) == 0 {
		log.Fatalf("配置文件错误.")
	}

	iv := ciphertext[:aes.BlockSize]

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	mode.CryptBlocks(plaintext, ciphertext[aes.BlockSize:])

	plaintext = unpad(plaintext)

	return string(plaintext)
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padText[i] = byte(padding)
	}
	return append(src, padText...)
}

func unpad(src []byte) []byte {
	padding := int(src[len(src)-1])
	return src[:len(src)-padding]
}
