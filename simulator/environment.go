package simulator

type Environment struct {
	ImplicitGenome *ImplicitGenome
	FitnessData map[*ImplicitLocus]*FitnessMapping
}

func (rec *Environment) FitnessForLocus(locus Locus) float32 {
	ilocus := locus.ImplicitLocus
	fdata := rec.FitnessData[ilocus]
	panic(fdata)
	// distFromOptimal := (
	return 1
}
