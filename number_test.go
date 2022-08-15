package godash

import "testing"

func TestI64ToHex(t *testing.T) {
	i := int64(839234)
	e := "CCE42"
	a := I64ToHex(i)
	if a != e {
		t.Fatalf("结果不正确，actul: %v, expect: %v", a, e)
	}
}

func TestHexToI64(t *testing.T) {
	i := "CCE42"
	e := int64(839234)
	a := HexToI64(i)
	if a != e {
		t.Fatalf("结果不正确，actul: %v, expect: %v", a, e)
	}
}
