package data_struct

import (
	"fmt"
	"testing"
)

func TestMyStack_IsEmpty(t *testing.T) {
	var stack MyStack
	fmt.Println(stack.IsEmpty())
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.Top())
}

func TestRev(t *testing.T) {
	node1 := LinkedList{1, nil}
	node2 := LinkedList{2, &node1}
	node3 := LinkedList{3, &node2}
	TraversalList(Rev(&node3))
}
