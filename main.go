package main

import (
	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
	"time"
	"math/rand"
)

func main() {
	// Create simulator (10 loci, 100 organisms)
	sim := simulator.NewSimulator(10, 100)

	// Logging setup
	simulator.DataContext = sim
	simulator.DataLog = simulator.DataLogBeneficialMutations

	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())

	// Run the simulation for X iterations
	sim.PerformIterations(1000)
}
