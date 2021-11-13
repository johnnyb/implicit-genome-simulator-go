# Implicit Genome Simulator

This is a population dynamics simulator based on the book "The Implicit Genome" by Caporale.
The idea is that genetic mutations are not as random as previously thought, but instead genomes have an implicit range within they normally mutate.
This is a population dynamics simulator built on that concept.

This is also meant to be a fairly pluggable simulator.
You can modify the loggers, the PRNG, and other important components at will without having to modify the main code.

## Running

A simple example of how to run the simulator is provided in `main.go`, and is reproduced here:

```
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

```

Note that seeding the PRNG is an "official" step, and is not done for you by the simulator.
The reason for this is that it allows you to have more reproducible results.
When testing, you should always seed the PRNG.
However, when doing official data gathering, leaving the PRNG unseeded (or with a fixed seed) will allow your data to be deterministically recreated.

You can also do a true random number generator by setting `simulator.RandomInt` and `simulator.RandomFloat` to be your own functions.

You can do a basic run by just doing:
```
go build
./implicit-genome-simulator-go
```
