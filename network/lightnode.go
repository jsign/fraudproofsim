package network

// Warning: Very verbose implementation.
// An resource-optimized implementation might generate the shares to ask on the fly,
// since light-node may be ask for cancellation of requests.
// Less CPU, less mem.

import (
	"errors"
	"math/rand"
)

// LightNodeConfig represents configuration of the light-node
type LightNodeConfig struct {
	K int
	S int
}

// LightNode represents a light-node in the network
type LightNode struct {
	nextShareToAskIdx int
	sharesToAsk       []ShareRequest
	fullNode          *FullNode
}

// See docs for justification about this implementation
func randomPermutations(r *rand.Rand, howMany int, matrixDimensionSize int) []int {
	cache := make(map[int32]bool, howMany)
	res := make([]int, howMany)
	resIdx := 0
	size := int32(matrixDimensionSize * matrixDimensionSize)

	for howMany > 0 {
		candidate := r.Int31n(size)
		if _, ok := cache[candidate]; !ok {
			cache[candidate] = true
			res[resIdx] = int(candidate)
			resIdx++
			howMany--
		}
	}

	return res
}

// NewLightNode returns a fresh light-node
func NewLightNode(r *rand.Rand, config *LightNodeConfig) *LightNode {
	matrixDimensionSize := 2 * config.K

	flattenedPermutationIndexes := randomPermutations(r, config.S, matrixDimensionSize)
	sharesToAsk := make([]ShareRequest, config.S)
	for i := 0; i < config.S; i++ {
		flattenedIndex := flattenedPermutationIndexes[i]

		x := flattenedIndex / matrixDimensionSize
		y := flattenedIndex % matrixDimensionSize
		sharesToAsk[i] = ShareRequest{x, y}
	}

	return &LightNode{0, sharesToAsk, nil}
}

// Connect the light-node to a full-node
func (ln *LightNode) Connect(fn *FullNode) error {
	if ln.fullNode != nil {
		return errors.New("Already connected to a full-node")
	}

	ln.fullNode = fn

	return nil
}

// Run ask for a set of shares to the full-node
func (ln *LightNode) Run() (bool, bool) {
	finished := false
	gotRejection := false
	for !finished && !gotRejection {
		gotRejection, finished = ln.AskForNextShare()
	}

	return gotRejection, finished
}

func (ln *LightNode) AskForNextShare() (bool, bool) {
	shareReceived := ln.fullNode.GetShare(&ln.sharesToAsk[ln.nextShareToAskIdx])
	ln.nextShareToAskIdx++

	return !shareReceived, ln.nextShareToAskIdx == len(ln.sharesToAsk)
}
