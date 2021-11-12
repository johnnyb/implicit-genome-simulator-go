package simulator

type Locus struct {
	ImplicitLocus *ImplicitLocus
	ValueDiscrete int
	ValueContinuous float32
}

func (rec *Locus) PossiblyMutate() {
	if RandomFloat() < rec.ImplicitLocus.Mutability {
		rec.Mutate()
	}
}

func (rec *Locus) Mutate() {
	switch rec.ImplicitLocus.LocusType {
		case LOCUS_CONTINUOUS:
			distance := RandomFloat() * 0.25
			if RandomFloat() < 0.5 {
				rec.ValueContinuous -= distance
				if rec.ValueContinuous < 0 {
					rec.ValueContinuous = 0
				}
			} else {
				rec.ValueContinuous += distance
				if rec.ValueContinuous > 1 {
					rec.ValueContinuous = 1
				}
			}
		case LOCUS_DISCRETE:

		default:
			panic("Bad locus type")
	}
}
