package complicated

type bstNode struct {
	value       interface{}
	left, right *bstNode

	valueComp interface{}
}

func IntValueComp(compValue, CompedValue int) bool {
	if compValue < CompedValue {
		return true
	} else {
		return false
	}
}
