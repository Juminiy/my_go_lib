package compile_fa

import "my_algo/compile_fa/set_fa"

/**
 * Deterministic finite automata
 */
type DFA struct {
	K            *set_fa.LimitSet // 有穷状态集合
	Sigma        *set_fa.InputSet // 输入符号表
	InitialState *set_fa.LimitSet // 初始状态
	FinalState   *set_fa.LimitSet // 终结状态

}

func (dfa *DFA) Delta(startState *set_fa.LimitSet, edgeA *set_fa.InputSet) *set_fa.LimitSet {

	return nil
}

/**
 * Nondeterministic finite automata
 */
type NFA struct {
	K            *set_fa.LimitSet
	Sigma        *set_fa.InputSet
	InitialState *set_fa.LimitSet
	FinalState   *set_fa.LimitSet

	withEpsilon bool // 是否带有空边
}

func (nfa *NFA) Delta(startState *set_fa.LimitSet, edgeA *set_fa.InputSet) *set_fa.LimitSet {

	return nil
}

func (nfa *NFA) RegexToNFA(regexStr string) {

}
func (*DFA) NFAToDFA(nfa *NFA) (dfa *DFA) {

	return dfa
}

func (dfa *DFA) MinDFA() {

}