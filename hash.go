package looklock

import (
	"crypto/sha256"
)

func hash(answers []byte) []byte {
	result := sha256.Sum256(answers)
	return result[:]
}
