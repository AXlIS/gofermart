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
	fmt.Println(password)
	return password
}

//func (s *SHA256Hasher) Base58Encoded(bytes []byte) (string, error) {
//	encoding := base58.BitcoinEncoding
//	encoded, err := encoding.Encode(bytes)
//	if err != nil {
//		log.Error().Msgf("encoded error: %s", err.Error())
//		return "", err
//	}
//	return string(encoded), nil
//}

//func (s *SHA256Hasher) Hash(url string) (string, error) {
//	urlHashBytes := s.Sha256Of(url)
//	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
//	finalString, err := s.Base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
//
//	if err != nil {
//		log.Error().Msgf("hash error: %s", err.Error())
//		return "", err
//	}
//
//	return finalString, nil
//}
