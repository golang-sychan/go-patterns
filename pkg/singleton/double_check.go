package singleton

import "sync"

type instance struct {
	attr struct{}
}

func new() *instance {
	return &instance{attr: struct{}{}}
}

var doubleChkInstance *instance
var mu sync.Mutex

func GetDCInstance() *instance {
	if doubleChkInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if doubleChkInstance == nil {
			doubleChkInstance = new()
		}
	}
	return doubleChkInstance
}
