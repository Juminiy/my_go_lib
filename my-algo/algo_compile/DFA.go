package algo_compile

import (
	"github.com/Juminiy/my_go_lib/my-algo/algo_compile/fa"
)

/**
 * Deterministic finite automata
 */
type DFA struct {
	K            *fa.ISet // 有穷状态集合
	Sigma        *fa.ISet // 输入符号表
	InitialState *fa.ISet // 初始状态
	FinalState   *fa.ISet // 终结状态

}

//Delta function f if Delta unction
func (dfa *DFA) Delta(startState *fa.ISet, edgeA *fa.ISet) *fa.ISet {

	return nil
}

/**
 * Nondeterministic finite automata
 */
type NFA struct {
	K            *fa.ISet
	Sigma        *fa.ISet
	InitialState *fa.ISet
	FinalState   *fa.ISet

	withEpsilon bool // 是否带有空边
}

func (nfa *NFA) Delta(startState *fa.ISet, edgeA *fa.ISet) *fa.ISet {

	return nil
}

func (nfa *NFA) RegexToNFA(regexStr string) {

}
func (*DFA) NFAToDFA(nfa *NFA) (dfa *DFA) {

	return dfa
}

func (dfa *DFA) MinDFA() {

}
