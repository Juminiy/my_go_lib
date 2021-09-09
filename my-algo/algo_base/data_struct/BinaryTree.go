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

// probloly a mistake
func NonRecursionDfs(bt *BinaryTree) {
	var s MyStack
	s.Push(bt)
	for bt != nil || !s.IsEmpty() {
		for bt != nil {
			s.Push(bt)
			bt = bt.left
		}
		if !s.IsEmpty() {
			if tBin, err := s.Top(); err != nil {
				bt = tBin.(*BinaryTree)
			}
			if err := s.Pop(); err == nil {
				s.Push(bt.right)
			}
		}
	}
}

func Bfs(bt *BinaryTree) {
	var q MyQueue
	q.Push(bt)
	for !q.IsEmpty() {
		tBin, err := q.Front()
		seq = append(seq, tBin)
		if err == nil {
			q.Push(tBin.(*BinaryTree).left)
			q.Push(tBin.(*BinaryTree).right)
		}
	}
}
