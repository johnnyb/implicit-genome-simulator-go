package simulator

import (
	"testing"
	"time"
	"math/rand"
)

func TestSimulator(t *testing.T) {
	// Create simulator
	sim := NewSimulator(10, 100)

	// Setup Data Logging
	DataContext = sim 
	DataLog = DataLogBeneficialMutations

	// Seed PRNG
	rand.Seed(time.Now().UnixNano()) 

	// DO IT!
	sim.PerformIterations(1000)
}
