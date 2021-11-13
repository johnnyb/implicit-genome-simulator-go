package simulator

type Environment struct {
	ImplicitGenome *ImplicitGenome
	FitnessData map[*ImplicitLocus]FitnessMapping
}

func NewEnvironment(igenome *ImplicitGenome) *Environment {
	rec := Environment{
		ImplicitGenome: igenome,
		FitnessData: map[*ImplicitLocus]FitnessMapping{},
	}
	for _, ilocus := range(igenome.ImplicitLoci) {
		minFitness := RandomFloat()
		maxFitness := RandomFloat() * 3
		if minFitness > maxFitness {
			tmp := minFitness
			minFitness = maxFitness
			maxFitness = tmp
		}
		if maxFitness < 1.2 {
			maxFitness += 1.0
		}
		optimal := ilocus.GenerateValue()
		rec.FitnessData[ilocus] = FitnessMapping{
			OptimalValue: optimal,
			FitnessMin: minFitness,
			FitnessMax: maxFitness,
		}
	}

	return &rec
}

func (rec *Environment) FitnessForLocus(locus Locus) float32 {
	ilocus := locus.ImplicitLocus
	fdata := rec.FitnessData[ilocus]
	distFromOptimal := ((fdata.OptimalValue - locus.Value) / locus.ImplicitLocus.RangeMax)
	distFit := distFromOptimal * (fdata.FitnessMax - fdata.FitnessMin)
	fitness := fdata.FitnessMin + distFit

	return fitness
}
