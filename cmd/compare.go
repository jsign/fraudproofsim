package cmd

import (
	"fmt"
	"strconv"

	"github.com/jsign/fraudproofsimulation/simulator"
	"github.com/spf13/cobra"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func init() {
	rootCmd.AddCommand(compareCmd)
}

var compareCmd = &cobra.Command{
	Use:   "compare [k] [s] [#points]",
	Short: "Compares the Standard and Enhanced models",
	Long:  `Compares Standard and Enhanced model to understand their impact on soundness`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		numIterations, _ := flags.GetInt("n")

		simEnhanced := simulator.New(numIterations, true)
		simStandard := simulator.New(numIterations, false)

		k, _ := strconv.Atoi(args[0])
		s, _ := strconv.Atoi(args[1])
		numPoints, _ := strconv.Atoi(args[2])

		fmt.Printf("Solving c for (k: %v, s: %v) with precision .99+-.005:\n", k, s)
		c, _ := simStandard.Solve(k, s, .99, .005)
		start := int(float32(c) * .50)
		end := int(float32(c) * 1.50)
		stepsize := (end - start + 1) / numPoints

		enhancedPoints := make(plotter.XYs, (end-start+1)/stepsize)
		standardPoints := make(plotter.XYs, (end-start+1)/stepsize)
		fmt.Printf("Found solution c=%v, now generating %v points in [.50*c,1.5*c]=[%v, %v]:\n", c, numPoints, start, end)
		for i := start; i <= end; i += stepsize {
			_, b := simStandard.Run(k, s, i)
			standardPoints = append(standardPoints, plotter.XY{X: float64(i), Y: float64(b)})

			_, be := simEnhanced.Run(k, s, i)
			enhancedPoints = append(enhancedPoints, plotter.XY{X: float64(i), Y: float64(be)})
			fmt.Printf("%v%%\n", 100*(i-start)/int((end-start)))
		}

		generatePlot(enhancedPoints, standardPoints)
	},
}

func generatePlot(enhancedPoints plotter.XYs, standardPoints plotter.XYs) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Enhanced vs Standard"
	p.X.Label.Text = "c"
	p.Y.Label.Text = "# Completed light-nodes"

	err = plotutil.AddLinePoints(p,
		"Enhanced", enhancedPoints,
		"Standard", standardPoints)
	if err != nil {
		panic(err)
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}
	fmt.Printf("Plotted in plot.png")
}
