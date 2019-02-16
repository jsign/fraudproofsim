package simulator

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/jsign/fraudproofsimulation/network"
)

// Simulator runs simulations
type Simulator struct {
	numIterations   int
	enhancedModel   bool
	simulationsRand []*rand.Rand
}

// New generates a new simulator
func New(numIterations int, enhancedModel bool) *Simulator {
	simulationsRand := generateRandomSeeds(numIterations)

	return &Simulator{numIterations, enhancedModel, simulationsRand}
}

func generateRandomSeeds(num int) []*rand.Rand {
	simulationsRand := make([]*rand.Rand, num)
	for i := 0; i < num; i++ {
		simulationsRand[i] = rand.New(rand.NewSource(rand.Int63()))
	}

	return simulationsRand
}

// Solve for (k, s, p, threshold)
func (sm *Simulator) Solve(k int, s int, p float64, threshold float64) (int, float64) {
	max := 4 * k * k
	min := 1
	actual := (max + min) / 2
	q := float64(0)
	config := &SimulationConfig{K: k, S: s, NumLN: actual}
	quit := false

	for !quit && max > min {
		config.NumLN = actual
		q, _ = sm.Run(config.K, config.S, config.NumLN)

		probdiff := q - p
		fmt.Printf("[%v, %v]: c=%v p=%v\n", min, max, actual, q)

		if math.Abs(probdiff) < threshold {
			quit = true
		} else if probdiff < 0 {
			min = actual
			actual = (max + min) / 2
		} else {
			max = actual
			actual = (max + min) / 2
		}
	}

	return actual, q
}

// RunConfigs will run the simulator for the specified configuration
func (sm *Simulator) RunConfigs(configs []SimulationConfig) {
	for _, c := range configs {
		start := time.Now()
		p, _ := sm.Run(c.K, c.S, c.NumLN)
		elapsed := time.Since(start) / time.Millisecond
		fmt.Printf("k=%v, s=%v, c=%v => p=%v %dms\n", c.K, c.S, c.NumLN, p, elapsed)
	}
}

// Run runs a specific configuration
func (sm *Simulator) Run(k int, s int, c int) (float64, float64) {
	fnConfig := &network.FullNodeConfig{K: k}
	lnConfig := &network.LightNodeConfig{K: k, S: s}

	successSimulationCount := 0
	sscMux := &sync.Mutex{}
	lnCompleted := 0
	lnCompletedMux := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(sm.numIterations)
	for i := 0; i < sm.numIterations; i++ {
		if sm.enhancedModel {
			go simulationIterationEnhancedModel(sm.simulationsRand[i], wg, c, fnConfig, lnConfig, &successSimulationCount, &lnCompleted, sscMux, lnCompletedMux)
		} else {
			go simulationIterationStandardModel(sm.simulationsRand[i], wg, c, fnConfig, lnConfig, &successSimulationCount, &lnCompleted, sscMux, lnCompletedMux)
		}
	}
	wg.Wait()

	return float64(successSimulationCount) / float64(sm.numIterations), float64(lnCompleted) / float64(sm.numIterations)
}

func simulationIterationEnhancedModel(r *rand.Rand, wg *sync.WaitGroup, numLightNodes int, fnCondif *network.FullNodeConfig, lnConfig *network.LightNodeConfig, successSimulationCount *int, lnCompleted *int, sscMux *sync.Mutex, lnCompletedMux *sync.Mutex) {
	fullNode := network.NewFullNode(fnCondif)

	lightNodes := make([]*network.LightNode, numLightNodes)
	for i := 0; i < numLightNodes; i++ {
		lightNodes[i] = network.NewLightNode(r, lnConfig)
		lightNodes[i].Connect(fullNode)
	}

	somebodyRequestRejected := false
	finishedLN := make(map[int32]bool)
	for !somebodyRequestRejected && len(finishedLN) < numLightNodes {
		electedLN := r.Int31n(int32(numLightNodes))
		if _, ok := finishedLN[electedLN]; !ok {
			lightNode := lightNodes[electedLN]

			noMoreSharesToAsk := false
			somebodyRequestRejected, noMoreSharesToAsk = lightNode.AskForNextShare()

			if noMoreSharesToAsk {
				finishedLN[electedLN] = true
			}
		}
	}

	if somebodyRequestRejected {
		sscMux.Lock()
		*successSimulationCount++
		sscMux.Unlock()
	}

	lnCompletedMux.Lock()
	*lnCompleted += len(finishedLN)
	lnCompletedMux.Unlock()

	wg.Done()
}

func simulationIterationStandardModel(r *rand.Rand, wg *sync.WaitGroup, numLightNodes int, fnCondif *network.FullNodeConfig, lnConfig *network.LightNodeConfig, successSimulationCount *int, lnCompleted *int, sscMux *sync.Mutex, lnCompletedMux *sync.Mutex) {
	fullNode := network.NewFullNode(fnCondif)

	somebodyRequestRejected := false
	var j int
	for j = 1; j <= numLightNodes && !somebodyRequestRejected; j++ {
		lightNode := network.NewLightNode(r, lnConfig)
		lightNode.Connect(fullNode)

		somebodyRequestRejected, _ = lightNode.Run()
	}

	if somebodyRequestRejected {
		sscMux.Lock()
		*successSimulationCount++
		sscMux.Unlock()
	}

	lnCompletedMux.Lock()
	*lnCompleted += j
	lnCompletedMux.Unlock()

	wg.Done()
}
