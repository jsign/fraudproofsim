package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jsign/fraudproofsimulation/simulator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(solveCmd)
}

var solveCmd = &cobra.Command{
	Use:   "solve [k] [s] [p] [threshold?]",
	Short: "Solves c for k, s and p",
	Long:  `It solves c for k, s and p (p, within a threshold)`,
	Args:  cobra.RangeArgs(3, 4),
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		numIterations, _ := flags.GetInt("n")
		enhanced, _ := flags.GetBool("enhanced")
		simulator := simulator.New(numIterations, enhanced)

		k, _ := strconv.Atoi(args[0])
		s, _ := strconv.Atoi(args[1])
		p, _ := strconv.ParseFloat(args[2], 64)
		threshold := 1 - p
		if len(args) == 4 {
			threshold, _ = strconv.ParseFloat(args[3], 64)
		}

		fmt.Printf("Solving for (k:%v, s:%v, p:%v, threshold:%v)\n", k, s, p, threshold)
		start := time.Now()
		c, p := simulator.Solve(k, s, p, threshold)
		elapsed := time.Since(start) / time.Millisecond
		fmt.Printf("Solution c=%v with p=%v (%dms)", c, p, elapsed)
	},
}
