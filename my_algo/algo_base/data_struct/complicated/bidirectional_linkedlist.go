package complicated

import "fmt"

type bNode struct {
	precursor, successor *bNode
	value                interface{}
}

type BList struct {
	head, tail *bNode
	len        int
}

func constructPrimaryNode(value interface{}) *BList {
	tNode := &bNode{precursor: nil, successor: nil, value: value}
	tList := &BList{head: tNode, tail: tNode, len: 0}
	return tList
}
func (list *BList) IsEmpty() bool {
	return list == nil || list.len == 0 || list.head == nil
}

func (list *BList) ForwardTraversal() {
	if list.IsEmpty() {
		return
	}
	tNode := list.head
	for tNode != nil {
		fmt.Printf("Address[%p]={%v}->", tNode, tNode.value)
		tNode = tNode.successor
	}
	fmt.Println("nil")
}

func (list *BList) OppositeTraversal() {
	if list.IsEmpty() {
		return
	}
	tNode := list.tail
	for tNode != nil {
		fmt.Printf("Address[%p]={%v}->", tNode, tNode.value)
		tNode = tNode.precursor
	}
	fmt.Println("nil")
}

func PushFront(list *BList, value interface{}) *BList {
	if value == nil {
		return nil
	}
	if list.IsEmpty() {
		list = constructPrimaryNode(value)
	} else {
		tNode := &bNode{precursor: nil, successor: list.head, value: value}
		list.head.precursor = tNode
		list.head = tNode
	}
	list.len++
	return list
}

func PopFront(list *BList) (*BList, interface{}) {
	if list.IsEmpty() {
		return nil, nil
	}
	tValue := list.head.value
	list.head = list.head.successor
	if list.head != nil {
		list.head.precursor = nil
	}
	list.len--
	if list.len == 0 {
		list = nil
	}
	return list, tValue
}

func PushBack(list *BList, value interface{}) *BList {
	if value == nil {
		return nil
	}
	if list.IsEmpty() {
		list = constructPrimaryNode(value)
	} else {
		tNode := &bNode{precursor: list.tail, successor: nil, value: value}
		list.tail.successor = tNode
		list.tail = tNode
	}
	list.len++
	return list
}

func PopBack(list *BList) (*BList, interface{}) {
	if list.IsEmpty() {
		return nil, nil
	}
	tValue := list.tail.value
	list.tail = list.tail.precursor
	if list.tail != nil {
		list.tail.successor = nil
	}
	list.len--
	if list.len == 0 {
		list = nil
	}
	return list, tValue
}
