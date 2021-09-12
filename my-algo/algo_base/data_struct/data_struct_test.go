package data_struct

import (
	"fmt"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/simple"
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
	node3 := simple.BinaryTree{Value: 4}
	node2 := simple.BinaryTree{Value: 3}
	node1 := simple.BinaryTree{Left: &node3, Value: 2}
	root := &simple.BinaryTree{Left: &node1, Right: &node2, Value: 1}
	fmt.Println(simple.Bfs(root))
}

func TestNonRecursionDfs(t *testing.T) {
	node3 := simple.BinaryTree{Value: 4}
	node2 := simple.BinaryTree{Value: 3}
	node1 := simple.BinaryTree{Left: &node3, Value: 2}
	root := &simple.BinaryTree{Left: &node1, Right: &node2, Value: 1}
	fmt.Println(simple.NonRecursionDfs(root))
}

func TestBfsGraph(t *testing.T) {
	adj := &complicated.AdjGraph{}
	adj.Construct(true)
	nodea, nodeb, nodec, noded, nodee := &complicated.GraphNode{Value: "A"}, &complicated.GraphNode{Value: "B"}, &complicated.GraphNode{Value: "C"}, &complicated.GraphNode{Value: "D"}, &complicated.GraphNode{Value: "E"}
	adj.AddNode(nodea)
	adj.AddNode(nodeb)
	adj.AddNode(nodec)
	adj.AddNode(noded)
	adj.AddNode(nodee)
	edgeab, edgeac, edgead, edgebe := &complicated.GraphEdge{"1"}, &complicated.GraphEdge{"2"}, &complicated.GraphEdge{"3"}, &complicated.GraphEdge{"4"}
	edgeea := &complicated.GraphEdge{5}
	nodeg := &complicated.GraphNode{"G"}
	adj.AddNode(nodeg)
	edgega := &complicated.GraphEdge{8}
	adj.AddEdge(nodeg, nodea, edgega)
	adj.AddEdge(nodea, nodeb, edgeab)
	adj.AddEdge(nodea, nodec, edgeac)
	adj.AddEdge(nodea, noded, edgead)
	adj.AddEdge(nodeb, nodee, edgebe)
	adj.AddEdge(nodee, nodea, edgeea)
	nodef, edgebf := &complicated.GraphNode{"F"}, &complicated.GraphEdge{"7"}
	adj.AddNode(nodef)
	adj.AddEdge(nodeb, nodef, edgebf)
	//fmt.Println("nodea",&nodea);fmt.Println("nodeb",&nodeb);fmt.Println("nodec",&nodec);fmt.Println("noded",&noded);fmt.Println("nodee",&nodee)
	//fmt.Println(adj.Adjacent)
	//fmt.Println(&adj)
	adj.PrintNodes(adj.BfsGraph())
}

func TestUnite(t *testing.T) {
	s1, s2 := &complicated.MySet{}, &complicated.MySet{}
	s1.Construct()
	s2.Construct()
	s1.Insert("121")
	s1.Insert('2')
	s1.Insert('3')
	s2.Insert("121")
	s2.Insert('3')
	s2.Insert('4')
	fmt.Println("交", s1.Union(s2))
	fmt.Println("并", s1.Unite(s2))
	fmt.Println("差", s1.Diff(s2))
}

type testStruct struct {
	Age  int
	Name string
}

func TestUnion(t *testing.T) {
	s1, s2 := &complicated.MySet{}, &complicated.MySet{}
	s1.Construct()
	s2.Construct()
	s1.Insert(testStruct{18, "Kami"})
	s1.Insert(testStruct{16, "Rilo"})
	s2.Insert(testStruct{18, "Ops"})
	s2.Insert(testStruct{18, "Kami"})
	fmt.Println(s1.Union(s2))
}
