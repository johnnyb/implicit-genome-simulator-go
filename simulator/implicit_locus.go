package simulator

type LocusType int
const (
	LOCUS_UNDEFINED LocusType = iota
	LOCUS_BOOLEAN
	LOCUS_DISCRETE
	LOCUS_CONTINUOUS
)

type ImplicitLocus struct {
	LocusType LocusType
	RangeMaxInt int
	RangeMaxFloat float32
	Mutability float32
}

func NewImplicitLocus() {

}
