package finite_automata

type IFA interface {
	Delta(startState *ISet, edgeA *ISet) *ISet
}

// DFA Deterministic finite automata
type DFA struct {
	S            *ISet // 有穷状态集合
	Sigma        *ISet // 输入符号表
	InitialState *ISet // 初始状态
	FinalState   *ISet // 终结状态
}

//Delta function f if Delta unction
func (dfa *DFA) Delta(startState *ISet, edgeA *ISet) *ISet {
	return nil
}

// NFA NonDeterministic finite automata
type NFA struct {
	S            *ISet
	Sigma        *ISet
	InitialState *ISet
	FinalState   *ISet

	withEpsilon bool // 是否带有空边
}

// Delta is Edge Description
func (nfa *NFA) Delta(startState *ISet, edgeA *ISet) *ISet {
	return nil
}

type DFATable struct {
	State       *ISet
	Input       *ISet
	ChangeTable [][]interface{}
}
