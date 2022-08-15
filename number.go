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

func HexToI64(s string) (ret int64) {
	ret, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func DecConvertToX(n int, base int) string {
	if n < 0 {
		panic("n < 0")
	}
	if base != 2 && base != 8 && base != 16 {
		panic("only support base : 2, 8, 16")
	}
	result := ""
	h:=map[int]string{
		0:"0",
		1:"1",
		2:"2",
		3:"3",
		4:"4",
		5:"5",
		6:"6",
		7:"7",
		8:"8",
		9:"9",
		10:"A",
		11:"B",
		12:"C",
		13:"D",
		14:"E",
		15:"F",
	}
	for ; n > 0; n /= base {
		lsb := h[n % base]
		result = lsb + result
	}
	return result
}

func I64ToHex(i int64) (ret string) {
	return DecConvertToX(int(i), 16)
}