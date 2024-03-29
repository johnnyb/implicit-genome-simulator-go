package simulator

type Metric int

const (
	SIMULATION_START = iota
	SIMULATION_COMPLETE
	ITERATION_START
	ITERATION_COMPLETE
	ORGANISM_FITNESS_DIFFERENCE
	ORGANISM_MUTATIONS_BENEFICIAL
	ENVIRONMENT_START
	ENVIRONMENT_COMPLETE
)
