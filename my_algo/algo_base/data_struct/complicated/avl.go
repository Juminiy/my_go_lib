package complicated

import (
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic"
)

const (
	immutableBalanceFactor = 1
	maxBin                 = 31
	maxAVLNodeAmount       = 1 << maxBin
)

// 尽量不使用递归，效率低

type AvlNode struct {
	Bf          int
	Left, Right *AvlNode
	Value       interface{}
}

type AvlTree struct {
	NodeAmount int
	AvlRoot    *AvlNode
}

// for realTime calculation
func nodeDepth(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return algo_basic.MaxValue(nodeDepth(node.Left), nodeDepth(node.Right)) + 1
}

func (node *AvlNode) balanceFactor() {
	node.Bf = nodeDepth(node.Left) - nodeDepth(node.Right)
}

// llRotate ll型 向右转
func llRotate(node *AvlNode) *AvlNode {
	lNode := node.Left
	node.Left = lNode.Right
	lNode.Right = node
	return lNode
}

// rrRotate rr型 向左转
func rrRotate(node *AvlNode) *AvlNode {
	rNode := node.Right
	node.Right = rNode.Left
	rNode.Left = node
	return rNode
}

// lrRotate lr型
func lrRotate(node *AvlNode) *AvlNode {
	node.Left = rrRotate(node.Left)
	node = llRotate(node)
	return node
}

// rlRotate rl型
func rlRotate(node *AvlNode) *AvlNode {
	node.Right = llRotate(node.Right)
	node = rrRotate(node)
	return node
}

func Insert() {

}
func Delete() {

}
func (avl *AvlTree) DfsAvl() {

}
func (avl *AvlTree) BfsAvl() {

}
