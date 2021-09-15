package finite_automata

// DFA Deterministic finite automata
type DFA struct {
	K            *ISet // 有穷状态集合
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
	K            *ISet
	Sigma        *ISet
	InitialState *ISet
	FinalState   *ISet

	withEpsilon bool // 是否带有空边
}

// Delta is Edge Description
func (nfa *NFA) Delta(startState *ISet, edgeA *ISet) *ISet {
	return nil
}
