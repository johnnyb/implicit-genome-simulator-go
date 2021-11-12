package simulator

type LocusType int
const (
	LOCUS_UNDEFINED LocusType = iota
	LOCUS_DISCRETE
	LOCUS_CONTINUOUS
)

type ImplicitLocus struct {
	LocusType LocusType
	RangeMax int
	Mutability float32
}

func NewImplicitLocus() *ImplicitLocus {
	rec := ImplicitLocus{}
	rec.Mutability = 0.0001
	if RandomFloat() < 0.5 {
		rec.LocusType = LOCUS_CONTINUOUS
	} else {
		rec.LocusType = LOCUS_DISCRETE
		rec.RangeMax = RandomInt(2, 50)
	}
	return &rec
}

func (rec *ImplicitLocus) GenerateDiscreteValue() int {
	return RandomInt(0, rec.RangeMax)
}

func (rec *ImplicitLocus) GenerateContinuousValue() float32 {
	return RandomFloat()
}
