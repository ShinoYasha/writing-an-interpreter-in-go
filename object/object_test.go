package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "hello"}
	hello2 := &String{Value: "hello"}
	diff1 := &String{Value: "diff"}
	diff2 := &String{Value: "diff"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("string with same content have different value")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("string with same content have different value")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("string with different content have same value")
	}
}
