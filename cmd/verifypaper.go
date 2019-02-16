package cmd

import (
	"fmt"
	"time"

	"github.com/jsign/fraudproofsimulation/simulator"
	"github.com/spf13/cobra"
)

var (
	configs = []simulator.SimulationConfig{
		simulator.SimulationConfig{K: 16, S: 50, NumLN: 28},
		simulator.SimulationConfig{K: 16, S: 20, NumLN: 69},
		simulator.SimulationConfig{K: 16, S: 10, NumLN: 138},
		simulator.SimulationConfig{K: 16, S: 5, NumLN: 275},
		simulator.SimulationConfig{K: 16, S: 2, NumLN: 690},

		simulator.SimulationConfig{K: 32, S: 50, NumLN: 112},
		simulator.SimulationConfig{K: 32, S: 20, NumLN: 280},
		simulator.SimulationConfig{K: 32, S: 10, NumLN: 561},
		simulator.SimulationConfig{K: 32, S: 5, NumLN: 1122},
		simulator.SimulationConfig{K: 32, S: 2, NumLN: 2805},

		simulator.SimulationConfig{K: 64, S: 50, NumLN: 451},
		simulator.SimulationConfig{K: 64, S: 20, NumLN: 1129},
		simulator.SimulationConfig{K: 64, S: 10, NumLN: 2258},
		simulator.SimulationConfig{K: 64, S: 5, NumLN: 4516},
		simulator.SimulationConfig{K: 64, S: 2, NumLN: 11289},

		simulator.SimulationConfig{K: 128, S: 50, NumLN: 1811},
		simulator.SimulationConfig{K: 128, S: 20, NumLN: 4500},
		simulator.SimulationConfig{K: 128, S: 10, NumLN: 9000},
		simulator.SimulationConfig{K: 128, S: 5, NumLN: 18000},
		simulator.SimulationConfig{K: 128, S: 2, NumLN: 40000},
	}
)

func init() {
	rootCmd.AddCommand(verifypaperCmd)
}

var verifypaperCmd = &cobra.Command{
	Use:   "verifypaper",
	Short: "Verifies setups calculated in the paper",
	Long:  `It runs simulations for all the configurations mentioned in the paper in order to verify p.`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		numIterations, _ := flags.GetInt("n")
		enhanced, _ := flags.GetBool("enhanced")
		simulator := simulator.New(numIterations, enhanced)

		start := time.Now()
		simulator.RunConfigs(configs)
		elapsed := time.Since(start) / time.Second
		fmt.Printf("Total time %ds", elapsed)
	},
}
