package simulator

import (
	"fmt"
)

var MAX_FITNESS float32 = 5

type Environment struct {
	EnvironmentId  int32
	ImplicitGenome *ImplicitGenome
	FitnessData    map[*ImplicitLocus]FitnessMapping
}

var lastEnvironmentId int32 = 0

func NewEnvironment(igenome *ImplicitGenome) *Environment {
	lastEnvironmentId += 1
	rec := Environment{
		EnvironmentId:  lastEnvironmentId,
		ImplicitGenome: igenome,
		FitnessData:    map[*ImplicitLocus]FitnessMapping{},
	}
	for _, ilocus := range igenome.ImplicitLoci {
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
			FitnessMin:   minFitness,
			FitnessMax:   maxFitness,
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
	fitness := fdata.FitnessMax - distFit

	return fitness
}

func (rec *Environment) String() string {
	desc := fmt.Sprintf("Environment %d\n", rec.EnvironmentId)
	for l, f := range rec.FitnessData {
		desc += fmt.Sprintf("  L%d: %f/%f/%f\n", l.LocusId, f.FitnessMin, f.FitnessMax, f.OptimalValue)
	}
	return desc
}
