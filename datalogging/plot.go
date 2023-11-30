package datalogging

import (
	"image/color"
	"os"

	"github.com/johnnyb/implicit-genome-simulator-go/simulator"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func NewDataLogPlotter(plotFile string) DataLogger {
	currentTime := 1.0
	beneficialCount := 0.0
	deleteriousCount := 0.0
	currentEnvironmentID := int32(-1)

	time := []float64{}
	BDRatio := []float64{}
	fitness := []float64{}
	environmentChangeIdx := []int{}

	// DataLogBeneficialMutations is a logging function which gives the beneficial/deleterious ratio.
	return func(sim *simulator.Simulator, metric simulator.Metric, value interface{}) {
		switch metric {
		case simulator.ENVIRONMENT_START:
			currentEnvironment = value.(*simulator.Environment)
			if currentEnvironmentID == -1 {
				currentEnvironmentID = currentEnvironment.EnvironmentId
			}
			if currentEnvironmentID != currentEnvironment.EnvironmentId {
				environmentChangeIdx = append(environmentChangeIdx, len(time))
				currentEnvironmentID = currentEnvironment.EnvironmentId
			}
		case simulator.ORGANISM_MUTATIONS_BENEFICIAL:
			if value.(bool) {
				beneficialCount += 1.0
			} else {
				deleteriousCount += 1.0
			}
		case simulator.ITERATION_COMPLETE:
			time = append(time, currentTime)
			BDRatio = append(BDRatio, beneficialCount/deleteriousCount)

			var ftotal float32 = 0
			for _, o := range sim.Organisms {
				ftotal += o.FitnessForEnvironment(sim.Environment)
			}
			favg := ftotal / float32(len(sim.Organisms))
			fitness = append(fitness, float64(favg))

			currentTime += 1
			beneficialCount = 0
			deleteriousCount = 0
		case simulator.SIMULATION_COMPLETE:
			const rows, cols = 1, 2
			plots := make([][]*plot.Plot, rows)
			for j := 0; j < rows; j++ {
				plots[j] = make([]*plot.Plot, cols)
				for i := 0; i < cols; i++ {
					p := plot.New()

					if i == 0 {
						p.Title.Text = "Beneficial/Deleterious Ratio"
						p.X.Label.Text = "Time"
						p.Y.Label.Text = "B/D"
						pts := make(plotter.XYs, len(time))
						for i := range pts {
							pts[i].X = time[i]
							pts[i].Y = BDRatio[i]
						}
						plotutil.AddLines(p, pts)

						pts = make(plotter.XYs, len(environmentChangeIdx))
						for i, idx := range environmentChangeIdx {
							pts[i].X = time[idx]
							pts[i].Y = BDRatio[idx]
						}
						s, _ := plotter.NewScatter(pts)
						s.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
						s.GlyphStyle.Shape = draw.CircleGlyph{}
						p.Add(s)

					} else {
						p.Title.Text = "Fitness"
						p.X.Label.Text = "Time"
						p.Y.Label.Text = "Fitness"
						pts := make(plotter.XYs, len(time))
						for i := range pts {
							pts[i].X = time[i]
							pts[i].Y = fitness[i]
						}

						plotutil.AddLines(p, pts)

						pts = make(plotter.XYs, len(environmentChangeIdx))
						for i, idx := range environmentChangeIdx {
							pts[i].X = time[idx]
							pts[i].Y = fitness[idx]
						}
						s, _ := plotter.NewScatter(pts)
						s.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
						s.GlyphStyle.Shape = draw.CircleGlyph{}
						p.Add(s)
					}

					p.X.Min = 0
					p.X.Max = currentTime
					plots[j][i] = p
				}
			}

			img := vgimg.New(vg.Points(800), vg.Points(300))
			dc := draw.New(img)

			t := draw.Tiles{
				Rows: rows,
				Cols: cols,
			}

			canvases := plot.Align(plots, t, dc)
			for j := 0; j < rows; j++ {
				for i := 0; i < cols; i++ {
					if plots[j][i] != nil {
						plots[j][i].Draw(canvases[j][i])
					}
				}
			}

			w, err := os.Create(plotFile + ".png")
			if err != nil {
				panic(err)
			}

			png := vgimg.PngCanvas{Canvas: img}
			if _, err := png.WriteTo(w); err != nil {
				panic(err)
			}
		}
	}

}
