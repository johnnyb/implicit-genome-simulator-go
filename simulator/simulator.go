package simulator

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Simulator struct {
	ImplicitGenome *ImplicitGenome
	Organisms      []*Organism
	Environment    *Environment
	MaxOrganisms   int
	Time           int
	DataStream     io.WriteCloser
	DataLogger     func(sim *Simulator, metric Metric, value interface{})
	Logger         func(sim *Simulator, message string)
}

var RandomSeed int64

func Seed(val int64) {
	RandomSeed = val
	rand.Seed(val)
}

func Reseed() int64 {
	seed := time.Now().UnixNano()
	Seed(seed)
	return seed
}

func ReseedAndPrint() {
	fmt.Printf("Running with seed: %d\n", Reseed())
}

func NewSimulator(numLoci, numOrganisms int, defaultMutability float32) *Simulator {
	igenome := NewImplicitGenome(numLoci, defaultMutability)
	rec := &Simulator{
		ImplicitGenome: igenome,
		Organisms:      make([]*Organism, numOrganisms),
		Time:           0,
		DataLogger:     func(sim *Simulator, metric Metric, value interface{}) {},
		Logger:         func(sim *Simulator, msg string) { fmt.Println(msg) },
	}

	for i := 0; i < numOrganisms; i++ {
		rec.Organisms[i] = NewOrganism(rec, igenome)
	}

	return rec
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

	rec.Log(fmt.Sprintf("Culled %d organisms", len(rec.Organisms)-len(culledOrganisms)))

	rec.Organisms = culledOrganisms
}

func (rec *Simulator) PerformIteration() {
	// If this is the first run, we might not have an environment yet
	if rec.Environment == nil {
		rec.SetEnvironment(NewEnvironment(rec.ImplicitGenome))
	}

	rec.DataLog(ITERATION_START, nil)
	rec.Time += 1
	rec.Log(fmt.Sprintf("Iteration %d: Organisms %d", rec.Time, len(rec.Organisms)))
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

	rec.DataLog(ITERATION_COMPLETE, nil)
}

func (rec *Simulator) PerformIterations(numIterations int) {
	for i := 0; i < numIterations; i++ {
		rec.PerformIteration()
	}
}

func (rec *Simulator) SetEnvironment(env *Environment) {
	if rec.Environment != nil {
		rec.DataLogger(rec, ENVIRONMENT_COMPLETE, rec.Environment)
	}
	rec.Environment = env
	rec.DataLogger(rec, ENVIRONMENT_START, env)
}

func (rec *Simulator) DataLog(metric Metric, value interface{}) {
	rec.DataLogger(rec, metric, value)
}

func (rec *Simulator) DataLogOutput(msg string) {
	rec.DataStream.Write([]byte(msg))
}

func (rec *Simulator) Log(msg string) {
	rec.Logger(rec, msg)
}

func (rec *Simulator) PossiblyChangeEnvironment() {
	// FIXME
}

func (rec *Simulator) Initialize() {
	rec.DataLog(SIMULATION_START, nil)
}

func (rec *Simulator) Finish() {
	rec.DataLog(SIMULATION_COMPLETE, nil)
	err := rec.DataStream.Close()
	if err != nil {
		rec.Log("Error closing datafile: " + err.Error())
	}
}
