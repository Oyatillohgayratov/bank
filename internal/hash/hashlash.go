package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}