package godash

func Unique(s []interface{}) []interface{} {

	m := make(map[interface{}]bool)
	for _, v := range s {
		m[v] = true
	}
	var res []interface{}
	for v, _ := range m {
		res = append(res, v)
	}
	return res
}
