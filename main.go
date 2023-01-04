package main

import (
	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
)

func main() {
	// Seed the PRNG
	simulator.ReseedAndPrint()

	// Create simulator (10 loci, 100 organisms)
	// sim := simulator.NewSimulator(10, 100, simulator.DEFAULT_MUTABILITY)
	sim := simulator.NewSimulator(1, 100, 0.1)
	sim.MaxOrganisms = 10000

	// Logging setup
	simulator.DataContext = sim
	simulator.DataLog = simulator.DataLogBeneficialMutations

	for i := 0; i < 4; i++ {
		// Run the simulation for X iterations
		sim.PerformIterations(10000)

		// Change environment
		newEnv := simulator.NewEnvironment(sim.ImplicitGenome)
		sim.Environment = newEnv

		simulator.Log("**** ENVIRONMENT CHANGE ****")
	}
}
