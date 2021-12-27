package godash

import (
	"bytes"
	"encoding/json"
)

func StructToMap(srcStruct interface{}) (map[string]interface{}, error) {
	dst := make(map[string]interface{})
	srcBtyes, err := json.Marshal(srcStruct)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(srcBtyes))
	decoder.UseNumber()
	err = decoder.Decode(&dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}


func StructToMapStrict(srcStruct interface{}) (map[string]interface{}, error) {
	dst := make(map[string]interface{})
	srcBtyes, err := json.Marshal(srcStruct)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(srcBtyes))
	decoder.UseNumber()
	err = decoder.Decode(&dst)
	if err != nil {
		return nil, err
	}
	for k, v := range dst {
		if nil == v {
			delete(dst, k)
		}
	}
	return dst, nil
}
