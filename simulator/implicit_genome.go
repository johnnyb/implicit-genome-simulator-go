package simulator

import "fmt"

// ImplicitGenome is a collection of implicit loci
type ImplicitGenome struct {
	ImplicitLoci []*ImplicitLocus
}

// NewImplicitGenome creates an ImplicitGenome with a specific number of loci
func NewImplicitGenome(numLoci int) *ImplicitGenome {
	loci := make([]*ImplicitLocus, numLoci)
	for i := 0; i < numLoci; i++ {
		loci[i] = NewImplicitLocus()
	}

	return &ImplicitGenome{
		ImplicitLoci: loci,
	}
}

func (rec *ImplicitGenome) String() string {
	description := "IGENOME:\n"
	for _, ilocus := range rec.ImplicitLoci {
		description += fmt.Sprintf("  %s\n", ilocus.String())
	}

	return description
}
