package data_struct

type BinaryTree struct {
	left, right *BinaryTree
	value       interface{}
}
type BinSequence []interface{}

var seq BinSequence

func PreOrderVisit(bt *BinaryTree) {
	if bt != nil {
		seq = append(seq, bt.value)
		PreOrderVisit(bt.left)
		PreOrderVisit(bt.right)
	}
}

func MidOrderVisit(bt *BinaryTree) {
	if bt != nil {
		MidOrderVisit(bt.left)
		seq = append(seq, bt.value)
		MidOrderVisit(bt.right)
	}
}

func PostOrderVisit(bt *BinaryTree) {

}
func NonRecursionDfs(bt *BinaryTree) {
	var s MyStack
	s.Push(bt)
	for bt != nil || !s.IsEmpty() {
		for bt != nil {
			s.Push(bt)
			bt = bt.left
		}
		if !s.IsEmpty() {
			if bt, err := s.Top(); err != nil {
				bt = bt.(*BinaryTree)
			}
			s.Pop()
			s.Push(bt.right)
		}
	}
}

func Bfs() {

}
