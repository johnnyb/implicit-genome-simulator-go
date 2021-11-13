package simulator

import (
	"testing"
)

func TestSimulator(t *testing.T) {
	sim := NewSimulator(20, 100)
	sim.PerformIterations(10)

	t.Errorf("Done")
}
