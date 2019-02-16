package network

// Config represents configurations for the full-node
type FullNodeConfig struct {
	K int
}

// FullNode represents a full-node in the network
type FullNode struct {
	matrixDimension   int
	maxSharesCanShare int

	exposedShares map[int]bool
}

// NewFullNode returns a new full-node configured as specified
func NewFullNode(conf *FullNodeConfig) *FullNode {
	matrixDimension := 2 * conf.K
	maxSharesCanShare := matrixDimension*matrixDimension - (conf.K+1)*(conf.K+1)

	exposedShares := make(map[int]bool)

	return &FullNode{matrixDimension, maxSharesCanShare, exposedShares}
}

// GetShare returns a boolean that represent that the shared was returned
func (fn *FullNode) GetShare(rq *ShareRequest) bool {
	flattenedIndex := rq.X*fn.matrixDimension + rq.Y
	if _, ok := fn.exposedShares[flattenedIndex]; ok {
		return true
	}

	fn.exposedShares[flattenedIndex] = true

	return len(fn.exposedShares) < fn.maxSharesCanShare
}
