package main

import (
	"fmt"
	"time"

	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
)

func main() {
	config := NewConfig()
	ParseFlags(config)

	simulator.DataLog = simulator.DataLogBeneficialMutations

	// Seed the PRNG
	if config.Seed == 0 {
		config.Seed = time.Now().UnixNano()
		simulator.Seed(config.Seed)
	}
	simulator.Log(fmt.Sprintf("Started with seed: %d", config.Seed))

	// Create simulator (10 loci, 100 organisms)
	// sim := simulator.NewSimulator(10, 100, simulator.DEFAULT_MUTABILITY)
	sim := simulator.NewSimulator(config.Loci, config.StartingOrganisms, config.Mutability)
	sim.MaxOrganisms = config.MaxOrganisms
	simulator.DataContext = sim

	simulator.Log("**** IGENOME ****")
	simulator.Log(sim.ImplicitGenome.String())
	for i := 0; i < config.Environments; i++ {
		// Report environment
		simulator.Log("**** ENVIRONMENT ****")
		simulator.Log(sim.Environment.String())

		// Run the simulation for X iterations
		sim.PerformIterations(config.Iterations)

		// Create New Environment
		sim.Environment = simulator.NewEnvironment(sim.ImplicitGenome)
	}
}
