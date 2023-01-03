package simulator

import (
	"fmt"
	"math/rand"
	"time"
)

type Simulator struct {
	ImplicitGenome *ImplicitGenome
	Organisms      []*Organism
	Environment    *Environment
	MaxOrganisms   int
	Time           int
}

func Reseed() int64 {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return seed
}

func ReseedAndPrint() {
	fmt.Printf("Running with seed: %d\n", Reseed())
}

func NewSimulator(numLoci, numOrganisms int) *Simulator {
	igenome := NewImplicitGenome(numLoci)
	rec := Simulator{
		ImplicitGenome: igenome,
		Environment:    NewEnvironment(igenome),
		Organisms:      make([]*Organism, numOrganisms),
		Time:           0,
	}

	for i := 0; i < numOrganisms; i++ {
		rec.Organisms[i] = NewOrganism(igenome)
	}

	return &rec
}

func (rec *Simulator) CullOrganisms() {
	if rec.MaxOrganisms > 0 {
		rec.CullToSizeNonStrict(rec.MaxOrganisms)
	}
}

func (rec *Simulator) CullToSizeNonStrict(numOrganisms int) {
	ratio := float32(numOrganisms) / float32(len(rec.Organisms))
	if ratio >= 1.0 {
		return
	}

	culledOrganisms := []*Organism{}
	for _, o := range rec.Organisms {
		if RandomFloat() <= ratio {
			culledOrganisms = append(culledOrganisms, o)
		}
	}

	Log("Culled %d organisms", len(rec.Organisms)-len(culledOrganisms))

	rec.Organisms = culledOrganisms
}

func (rec *Simulator) PerformIteration() {
	DataLog(ITERATION_START, nil)
	rec.Time += 1
	Log("Iteration %d: Organisms %d", rec.Time, len(rec.Organisms))
	newOrganisms := []*Organism{}
	for _, o := range rec.Organisms {
		offspring := o.OffspringForEnvironment(rec.Environment)
		for _, newO := range offspring {
			newOrganisms = append(newOrganisms, newO)
		}
	}
	rec.Organisms = newOrganisms

	rec.CullOrganisms()

	rec.PossiblyChangeEnvironment()

	DataLog(ITERATION_COMPLETE, nil)
}

func (rec *Simulator) PerformIterations(numIterations int) {
	for i := 0; i < numIterations; i++ {
		rec.PerformIteration()
	}
}

func (rec *Simulator) PossiblyChangeEnvironment() {
	// FIXME
}
