package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/jsign/fraudproofsimulation/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Will produce non-reproducible results.
	// Delete or provide deterministic seed if required.
	rand.Seed(time.Now().UTC().UnixNano())

	cmd.Execute()
}
