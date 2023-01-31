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
	}
	simulator.Seed(config.Seed)

	// Create simulator (10 loci, 100 organisms)
	// sim := simulator.NewSimulator(10, 100, simulator.DEFAULT_MUTABILITY)
	sim := simulator.NewSimulator(config.Loci, config.StartingOrganisms, config.Mutability)
	if config.DataFile == "" {
		sim.DataStream = os.Stdout
	} else {
		sim.DataStream, err = os.OpenFile(config.DataFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	}
	sim.MaxOrganisms = config.MaxOrganisms
	sim.DataLogger = datalogging.DataLogBeneficialMutations

	sim.Log(fmt.Sprintf("Started with seed: %d", config.Seed))

	if config.Quiet {
		sim.Logger = func(sim *simulator.Simulator, message string) {}
	}

	sim.Log("**** IGENOME ****")
	sim.Log(sim.ImplicitGenome.String())

	sim.Initialize()

	// Generate environments at the beginning, so that even
	// with different options you will get the same environments
	// if run with the same seed.
	environs := [](*simulator.Environment){}
	for i := 0; i < config.Environments; i++ {
		environs = append(environs, simulator.NewEnvironment(sim.ImplicitGenome))
	}

	for _, environ := range environs {
		sim.SetEnvironment(environ)

		// Report environment
		sim.Log("**** ENVIRONMENT ****")
		sim.Log(sim.Environment.String())

		// Run the simulation for X iterations
		sim.PerformIterations(config.Iterations)
	}
	sim.Finish()
}
