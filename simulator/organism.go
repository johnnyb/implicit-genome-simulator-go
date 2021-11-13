package simulator

import (
	// "fmt"
)

type Organism struct {
	ImplicitGenome *ImplicitGenome
	Loci []Locus
}

func NewOrganism(igenome *ImplicitGenome) *Organism {
	rec := Organism{
		ImplicitGenome: igenome,
		Loci: make([]Locus, len(igenome.ImplicitLoci)),
	}
	for idx, ilocus := range(igenome.ImplicitLoci) {
		rec.Loci[idx] = NewLocus(ilocus)
	}

	return &rec
}

func (rec *Organism) Duplicate() *Organism {
	newrec := Organism{
		ImplicitGenome: rec.ImplicitGenome,
		Loci: make([]Locus, len(rec.Loci)),
	}
	for i := 0; i < len(rec.Loci); i++ {
		newrec.Loci[i] = rec.Loci[i]
	}

	return &newrec
}

func (rec *Organism) FitnessForEnvironment(env *Environment) float32 {
	var fitnessSum float32 = 0
	for _, locus := range(rec.Loci) {
		fitnessSum += env.FitnessForLocus(locus)
	}
	fitness := fitnessSum / ((float32)(len(rec.Loci)))

	return fitness
}

func (rec *Organism) OffspringForEnvironment(env *Environment) []*Organism {
	offspring := []*Organism{}
	fitness := rec.FitnessForEnvironment(env)
	numOffspring := NumOffspringForFitness(fitness)

	// fmt.Printf("Organism: Fitness/%f Offspring/%d\n", fitness, numOffspring)

	for i := 0; i < numOffspring; i++ {
		newOrganism := rec.Duplicate()
		newOrganism.Evolve()
		offspring = append(offspring, newOrganism)
	}

	return offspring
}

func (rec *Organism) Evolve() {
	for idx := range(rec.Loci) {
		rec.Loci[idx].PossiblyMutate()
	}
}

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
