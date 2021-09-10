package simple

const (
	binLevel          = 31
	averageNodeAmount = 1 << 15
	maxNodeAmount     = 1 << binLevel
)

type BinaryTree struct {
	left, right *BinaryTree
	value       interface{}
}
type BinSequence []interface{}

var seq BinSequence

func forTestInit() {
	seq = make([]interface{}, 0)
}
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
func NonRecursionDfs(bt *BinaryTree) []interface{} {
	var s MyStack
	ansSeq := make([]interface{}, 0)
	for bt != nil || !s.IsEmpty() {
		for bt != nil {
			s.Push(bt)
			ansSeq = append(ansSeq, bt.value)
			bt = bt.left
		}
		if !s.IsEmpty() {
			if tBin, err := s.Top(); err == nil {
				bt = tBin.(*BinaryTree)
			}
			if err := s.Pop(); err == nil {
				bt = bt.right
			}
		}
	}
	return ansSeq
}

func Bfs(bt *BinaryTree) []interface{} {
	var q MyQueue
	ansSeq := make([]interface{}, 0)
	q.Push(bt)
	for !q.IsEmpty() {
		tBin, err := q.Front()
		q.Pop()
		if tBin != nil {
			ansSeq = append(ansSeq, tBin.(*BinaryTree).value)
		}
		if err == nil {
			if tBin.(*BinaryTree).left != nil {
				q.Push(tBin.(*BinaryTree).left)
			}
			if tBin.(*BinaryTree).right != nil {
				q.Push(tBin.(*BinaryTree).right)
			}
		}
	}
	return ansSeq
}
