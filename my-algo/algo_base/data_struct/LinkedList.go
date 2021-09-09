package data_struct

import "fmt"

type LinkedList struct {
	NodeData interface{}
	NextNode *LinkedList
}

func insertNode(data interface{}, head *LinkedList) *LinkedList {
	tNode := head
	if tNode == nil {
		return &LinkedList{data, nil}
	}
	for tNode != nil {
		tNode = tNode.NextNode
	}
	nNode := &LinkedList{data, nil}
	tNode.NextNode = nNode
	return head
}

func deleteNode(data interface{}, head *LinkedList) *LinkedList {
	tNode := head
	if tNode != nil && tNode.NodeData == data {
		return head.NextNode
	}
	for tNode != nil && tNode.NextNode != nil && tNode.NextNode.NodeData != data {
		tNode = tNode.NextNode
	}
	if tNode != nil && tNode.NextNode != nil {
		tNode.NextNode = tNode.NextNode.NextNode
	}
	return head
}

func Ins(data interface{}, head *LinkedList) *LinkedList {
	return insertNode(data, head)
}
func Del(data interface{}, head *LinkedList) *LinkedList {
	return deleteNode(data, head)
}
func TraversalList(head *LinkedList) {
	tNode := head
	for tNode.NextNode != nil {
		fmt.Printf("[%p](%v)->", tNode, tNode.NodeData)
		tNode = tNode.NextNode
	}
	fmt.Printf("[%p](%v)\n", tNode, tNode.NodeData)
}

func reverseList(head *LinkedList) *LinkedList {
	if head == nil || head.NextNode == nil {
		return head
	}
	var tNode, nNode *LinkedList = nil, head
	for nNode != nil {
		nNode = head.NextNode
		head.NextNode = tNode
		tNode = head
		head = nNode
	}
	return tNode
}

func Rev(head *LinkedList) *LinkedList {
	return reverseList(head)
}
