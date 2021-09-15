package simple

const (
	binLevel          = 31
	averageNodeAmount = 1 << 15
	maxNodeAmount     = 1 << binLevel
)

type BinaryTree struct {
	Left, Right *BinaryTree
	Value       interface{}
}
type BinSequence []interface{}

var seq BinSequence

func forTestInit() {
	seq = make([]interface{}, 0)
}
func PreOrderVisit(bt *BinaryTree) {
	if bt != nil {
		seq = append(seq, bt.Value)
		PreOrderVisit(bt.Left)
		PreOrderVisit(bt.Right)
	}
}

func MidOrderVisit(bt *BinaryTree) {
	if bt != nil {
		MidOrderVisit(bt.Left)
		seq = append(seq, bt.Value)
		MidOrderVisit(bt.Right)
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
			ansSeq = append(ansSeq, bt.Value)
			bt = bt.Left
		}
		if !s.IsEmpty() {
			if tBin, err := s.Top(); err == nil {
				bt = tBin.(*BinaryTree)
			}
			if err := s.Pop(); err == nil {
				bt = bt.Right
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
			ansSeq = append(ansSeq, tBin.(*BinaryTree).Value)
		}
		if err == nil {
			if tBin.(*BinaryTree).Left != nil {
				q.Push(tBin.(*BinaryTree).Left)
			}
			if tBin.(*BinaryTree).Right != nil {
				q.Push(tBin.(*BinaryTree).Right)
			}
		}
	}
	return ansSeq
}
