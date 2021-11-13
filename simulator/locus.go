package simulator

type Locus struct {
	ImplicitLocus *ImplicitLocus
	Value float32
}

func NewLocus(ilocus *ImplicitLocus) Locus {
	if ilocus == nil {
		panic("Nil ilocus")
	}
	l := Locus {
		ImplicitLocus: ilocus,
		Value: ilocus.GenerateValue(),
	}

	return l
}

func (rec *Locus) PossiblyMutate() {
	if RandomFloat() < rec.ImplicitLocus.Mutability {
		rec.Mutate()
	}
}

func (rec *Locus) Mutate() {
	rec.Value = rec.ImplicitLocus.GenerateModifiedValueFrom(rec.Value)
}

