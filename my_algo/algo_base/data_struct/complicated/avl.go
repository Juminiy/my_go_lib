package complicated

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/simple"
	"reflect"
)

const (
	immutableBalanceFactor = 1
	maxBin                 = 31
	maxAvlNodeAmount       = 1 << maxBin
)

// 尽量不使用递归，效率低

type avlNode struct {
	Bf          int
	Left, Right *avlNode
	Value       interface{}
}

type AvlTree struct {
	NodeAmount int
	AvlRoot    *avlNode
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

// llRotate ll型 向右转
func llRotate(node *avlNode) *avlNode {
	lNode := node.Left
	node.Left = lNode.Right
	lNode.Right = node
	return lNode
}

// rrRotate rr型 向左转
func rrRotate(node *avlNode) *avlNode {
	rNode := node.Right
	node.Right = rNode.Left
	rNode.Left = node
	return rNode
}

// lrRotate lr型
func lrRotate(node *avlNode) *avlNode {
	node.Left = rrRotate(node.Left)
	node = llRotate(node)
	return node
}

// rlRotate rl型
func rlRotate(node *avlNode) *avlNode {
	node.Right = llRotate(node.Right)
	node = rrRotate(node)
	return node
}

func Comp(compedValue, compValue interface{}) bool {
	if reflect.TypeOf(compedValue).String() == "int" && reflect.TypeOf(compValue).String() == "int" {
		return compedValue.(int) < compValue.(int)
	} else {
		return true
	}
}

func insertNode(root *avlNode, value interface{}) *avlNode {
	if root == nil {
		return &avlNode{Value: value}
	}
	if Comp(value, root.Value) {

	}
	return nil
}
func deleteNode() {

}

// DfsAvl midOrder arr in sequence
func (avl *AvlTree) DfsAvl() []interface{} {
	root := avl.AvlRoot
	stack := simple.MyStack{}
	seq := make([]interface{}, 0)
	for !stack.IsEmpty() || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		if !stack.IsEmpty() {
			tNode, _ := stack.Top()
			stack.Pop()
			root = tNode.(*avlNode)
			seq = append(seq, root.Value)
			root = root.Right
		}
	}
	return seq
}
func (avl *AvlTree) BfsAvl() []interface{} {
	root := avl.AvlRoot
	queue := simple.MyQueue{}
	queue.Push(root)
	seq := make([]interface{}, 0)
	for !queue.IsEmpty() {
		tNode, _ := queue.Front()
		queue.Pop()
		seq = append(seq, tNode.(*avlNode).Value)
		if lNode := tNode.(*avlNode).Left; lNode != nil {
			queue.Push(lNode)
		}
		if rNode := tNode.(*avlNode).Right; rNode != nil {
			queue.Push(rNode)
		}
	}
	return seq
}

func TestInternal() {
	n1 := &avlNode{Value: 1}
	n2 := &avlNode{Left: n1, Value: 2}
	n3 := &avlNode{Left: n2, Value: 3}
	fmt.Println(n3.Left.Value, n3.Left.Left.Value)
	n3 = llRotate(n3)
	fmt.Println(n3.Left.Value, n3.Right.Value)
}
