package main

import (
	"flag"

	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
)

type Config struct {
	DataFile          string
	Seed              int64
	Loci              int
	StartingOrganisms int
	MaxOrganisms      int
	Iterations        int
	Environments      int
	Mutability        float64
	NeutralRange      float64
	MaxFitness        float64
	Quiet             bool
}

func NewConfig() *Config {
	return &Config{
		Seed:              0,
		Loci:              10,
		StartingOrganisms: 100,
		MaxOrganisms:      10000,
		Iterations:        100,
		Environments:      5,
		Mutability:        float64(simulator.DEFAULT_MUTABILITY),
		MaxFitness:        float64(simulator.MAX_FITNESS),
		NeutralRange:      0.0,
	}
}

func ParseFlags(config *Config) {
	flag.Float64Var(&config.NeutralRange, "neutral-range", config.NeutralRange, "This is the fitness range that is allowed for a mutation to be considered 'neutral'")
	flag.BoolVar(&config.Quiet, "quiet", config.Quiet, "Limits output sent to STDOUT")
	flag.StringVar(&config.DataFile, "datafile", "", "Sends the output to a file")
	flag.Int64Var(&config.Seed, "seed", config.Seed, "Sets the random number generator seed")
	flag.IntVar(&config.Loci, "loci", config.Loci, "Sets the number of loci in the implicit genome")
	flag.IntVar(&config.Environments, "envs", config.Environments, "Sets the number of environments to use")
	flag.IntVar(&config.Iterations, "iterations", config.Iterations, "Sets the number of iterations for each environment")
	flag.IntVar(&config.MaxOrganisms, "maxorgs", config.MaxOrganisms, "Sets the maximum number of organisms in each generation.  This is a fuzzy maximum due to the culling process.")
	flag.IntVar(&config.StartingOrganisms, "startorgs", config.StartingOrganisms, "Sets the starting number of organisms for the simulation")
	flag.Float64Var(&config.Mutability, "mutability", config.Mutability, "The per-locus mutation rate for the organism (0.1 means that each locus will mutate 10% of the time")
	flag.Float64Var(&config.MaxFitness, "max-fitness", config.MaxFitness, "Set the maximum fitness of a locus")
	flag.Parse()
}
