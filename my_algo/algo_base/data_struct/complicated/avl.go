package complicated

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/simple"
	"math"
	"reflect"
)

const (
	immutableBalanceFactor    = 1
	outOfBoundsBalanceFactor1 = 2
	outOfBoundsBalanceFactor2 = -2
	maxBin                    = 31
	maxAvlNodeAmount          = 1 << maxBin
	EQ                        = 10000
	GT                        = 10001
	LT                        = 10002
	QS                        = 10003
	INF                       = math.MaxInt
)

// 尽量不使用递归，效率低
// 树高的问题不能有效解决
type avlNode struct {
	Left, Right *avlNode    // Left Node Ptr, Right Node Ptr
	Value       interface{} // Value interface{}
	Count       int         // Same Value node in a node, Count
	Height      int         // Node Height
}

type AvlTree struct {
	NodeAmount int      // Tree Total Node Counts, NodeAmount
	AvlRoot    *avlNode // Root of Tree, AvlRoot
}

// for realTime calculation
func nodeDepth(node *avlNode) int {
	if node == nil {
		return 0
	}
	return algo_basic.MaxValue(nodeDepth(node.Left), nodeDepth(node.Right)) + 1
}

func balanceFactor(node *avlNode) int {
	if node == nil {
		return INF
	}
	lHeight, rHeight := 0, 0
	if node.Left != nil {
		lHeight = node.Left.Height
	}
	if node.Right != nil {
		rHeight = node.Right.Height
	}
	return lHeight - rHeight
}

// llRotate ll型 向右转
func llRotate(node *avlNode) *avlNode {
	lNode := node.Left
	node.Left = lNode.Right
	node.Height = lNode.Height - 1
	lNode.Right = node
	return lNode
}

// rrRotate rr型 向左转
func rrRotate(node *avlNode) *avlNode {
	rNode := node.Right
	node.Right = rNode.Left
	node.Height = rNode.Height - 1
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
func IntType(value interface{}) bool {
	if reflect.TypeOf(value).String() == "int" ||
		reflect.TypeOf(value).String() == "int32" ||
		reflect.TypeOf(value).String() == "int64" {
		return true
	} else {
		return false
	}
}

func Comp(compedValue, compValue interface{}) int {
	if IntType(compedValue) && IntType(compValue) {
		compedInt, compInt := compedValue.(int), compValue.(int)
		if compedInt == compInt {
			return EQ
		} else if compedInt > compInt {
			return GT
		} else {
			return LT
		}
	} else {
		return QS
	}
}

func insertNode(root *avlNode, value interface{}) *avlNode {
	if root == nil {
		return &avlNode{Value: value, Count: 1, Height: 1}
	}
	compRes := Comp(value, root.Value)

	switch compRes {
	case EQ:
		{
			root.Count++
		}
	case GT:
		{
			root.Right = insertNode(root.Right, value)
			root.Height++
			if balanceFactor(root) == outOfBoundsBalanceFactor2 {
				if balanceFactor(root.Right) == -1 {
					root = rrRotate(root)
				} else {
					root = lrRotate(root)
				}
			}
		}
	case LT:
		{
			root.Left = insertNode(root.Left, value)
			root.Height++
			if balanceFactor(root) == outOfBoundsBalanceFactor1 {
				if balanceFactor(root.Left) == 1 {
					root = llRotate(root)
				} else {
					root = lrRotate(root)
				}
			}
		}
	}
	return root
}
func deleteNode(root *avlNode, value interface{}) *avlNode {
	return nil
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
			for i := 0; i < root.Count; i++ {
				seq = append(seq, root.Value)
			}
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
		for i := 0; i < root.Count; i++ {
			seq = append(seq, tNode.(*avlNode).Value)
		}
		if lNode := tNode.(*avlNode).Left; lNode != nil {
			queue.Push(lNode)
		}
		if rNode := tNode.(*avlNode).Right; rNode != nil {
			queue.Push(rNode)
		}
	}
	return seq
}

func (avl *AvlTree) Insert(value interface{}) {
	avl.AvlRoot = insertNode(avl.AvlRoot, value)
	avl.NodeAmount++
}
func (avl *AvlTree) Delete(value interface{}) {
	avl.AvlRoot = deleteNode(avl.AvlRoot, value)
}

func TestInternal() {
	n1 := &avlNode{Value: 1}
	n2 := &avlNode{Left: n1, Value: 2}
	n3 := &avlNode{Left: n2, Value: 3}
	fmt.Println(n3.Left.Value, n3.Left.Left.Value)
	n3 = llRotate(n3)
	fmt.Println(n3.Left.Value, n3.Right.Value)
}
