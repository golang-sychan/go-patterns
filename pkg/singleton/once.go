package singleton

import "sync"

var onceInstance *instance

var o sync.Once

func GetOnceInstance () *instance {
	o.Do(func() {
		onceInstance = new()
	})
	return onceInstance
}
