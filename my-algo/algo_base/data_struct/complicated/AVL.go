package complicated

import "my_algo/algo_base/algo_basic"

const (
	immutableBalanceFactor = 1
	maxBin                 = 31
	maxAVLNodeAmount       = 1 << maxBin
)

type avlNode struct {
	bf          int
	left, right *avlNode
	value       interface{}
}

// for realTime calculation
func nodeDepth(node *avlNode) int {
	if node == nil {
		return 0
	}
	return algo_basic.MaxValue(nodeDepth(node.left), nodeDepth(node.right)) + 1
}

func (node *avlNode) balanceFactor() {
	node.bf = nodeDepth(node.left) - nodeDepth(node.right)
}

// ll型
func (node *avlNode) llRotate() {

}
