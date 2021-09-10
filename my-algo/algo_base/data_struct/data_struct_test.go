package data_struct

import (
	"fmt"
	"my_algo/algo_base/data_struct/simple"
	"testing"
)

func TestMyStack_IsEmpty(t *testing.T) {
	var stack simple.MyStack
	fmt.Println(stack.IsEmpty())
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.Top())
}

func TestRev(t *testing.T) {
	node1 := simple.LinkedList{1, nil}
	node2 := simple.LinkedList{2, &node1}
	node3 := simple.LinkedList{3, &node2}
	simple.TraversalList(simple.Rev(&node3))
}

func TestMyQueue_Front(t *testing.T) {
	q := &simple.MyQueue{}
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	for i := 0; i < 5; i++ {
		v, err := q.Front()
		if err == nil {
			fmt.Printf("%v ", v)
		}
		q.Pop()
	}
	for i := 0; i < 5; i++ {
		v, err := q.Front()
		if err == nil {
			fmt.Printf("%v ", v)
		}
		q.Pop()
	}
	fmt.Println(q.IsEmpty())
}

func TestBfs(t *testing.T) {
	node3 := simple.BinaryTree{nil, nil, 4}
	node2 := simple.BinaryTree{nil, nil, 3}
	node1 := simple.BinaryTree{&node3, nil, 2}
	root := &simple.BinaryTree{&node1, &node2, 1}
	fmt.Println(simple.Bfs(root))
}

func TestNonRecursionDfs(t *testing.T) {
	node3 := simple.BinaryTree{nil, nil, 4}
	node2 := simple.BinaryTree{nil, nil, 3}
	node1 := simple.BinaryTree{&node3, nil, 2}
	root := &simple.BinaryTree{&node1, &node2, 1}
	fmt.Println(simple.NonRecursionDfs(root))
}
