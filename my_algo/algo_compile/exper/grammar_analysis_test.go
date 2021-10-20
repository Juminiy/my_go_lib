package exper

import (
	"fmt"
	"testing"
)

func TestRemoveLeftToken(t *testing.T) {
	RemoveLeftToken("./testdata/ans/grammar_ans.txt", "./testdata/out/lexical/output_left_token.txt")
	CompareTwoFile("./testdata/ans/lexical_ans.txt", "./testdata/out/lexical/output_left_token.txt")
}

func TestGrammarAnalysis(t *testing.T) {
	GrammarAnalysis("./testdata/in/grammar/const_1.txt", "./testdata/out/grammar/output.txt")
	fmt.Println("Finished ", len(TokenSequence), " lines! ")
	CompareTwoFile("./testdata/out/grammar/output.txt", "./testdata/ans/grammar_ans.txt")
}
