package simulator

type Locus struct {
	ImplicitLocus *ImplicitLocus
	ValueDiscrete int
	ValueContinuous float32
}

func NewLocus(ilocus *ImplicitLocus) Locus {
	l := Locus {
		ImplicitLocus: ilocus,
	}
	switch ilocus.LocusType {
		case LOCUS_CONTINUOUS:
			l.ValueContinuous = ilocus.GenerateContinuousValue()

		case LOCUS_DISCRETE:
			l.ValueDiscrete = ilocus.GenerateDiscreteValue()

		default:
			panic("Invalid Locus Type")
	}

	return l
}

func (rec *Locus) PossiblyMutate() {
	if RandomFloat() < rec.ImplicitLocus.Mutability {
		rec.Mutate()
	}
}

func (rec *Locus) Mutate() {
	switch rec.ImplicitLocus.LocusType {
		case LOCUS_CONTINUOUS:
			rec.ValueContinuous = rec.ImplicitLocus.GenerateContinuousModificationFrom(rec.ValueContinuous)

		case LOCUS_DISCRETE:
			rec.ValueDiscrete = rec.ImplicitLocus.GenerateDiscreteValue()

		default:
			panic("Bad locus type")
	}
}
