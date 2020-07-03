package godash

import (
	"encoding/json"
	"log"
)

func JSONStringify(origin interface{}) string {
	if nil == origin {
		return ""
	}
	d, err := json.Marshal(origin)
	if nil != err {
		log.Fatalln(err)
		return ""
	}
	return string(d)
}

func JSONIndent(origin interface{}) string {
	if nil == origin {
		return ""
	}
	d, err := json.MarshalIndent(origin, "", " ")
	if nil != err {
		log.Fatalln(err)
		return ""
	}
	return string(d)
}


func JSONnumber(v interface{}) int64 {
	if nil == v {
		return int64(0)
	}
	return int64(v.(float64))
}
