package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
)

const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
const newUrlLen = 10

func GetSeed(s string) int64 {
	hash := sha256.Sum256([]byte(s))
	seed := hash[:]
	seedInt64 := binary.BigEndian.Uint64(seed)
	return int64(seedInt64)
}

func ShortenUrl(url string) string {
	source := rand.NewSource(GetSeed(url))
	result := make([]rune, 10)
	for i := 0; i < newUrlLen; i++ {
		result[i] = rune(letterRunes[source.Int63()%int64(len(letterRunes))])
	}
	return string(result)
}
