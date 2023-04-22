package hash_passwords

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordHash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func CompareHashAndPassword(password, hash string) bool {
	passwordHash := GeneratePasswordHash(password)
	return hash == passwordHash
}
