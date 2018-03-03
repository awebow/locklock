package looklock

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(block cipher.Block, raw []byte) []byte {
	padding := aes.BlockSize - len(raw)%aes.BlockSize
	raw = append(raw, bytes.Repeat([]byte{byte(padding)}, padding)...)

	result := make([]byte, aes.BlockSize+len(raw))
	iv := result[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(result[aes.BlockSize:], raw)

	return result
}

func decrypt(block cipher.Block, ciphertext []byte) []byte {
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	result := make([]byte, len(ciphertext))
	paddedLength := len(result)
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(result, ciphertext)

	result = result[:paddedLength-int(result[paddedLength-1])]

	return result
}
