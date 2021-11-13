package simulator

type LocusType int
const (
	LOCUS_UNDEFINED LocusType = iota
	LOCUS_DISCRETE
	LOCUS_CONTINUOUS
)

type ImplicitLocus struct {
	LocusType LocusType
	RangeMax float32
	Mutability float32
}

func NewImplicitLocus() *ImplicitLocus {
	rec := ImplicitLocus{}
	rec.Mutability = 0.0001
	if RandomFloat() < 0.5 {
		rec.LocusType = LOCUS_CONTINUOUS
		rec.RangeMax = 1.0
	} else {
		rec.LocusType = LOCUS_DISCRETE
		rec.RangeMax = float32(RandomInt(2, 50))
	}
	return &rec
}

func (rec *ImplicitLocus) GenerateValue() float32 {
	switch rec.LocusType {
		case LOCUS_CONTINUOUS:
			return RandomFloat()

		case LOCUS_DISCRETE:
			return float32(RandomInt(0, int(rec.RangeMax)))

		default:
			panic("Invalid locus type")
	}
}

func (rec *ImplicitLocus) GenerateModifiedValueFrom(prev float32) float32 {
	switch rec.LocusType {
		case LOCUS_CONTINUOUS:
			distance := RandomFloat() * 0.25
			if RandomFloat() < 0.5 {
				prev -= distance
				if prev < 0 {
					prev = 0
				}
			} else {
				prev += distance
				if prev > 1 {
					prev = 1
				}
			}
			return prev

		case LOCUS_DISCRETE:
			return rec.GenerateValue()

		default:
			panic("Invalid locus type")
	}
}

