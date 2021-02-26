package godash

import (
	"log"
	"strconv"
	"strings"
)

// 将十进制转为 32 进制
const safeStrCode32 = "0123456789abcdefghijklmnopqrstuv"
const safeStrCode37 = "0123456789abcdefghijklmnopqrstuvxyz_-"

func NumberToShortStr(num int64, args ...int64) string {
	base := int64(32)
	if len(args) > 0 && args[0] > 0 && args[0] < int64(len(safeStrCode37)) {
		base = args[0]
	}
	var converNums []string
	for {
		rem := num % base
		num = num / base
		converNums = append(converNums, safeStrCode37[rem:rem+1])
		if num == 0 {
			break
		}
	}
	for i, j := 0, len(converNums)-1; i < j; i, j = i+1, j-1 {
		converNums[i], converNums[j] = converNums[j], converNums[i]
	}
	return strings.Join(converNums, "")
}


func StrPtrIsBlank(v *string) bool {
	if v == nil {
		return true
	}
	return StringIsBlank(*v)
}

func StringIsBlank(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}


func Email(s string) bool {
	if !strings.Contains(s, "@") || s[0] == '@' || s[len(s)-1] == '@' {
		return false
	}
	return true
}
func MustString(v interface{}) string {
	if nil == v {
		return ""
	}
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return ""
	}
}
func StrPtr(v string) *string {
	return &v
}

func PtrStr(v *string) string {
	if nil == v {
		return ""
	}
	return *v
}

func Stoi(v interface{}) int {
	if nil == v {
		return 0
	}
	s := v.(string)
	if "" == s {
		return 0
	}
	i, err := strconv.Atoi(s)
	if nil != err {
		log.Printf("strconv.Atoi: %s err: %s", s, err.Error())
		return 0
	}
	return i
}

func Stoi64(v interface{}) int64 {
	if nil == v {
		return int64(0)
	}
	s := v.(string)
	i, err := strconv.ParseInt(s, 0, 64)
	if nil != err {
		log.Printf("strconv.ParseInt: %s err: %s", s, err.Error())
		return int64(0)
	}
	return i
}

func Stof64(v interface{}) float64 {
	if nil == v {
		return float64(0)
	}
	s := v.(string)
	if "" == s {
		return float64(0)
	}
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		log.Printf("strconv.ParseFloat: %s err: %s", s, err.Error())
		return float64(0)
	}
	return i
}
