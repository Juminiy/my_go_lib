package grammar_analysis

import (
	"fmt"
	"testing"
)

// "S -> MH | a","H->LSo|ε","K->dML|ε","L->eHf","M->K|bLM"
// "S->ABCDE","A->a|ε","B->b|ε","C->c|ε","D->d|ε","E->e|ε"
// "S ->ABc","A->a|ε","B->b|ε"
// 消除左递归 S -> aB|aC|aD|aE|bF|bG ---> S->aA'|bB',A'->B|C|D|E,B'->F|G
// S -> aBC|aBD|aE --->
// 输入太麻烦了，应该弄个手机拍照经过神经网络识别串组直接传到后端运行返回结果
func TestMultiMap(t *testing.T) {
	regulars := GrammarRegulars{}
	regulars.Construct()
	regulars.GrammarAnalysis([]string{"E->TY", "Y->+E|ε", ""})
	fmt.Println("map = ", regulars.RegularMap)
	fmt.Println("VT = ", regulars.VT)
	fmt.Println("VN = ", regulars.VN)
	bodyFirstSet := regulars.CalculateTerminatorsFIRST()
	fmt.Println(bodyFirstSet)
}
