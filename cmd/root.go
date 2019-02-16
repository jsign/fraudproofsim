package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	numIterations int
	enhancedModel bool
)

func init() {
	rootCmd.PersistentFlags().IntVar(&numIterations, "n", 500, "number of iterations to run per instance")
	rootCmd.PersistentFlags().BoolVar(&enhancedModel, "enhanced", false, "run an Enhanced Model")
}

var rootCmd = &cobra.Command{
	Use:   "fraudproofsim",
	Short: "fraudproofsim is a tool for simulating a fraud-proof network",
	Long:  "It permits to compare, solve and verify fraud-proof networks.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
