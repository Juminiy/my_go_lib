package complicated

import (
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic"
)

const (
	immutableBalanceFactor = 1
	maxBin                 = 31
	maxAVLNodeAmount       = 1 << maxBin
)

type avlNode struct {
	Bf          int
	Left, Right *avlNode
	Value       interface{}
}

// for realTime calculation
func nodeDepth(node *avlNode) int {
	if node == nil {
		return 0
	}
	return algo_basic.MaxValue(nodeDepth(node.Left), nodeDepth(node.Right)) + 1
}

func (node *avlNode) balanceFactor() {
	node.Bf = nodeDepth(node.Left) - nodeDepth(node.Right)
}

// ll型
func (node *avlNode) llRotate() {

}
