package simulator

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNumOffspringForFitness(t *testing.T) {
	seed := time.Now().UnixNano()
	// fmt.Printf("Running with seed: %d\n", seed)
	rand.Seed(seed)

	fitnesses := []float32{0.5, 1.0, 2.0, 4.0, 0.25}
	for _, fitness := range fitnesses {
		offspring := 0
		iterations := 1000
		expectation := int(fitness * float32(iterations))
		for idx := 0; idx < iterations; idx++ {
			offspring += NumOffspringForFitness(fitness)
		}
		difference := math.Abs(float64(offspring)/float64(expectation) - 1.0)
		if difference > 0.2 {
			t.Errorf("Substantial difference detected: %f/%d/%d/%f", fitness, expectation, offspring, difference)
		}
	}
}
