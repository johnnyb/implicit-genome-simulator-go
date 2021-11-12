package simulator

type Environment struct {
	ImplicitGenome *ImplicitGenome
	FitnessData map[*ImplicitLocus]*FitnessMapping
}
