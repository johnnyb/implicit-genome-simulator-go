package simulator

import (
	"math/rand"
	"testing"
	"time"
)

// This test just makes sure that the simulator runs without error
func TestSimulator(t *testing.T) {
	// Create simulator
	sim := NewSimulator(10, 100, DEFAULT_MUTABILITY)
	sim.MaxOrganisms = 10000

	// Setup Data Logging
	DataContext = sim
	DataLog = DataLogBeneficialMutations

	// Seed PRNG
	rand.Seed(time.Now().UnixNano())

	// DO IT!
	sim.PerformIterations(100)
}
