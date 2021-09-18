package grammar_analysis

import (
	"fmt"
	"testing"
)

// "S -> MH | a","H->LSo|ε","K->dML|ε","L->eHf","M->K|bLM"
// "S->ABCDE","A->a|ε","B->b|ε","C->c|ε","D->d|ε","E->e|ε"
// "S ->ABc","A->a|ε","B->b|ε"
func TestMultiMap(t *testing.T) {
	regulars := GrammarRegulars{}
	regulars.Construct()
	regulars.GrammarAnalysis([]string{"S ->ABc", "A->a|ε", "B->b|ε"})
	fmt.Println("map = ", regulars.RegularMap)
	fmt.Println("VT = ", regulars.VT)
	fmt.Println("VN = ", regulars.VN)
	bodyFirstSet := regulars.CalculateTerminatorsFIRST()
	fmt.Println(bodyFirstSet)
}
