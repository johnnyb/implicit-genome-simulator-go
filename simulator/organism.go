package simulator

import "fmt"

// Organism is a collection of loci.
type Organism struct {
	ImplicitGenome *ImplicitGenome
	Loci           []Locus
}

// Creates a new organism with the given implicit genome.
// Assigns initial values of loci randomly
func NewOrganism(igenome *ImplicitGenome) *Organism {
	rec := Organism{
		ImplicitGenome: igenome,
		Loci:           make([]Locus, len(igenome.ImplicitLoci)),
	}
	for idx, ilocus := range igenome.ImplicitLoci {
		rec.Loci[idx] = NewLocus(ilocus)
	}

	return &rec
}

// Duplicate does a perfect duplication of the organism (no mutations).  Use Evolve() on the duplicated organism to potentially get mutations.
func (rec *Organism) Duplicate() *Organism {
	newrec := Organism{
		ImplicitGenome: rec.ImplicitGenome,
		Loci:           make([]Locus, len(rec.Loci)),
	}
	for i := 0; i < len(rec.Loci); i++ {
		newrec.Loci[i] = rec.Loci[i]
	}

	return &newrec
}

// Evolve goes through an organism's loci and applies random changes based on the mutation rate (using PossiblyMutate from the Locus).
func (rec *Organism) Evolve() bool {
	mutated := false
	for idx := range rec.Loci {
		if rec.Loci[idx].PossiblyMutate() {
			mutated = true
		}
	}
	return mutated
}

// FitnessForEnvironment evaluates the fitness of the total organism in a given environment.
func (rec *Organism) FitnessForEnvironment(env *Environment) float32 {
	var fitnessSum float32 = 0
	for _, locus := range rec.Loci {
		fitnessSum += env.FitnessForLocus(locus)
	}
	fitness := fitnessSum / ((float32)(len(rec.Loci)))

	return fitness
}

// OffspringForEnvironment duplicates and evolves an organism based on their environment.
func (rec *Organism) OffspringForEnvironment(env *Environment) []*Organism {
	offspring := []*Organism{}
	fitness := rec.FitnessForEnvironment(env)
	numOffspring := NumOffspringForFitness(fitness)

	// Log("Organism: Fitness/%f Offspring/%d\n", fitness, numOffspring)

	for i := 0; i < numOffspring; i++ {
		newOrganism := rec.Duplicate()
		didEvolve := newOrganism.Evolve()
		offspring = append(offspring, newOrganism)

		// Is it more or less fit?
		newFitness := newOrganism.FitnessForEnvironment(env)

		if didEvolve {
			DataLog(ORGANISM_FITNESS_DIFFERENCE, newFitness-fitness)
			DataLog(ORGANISM_MUTATIONS_BENEFICIAL, newFitness > fitness)
		}
	}

	return offspring
}

func (rec *Organism) String() string {
	description := "ORGANISM:\n"
	for _, locus := range rec.Loci {
		description += fmt.Sprintf("  %s\n", locus.String())
	}

	return description
}

// NumOffspringForFitness uses the PRNG to determine, based on fitness, how many offspring an organism should have.
func NumOffspringForFitness(fitness float32) int {
	num := 0
	ratio := fitness / (fitness + 1)
	for {
		tmp := RandomFloat()
		if tmp > ratio {
			break
		}
		num += 1
	}

	return num
}
