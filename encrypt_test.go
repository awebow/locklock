package looklock

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	block, _ := aes.NewCipher([]byte("acskdkwkgoaosockxkzkwkskakd12035"))
	fmt.Println(hex.EncodeToString(encrypt(block, []byte("Test <16B"))))
	fmt.Println(hex.EncodeToString(encrypt(block, []byte("Test 16BytesText"))))
	fmt.Println(hex.EncodeToString(encrypt(block, []byte("Test more than 16bytes string"))))
}

func TestDecrypt(t *testing.T) {
	block, _ := aes.NewCipher([]byte("acskdkwkgoaosockxkzkwkskakd12035"))
	data, _ := hex.DecodeString("fc99ec4df059fd3d4ceb65a66fd4947172b48a2051f35aa49d56fe5f5f8aeafb")
	plain := string(decrypt(block, data))
	assert.Equal(t, "Test <16B", plain)
}
