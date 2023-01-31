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
	Mutability        float32
}

func NewConfig() *Config {
	return &Config{
		Seed:              0,
		Loci:              10,
		StartingOrganisms: 100,
		MaxOrganisms:      10000,
		Iterations:        100,
		Environments:      5,
		Mutability:        simulator.DEFAULT_MUTABILITY,
	}
}

func ParseFlags(config *Config) {
	mutability := float64(config.Mutability)
	flag.StringVar(&config.DataFile, "datafile", "", "Sends the output to a file")
	flag.Int64Var(&config.Seed, "seed", config.Seed, "Sets the random number generator seed")
	flag.IntVar(&config.Loci, "loci", config.Loci, "Sets the number of loci in the implicit genome")
	flag.IntVar(&config.Environments, "envs", config.Environments, "Sets the number of environments to use")
	flag.IntVar(&config.Iterations, "iterations", config.Iterations, "Sets the number of iterations for each environment")
	flag.IntVar(&config.MaxOrganisms, "maxorgs", config.MaxOrganisms, "Sets the maximum number of organisms in each generation.  This is a fuzzy maximum due to the culling process.")
	flag.IntVar(&config.StartingOrganisms, "startorgs", config.StartingOrganisms, "Sets the starting number of organisms for the simulation")
	flag.Float64Var(&mutability, "mutability", mutability, "The per-locus mutation rate for the organism (0.1 means that each locus will mutate 10% of the time")
	flag.Parse()
}
