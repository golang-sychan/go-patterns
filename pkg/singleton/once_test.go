package singleton

import "testing"

func TestGetOnceInstance(t *testing.T) {
	i := GetOnceInstance()
	if i == nil {
		t.Fatal("can not init instance")
	}
	i2 := GetOnceInstance()
	if i != i2 {
		t.Fatal("not singleton")
	}
	t.Logf("get once instance :%v", i)
}
