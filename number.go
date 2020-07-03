package godash

import (
	"strconv"
)

func Itos(i int) string {
	s := strconv.Itoa(i)
	return s
}

func I64tos(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}
