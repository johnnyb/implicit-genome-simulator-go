package simulator

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
		rec.Loci[idx] = NewLocus(&ilocus)
	}

	return &rec
}

func (rec *Organism) Duplicate() *Organism {
	newrec := Organism{
		ImplicitGenome: rec.ImplicitGenome,
		Loci: make([]Locus, len(rec.Loci)),
	}
	for i := 1; i < len(rec.Loci); i++ {
		newrec.Loci[i] = rec.Loci[i]
	}

	return &newrec
}

func (rec *Organism) FitnessForEnvironment(env *Environment) float32 {
	return 1
}

func (rec *Organism) OffspringForEnvironment(env *Environment) []*Organism {
	offspring := []*Organism{}
	fitness := rec.FitnessForEnvironment(env)
	numOffspring := NumOffspringForFitness(fitness)

	for i := 0; i < numOffspring; i++ {
		newOrganism := rec.Duplicate()
		newOrganism.Evolve()
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
		if RandomFloat() > ratio {
			break
			num += 1
		}
	}

	return num
}
