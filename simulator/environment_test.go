package simulator

import (
	"fmt"
	"testing"
)

func TestEnvironmentDynamics(t *testing.T) {
	ReseedAndPrint()

	igenome := NewImplicitGenome(10, DEFAULT_MUTABILITY)
	fmt.Println(igenome.String())
	env := NewEnvironment(igenome)
	fmt.Println(env.String())
	sim := &Simulator{}
	o := NewOrganism(sim, igenome)
	fmt.Println(o.String())

	fmt.Println("Fitness:")
	for _, locus := range o.Loci {
		curFit := env.FitnessForLocus(locus)
		fmt.Printf("  L%d: %f\n", locus.ImplicitLocus.LocusId, curFit)
	}

	fitness := o.FitnessForEnvironment(env)
	fmt.Printf("Total Fitness: %f\n", fitness)
}
