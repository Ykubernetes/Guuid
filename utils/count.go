package utils

import "sync/atomic"

var num uint32 = 0

func DeltaIdCounter() (newValue uint32) {
	addValue := uint32(1)
	newValue = atomic.AddUint32(&num, addValue)
	return
}
