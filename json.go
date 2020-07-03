package godash

import (
	"encoding/json"
	"log"
)


func ConvertToAny(source interface{}, ptrTo interface{}) error {
	by, err := json.Marshal(source)
	if nil != err {
		return err
	}
	err = json.Unmarshal(by, ptrTo)
	if nil != err {
		return err
	}
	return nil
}

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
