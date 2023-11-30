package datalogging

import (
	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
)

type DataLogger func(sim *simulator.Simulator, metric simulator.Metric, value interface{})

func DataLogTee(loggers ...DataLogger) DataLogger {
	return func(sim *simulator.Simulator, metric simulator.Metric, value interface{}) {
		for _, logger := range loggers {
			logger(sim, metric, value)
		}
	}
}
