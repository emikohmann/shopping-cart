package hashing

import (
	"crypto/md5"
	"fmt"
)

func MD5(input string) string {
	data := []byte(input)
	return fmt.Sprintf("%x", md5.Sum(data))
}
