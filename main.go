package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnnyb/implicit-genome-simulator-go/datalogging"

	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
)

func main() {
	var err error
	config := NewConfig()
	ParseFlags(config)

	// Seed the PRNG
	if config.Seed == 0 {
		config.Seed = time.Now().UnixNano()
		simulator.Seed(config.Seed)
	}

	// Create simulator (10 loci, 100 organisms)
	// sim := simulator.NewSimulator(10, 100, simulator.DEFAULT_MUTABILITY)
	sim := simulator.NewSimulator(config.Loci, config.StartingOrganisms, config.Mutability)
	if config.DataFile == "" {
		sim.DataStream = os.Stdout
	} else {
		sim.DataStream, err = os.OpenFile(config.DataFile, os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
	}
	sim.MaxOrganisms = config.MaxOrganisms
	sim.DataLogger = datalogging.DataLogBeneficialMutations

	sim.Log(fmt.Sprintf("Started with seed: %d", config.Seed))

	sim.Log("**** IGENOME ****")
	sim.Log(sim.ImplicitGenome.String())

	sim.Initialize()
	sim.SetEnvironment(simulator.NewEnvironment(sim.ImplicitGenome))

	for i := 0; i < config.Environments; i++ {
		// Report environment
		sim.Log("**** ENVIRONMENT ****")
		sim.Log(sim.Environment.String())

		// Run the simulation for X iterations
		sim.PerformIterations(config.Iterations)

		// Create New Environment
		sim.SetEnvironment(simulator.NewEnvironment(sim.ImplicitGenome))
	}
	sim.Finish()
}
