package simulator

type ImplicitGenome struct {
	ImplicitLoci []*ImplicitLocus
}

func NewImplicitGenome(numLoci int) *ImplicitGenome {
	loci := make([]*ImplicitLocus, numLoci)
	for i := 0; i < numLoci; i++ {
		loci[i] = NewImplicitLocus()
	}

	return &ImplicitGenome{
		ImplicitLoci: loci,
	}
}
