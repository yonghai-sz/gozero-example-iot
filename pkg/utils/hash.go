package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
