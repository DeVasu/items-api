package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	defer hash.Reset()
	return hex.EncodeToString(hash.Sum(nil))
}
