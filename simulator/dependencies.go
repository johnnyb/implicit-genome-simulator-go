package simulator

// This file is a list of external and/or replaceable dependencies.

import (
	"fmt"
	"math/rand"
)

// RandomFloat is a function which returns a random floating point value between 0 and 1.  By default it uses the standard PRNG
var RandomFloat = func()float32 {
	return rand.Float32()
}

// RandomInt is a function which returns a random integer between minval and maxval.  By default it uses the standard PRNG.
var RandomInt = func(minval, maxval int) int {
	return rand.Intn(maxval - minval) + minval
}

// Log is the generic logging function.  It's basically a replacement for fmt.Printf so that I don't have to import/unimport "fmt" when I want to add Printf statements.
// In general, DataLog should be used instead of Log.
var Log = func(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// DataContext is the active Simulator 
var DataContext *Simulator

// DataLog is the generic data logging function.  
// It takes a metric and a value for that metric.
// The default mechanism is to do nothing.
// data_logging.go has some basic data logging functions you can use.
var DataLog = func(metric Metric, value interface{}) {
	// default - do nothing
}
