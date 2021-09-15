package complicated

const (
	Less   = 1  // is considered to be ascending
	Equal  = 0  // is considered to be equal
	Larger = -1 // is considered to be descending

	Inc    = true
	Reduce = false
)

type bstNode struct {
	value       interface{}
	Left, Right *bstNode
	valueAmount int
	// valueComp interface{}
}

func IntValueComp(compValue, compedValue int) int {
	if compValue == compedValue {
		return Equal
	} else if compValue < compedValue {
		return Less
	} else {
		return Larger
	}
}
func (node *bstNode) defaultIntCmpPtr(compValue, compedValue interface{}) int {
	return IntValueComp(compValue.(int), compedValue.(int))
}

// isInc = true

func addIntNode(node *bstNode, value interface{}) *bstNode {
	if node == nil {
		return &bstNode{value, nil, nil, 1}
	}
	cmpRes := IntValueComp(value.(int), node.value.(int))
	if cmpRes == 0 {
		node.valueAmount++
	} else if cmpRes == 1 {
		node = addIntNode(node.Left, value)
	} else {
		node = addIntNode(node.Right, value)
	}
	return node
}

func deleteIntNode(node *bstNode, value interface{}) *bstNode {
	if node == nil {
		return nil
	}
	cmpRes := IntValueComp(value.(int), node.value.(int))
	if cmpRes == 0 {
		if node.valueAmount > 0 {
			node.valueAmount--
		}
	} else if cmpRes == 1 {
		node = deleteIntNode(node.Left, value)
	} else {
		node = deleteIntNode(node.Right, value)
	}
	return node
}

func manageNode(node *bstNode, value interface{}, isInc bool) *bstNode {
	return nil
}
func strikeNode(node *bstNode, value interface{}) *bstNode {
	return nil
}
func triggerAdjust() {

}
