package simulator

import (
	"fmt"
	"math/rand"
)

var RandomFloat = func()float32 {
	return rand.Float32()
}

var RandomInt = func(minval, maxval int) int {
	return rand.Intn(maxval - minval) + minval
}

var Log = func(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

var DataContext *Simulator

var DataLog = func(metric Metric, value interface{}) {
	// default - do nothing
}
