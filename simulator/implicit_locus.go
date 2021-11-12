package simulator

type LocusType int
const (
	LOCUS_BOOLEAN = iota
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
