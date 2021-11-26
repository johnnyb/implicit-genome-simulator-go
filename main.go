package main

import (
	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
	"time"
	"math/rand"
)

func main() {
	// Create simulator (10 loci, 100 organisms)
	sim := simulator.NewSimulator(10, 100)
	sim.MaxOrganisms = 1000000

	// Logging setup
	simulator.DataContext = sim
	simulator.DataLog = simulator.DataLogBeneficialMutations

	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		// Run the simulation for X iterations
		sim.PerformIterations(50)

		// Change environment
		newEnv := simulator.NewEnvironment(sim.ImplicitGenome)
		sim.Environment = newEnv

		simulator.Log("ENVIRONMENT CHANGE")
	}
}
