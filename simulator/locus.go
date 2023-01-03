package simulator

import "fmt"

// Locus is the real locus within an Organism, based on an ImplicitLocus (which defines the possible range and mutability).
type Locus struct {
	ImplicitLocus *ImplicitLocus
	Value         float32
}

// Generates a new locus based on an implicit locus and generates a random initial value for it (within the proper range).
func NewLocus(ilocus *ImplicitLocus) Locus {
	if ilocus == nil {
		panic("Nil ilocus")
	}
	l := Locus{
		ImplicitLocus: ilocus,
		Value:         ilocus.GenerateValue(),
	}

	return l
}

// Determines, based on the mutation rate, if we should mutate this locus.
func (rec *Locus) PossiblyMutate() bool {
	if RandomFloat() < rec.ImplicitLocus.Mutability {
		rec.Mutate()
		return true
	}

	return false
}

// Generates a new value for this locus, potentially based on the previous value.
func (rec *Locus) Mutate() {
	rec.Value = rec.ImplicitLocus.GenerateModifiedValueFrom(rec.Value)
}

func (rec *Locus) String() string {
	return fmt.Sprintf("L%d: %f", rec.ImplicitLocus.LocusId, rec.Value)
}
