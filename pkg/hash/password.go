package hash

import (
	"crypto/sha256"
	"fmt"
)

type Hasher interface {
	Hash(value string) string
}

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{
		salt: salt,
	}
}

func (s *SHA256Hasher) Hash(initialString string) string {
	encoder := sha256.New()
	encoder.Write([]byte(initialString))
	password := fmt.Sprintf("%x", encoder.Sum([]byte(s.salt)))
	return password
}
