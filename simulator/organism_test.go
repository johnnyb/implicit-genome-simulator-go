package simulator

import (
	"fmt"
	"math"
	"testing"
)

func TestOrganismDynamics(t *testing.T) {
	ReseedAndPrint()

	igenome := NewImplicitGenome(10, DEFAULT_MUTABILITY)
	sim := &Simulator{}
	o := NewOrganism(sim, igenome)
	fmt.Println(o.String())

	newO := o.Duplicate()
	didEvolve := newO.Evolve()
	if didEvolve {
		fmt.Println(newO.String())
		fmt.Println(o.String())
	} else {
		fmt.Println("Did not evolve")
	}
}

func TestNumOffspringForFitness(t *testing.T) {
	ReseedAndPrint()

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
