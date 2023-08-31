package singleton

import "testing"

func TestGetDCInstance(t *testing.T) {
	i := GetDCInstance()
	if i == nil {
		t.Fatal("can not init instance")
	}
	i2 := GetDCInstance()
	if i != i2 {
		t.Fatal("not singleton")
	}
	t.Logf("get instance :%v", i)
}
