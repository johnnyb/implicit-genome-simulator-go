package simulator

// This file is a list of external and/or replaceable dependencies.

import (
	"math/rand"
)

// RandomFloat is a function which returns a random floating point value between 0 and 1.  By default it uses the standard PRNG
var RandomFloat = func() float32 {
	return rand.Float32()
}

// RandomInt is a function which returns a random integer between minval and maxval.  By default it uses the standard PRNG.
var RandomInt = func(minval, maxval int) int {
	return rand.Intn(maxval-minval) + minval
}
