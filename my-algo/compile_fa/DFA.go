package compile_fa

import (
	"github.com/Juminiy/my_go_lib/my-algo/compile_fa/fa_data_structure"
)

/**
 * Deterministic finite automata
 */
type DFA struct {
	K            *fa_data_structure.ISet // 有穷状态集合
	Sigma        *fa_data_structure.ISet // 输入符号表
	InitialState *fa_data_structure.ISet // 初始状态
	FinalState   *fa_data_structure.ISet // 终结状态

}

//Delta function f if Delta unction
func (dfa *DFA) Delta(startState *fa_data_structure.ISet, edgeA *fa_data_structure.ISet) *fa_data_structure.ISet {

	return nil
}

/**
 * Nondeterministic finite automata
 */
type NFA struct {
	K            *fa_data_structure.ISet
	Sigma        *fa_data_structure.ISet
	InitialState *fa_data_structure.ISet
	FinalState   *fa_data_structure.ISet

	withEpsilon bool // 是否带有空边
}

func (nfa *NFA) Delta(startState *fa_data_structure.ISet, edgeA *fa_data_structure.ISet) *fa_data_structure.ISet {

	return nil
}

func (nfa *NFA) RegexToNFA(regexStr string) {

}
func (*DFA) NFAToDFA(nfa *NFA) (dfa *DFA) {

	return dfa
}

func (dfa *DFA) MinDFA() {

}
