package simulator

var MAX_FITNESS float32 = 5

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
		maxFitness := RandomFloat() * MAX_FITNESS
		if minFitness > maxFitness {
			tmp := minFitness
			minFitness = maxFitness
			maxFitness = tmp
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
	if distFromOptimal < 0 {
		distFromOptimal = 0 - distFromOptimal // Absolute value
	}
	distFit := distFromOptimal * (fdata.FitnessMax - fdata.FitnessMin)
	fitness := fdata.FitnessMin + distFit

	return fitness
}
