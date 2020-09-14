package godash

import (
	"strconv"
)

func ItoS(i int) string {
	s := strconv.Itoa(i)
	return s
}

func I64toS(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}


func I64Ptr(v int64) *int64 {
	return &v
}

func IntPtr(v int) *int {
	return &v
}

func U64Ptr(v uint64) *uint64 {
	return &v
}

func I32Ptr(v int32) *int32 {
	return &v
}


func StrI64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StrF64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func F32Str(s string) (ret float64) {
	ret, _ = strconv.ParseFloat(s, 32)
	return ret
}

func F64Str(s string) (ret float64) {
	ret, _ =  strconv.ParseFloat(s, 64)
	return ret
}

func I32Str(s string) (ret int64) {
	ret, _ =  strconv.ParseInt(s, 10, 32)
	return ret
}

func I64Str(s string) (ret int64) {
	ret, _ =  strconv.ParseInt(s, 10, 64)
	return ret
}
