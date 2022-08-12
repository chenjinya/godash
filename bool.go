package godash

func BoolPtr(v bool) *bool {
	return &v
}

func PtrSBool(v *bool) bool {
	if nil == v {
		return false
	}
	return *v
}