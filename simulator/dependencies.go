package simulator

import (
	"math/rand"
)

var RandomFloat = func()float32 {
	return rand.Float32()
}

var RandomInt = func(minval, maxval int) int {
	return rand.Intn(maxval - minval) + minval
}
