package simulator

import (
	"testing"
)

func TestSimulator(t *testing.T) {
	sim := NewSimulator(20, 1000)
	sim.PerformIterations(100)

	t.Errorf("Done")
}
