package md5hex

import (
	"crypto/md5"
	"encoding/hex"
)

func Get(src string) string {
	hash := md5.Sum([]byte(src))
	return hex.EncodeToString(hash[:])
}
