package simulator

var currentTime int = 1

type Metric int
const (
	ITERATION_START = iota
	ITERATION_COMPLETE
	ORGANISM_FITNESS_DIFFERENCE
	ORGANISM_MUTATIONS_BENEFICIAL
)

var beneficialCount int = 0
var deleteriousCount int = 0
// DataLogBeneficialMutations is a logging function which gives the beneficial/deleterious ratio.
func DataLogBeneficialMutations(metric Metric, value interface{}) {
	switch metric {
		case ORGANISM_MUTATIONS_BENEFICIAL:
			if value.(bool) {
				beneficialCount += 1
			} else {
				deleteriousCount += 1
			}
		case ITERATION_COMPLETE:
			Log("B/D Ratio: %f, Total: %d\n", float32(beneficialCount) / float32(deleteriousCount), beneficialCount + deleteriousCount)
			currentTime += 1
			beneficialCount = 0
			deleteriousCount = 0
	}
}

// DataLogVerbose is a logging function that prints everything.  Mostly useful for debugging.
func DataLogVerbose(metric Metric, value interface{}) {
	if DataContext != nil {
		switch value.(type) {
			case float32:
				Log("METRIC: %d / %d / %f\n", DataContext.Time, metric, value)
			default:
				Log("METRIC: %d / %s / %+v\n", DataContext.Time, metric, value)
		}
	}
}
