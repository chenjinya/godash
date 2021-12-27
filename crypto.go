package godash

import (
	"crypto/md5"
	"fmt"
)

func MD5(origin string) string {
	data := []byte(origin)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}