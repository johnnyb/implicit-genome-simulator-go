# Implicit Genome Simulator

This is a population dynamics simulator based on the book "The Implicit Genome" by Caporale.
The idea is that genetic mutations are not as random as previously thought, but instead genomes have an implicit range within they normally mutate.
This is a population dynamics simulator built on that concept.

This is also meant to be a fairly pluggable simulator.
You can modify the loggers, the PRNG, and other important components at will without having to modify the main code.

## Running

To run the app, after building, do `./implicit-genome-simulator-go`. 
The app has a number of options which are shown by adding `--help`.

## Customizing

You can customize this heavily by providing your own `main()` function.  
The `main` package only parses command line options and sets up the simulation.
Below is a simple `main()` function you can work from.

```
package main

import (
	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
	"time"
	"math/rand"
)

func main() {
	// Create simulator (10 loci, 100 organisms)
	sim := simulator.NewSimulator(10, 100, simulator.DEFAULT_MUTABILITY)

	// Logging setup
	simulator.DataContext = sim
	simulator.DataLog = simulator.DataLogBeneficialMutations

	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())

	// Mostly for logging setup
	sim.Initialize()

	// Run the simulation for X iterations
	sim.PerformIterations(1000)

	// Mostly for logging
	sim.Finish()
}

```

## Replicating Results

Note that seeding the PRNG is an "official" step, and is not done for you by the simulator.
The reason for this is that it allows you to have more reproducible results.
When testing, you should always seed the PRNG.
However, when doing official data gathering, leaving the PRNG unseeded (or with a fixed seed) will allow your data to be deterministically recreated.

When reporting data, you should always report (a) the commit hash, (b) the PRNG seed, and (c) the command-line options.
That should allows others to exactly reproduce the data.

You can also do a true random number generator by setting `simulator.RandomInt` and `simulator.RandomFloat` to be your own functions.

You can do a basic run by just doing:
```
go build
./implicit-genome-simulator-go
```
