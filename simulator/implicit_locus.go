package simulator

import "fmt"

type LocusType int

const (
	LOCUS_UNDEFINED LocusType = iota
	LOCUS_DISCRETE
	LOCUS_CONTINUOUS
)

// ImplicitLocus is the generic type of a locus.
// Note that on the implementation, the value is always a float even if it is discrete.
// This makes the implementation easier and more uniform.
type ImplicitLocus struct {
	LocusId             int32     // Used for tracking
	LocusType           LocusType // LocusType tells whether this is a continuous or discrete-valued locus
	RangeMax            float32   // RangeMax tells the maximum of the states that this will take (total states = RangeMax + 1); only used on discrete-valued loci.
	ContinuousChangeMax float32   // The maximum amount that a continuous loci can change in a single mutation
	Mutability          float32   // Mutability gives the change that this locus will be mutated per iteration.
}

var lastLocusId int32 = 0
var DefaultMutability float32 = 0.01

// NewImplicitLocus creates a new implicit locus and sets the type and range randomly.
func NewImplicitLocus() *ImplicitLocus {
	lastLocusId += 1

	rec := ImplicitLocus{
		LocusId: lastLocusId,
	}
	rec.Mutability = DefaultMutability

	rec.ContinuousChangeMax = 0.25
	if RandomFloat() < 0.5 {
		rec.LocusType = LOCUS_CONTINUOUS
		rec.RangeMax = 1.0
	} else {
		rec.LocusType = LOCUS_DISCRETE
		rec.RangeMax = float32(RandomInt(2, 50))
	}
	return &rec
}

// GenerateValue generates a random valid value for this locus.
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

// GenerateModifiedValueFrom provides mutability for a locus.  Given a previous value, this determines what the next value should be.
// Discrete values are simply equivalent to GenerateValue().  Continuous values are bounded by (+/-) ContinuousChangeMax.
func (rec *ImplicitLocus) GenerateModifiedValueFrom(prev float32) float32 {
	switch rec.LocusType {
	case LOCUS_CONTINUOUS:
		// Calculate the new value
		distance := RandomFloat()*rec.ContinuousChangeMax*2 - rec.ContinuousChangeMax
		newVal := prev + distance

		// Bounds check
		if newVal > 1 {
			newVal = 1
		} else {
			if newVal < 0 {
				newVal = 0
			}
		}
		return newVal

	case LOCUS_DISCRETE:
		return rec.GenerateValue()

	default:
		panic("Invalid locus type")
	}
}

func (rec *ImplicitLocus) String() string {
	return fmt.Sprintf("L%d: %d/%f/%f/%f", rec.LocusId, rec.LocusType, rec.RangeMax, rec.ContinuousChangeMax, rec.Mutability)
}
