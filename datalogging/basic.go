package datalogging

import (
	"fmt"

	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
)

var currentTime int = 1
var beneficialCount int = 0
var deleteriousCount int = 0

var currentEnvironment *simulator.Environment

// DataLogBeneficialMutations is a logging function which gives the beneficial/deleterious ratio.
func DataLogBeneficialMutations(sim *simulator.Simulator, metric simulator.Metric, value interface{}) {
	switch metric {
	case simulator.ENVIRONMENT_START:
		currentEnvironment = value.(*simulator.Environment)
	case simulator.ENVIRONMENT_COMPLETE:
	case simulator.SIMULATION_START:
		sim.DataLogOutput("Generation,Environment,# Organisms Mutated,B/D Ratio,Fitness\n")
	case simulator.ORGANISM_MUTATIONS_BENEFICIAL:
		if value.(bool) {
			beneficialCount += 1
		} else {
			deleteriousCount += 1
		}
	case simulator.ITERATION_COMPLETE:
		var ftotal float32 = 0
		for _, o := range sim.Organisms {
			ftotal += o.FitnessForEnvironment(sim.Environment)
		}
		favg := ftotal / float32(len(sim.Organisms))
		sim.DataLogOutput(fmt.Sprintf(
			"%d,%d,%d,%f,%f\n",
			currentTime,
			currentEnvironment.EnvironmentId,
			beneficialCount+deleteriousCount,
			float32(beneficialCount)/float32(deleteriousCount),
			favg))
		currentTime += 1
		beneficialCount = 0
		deleteriousCount = 0
	}
}

// DataLogVerbose is a logging function that prints everything.  Mostly useful for debugging.
func DataLogVerbose(sim *simulator.Simulator, metric simulator.Metric, value interface{}) {
	switch value.(type) {
	case float32:
		sim.Log(fmt.Sprintf("METRIC: %d / %d / %f", sim.Time, metric, value))
	default:
		sim.Log(fmt.Sprintf("METRIC: %d / %d / %+v", sim.Time, metric, value))
	}
}
