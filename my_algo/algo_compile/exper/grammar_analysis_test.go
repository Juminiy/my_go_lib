package exper

import (
	"fmt"
	"testing"
)

func TestRemoveLeftToken(t *testing.T) {
	RemoveLeftToken("./testdata/ans/grammar_ans.txt", "./testdata/out/lexical/output_left_token.txt")
	CompareTwoFile("./testdata/ans/lexical_ans.txt", "./testdata/out/lexical/output_left_token.txt")
}

// 总测试
func TestGrammarAnalysis1(t *testing.T) {
	GrammarAnalysis("./testdata/in/grammar/const_1.txt", "./testdata/out/grammar/output.txt")
	fmt.Println("Finished ", len(TokenSequence), " lines! ")
	CompareTwoFile("./testdata/out/grammar/output.txt", "./testdata/ans/grammar_ans.txt")
}

// 常量定义 & 变量定义
func TestGrammarAnalysis2(t *testing.T) {
	GrammarAnalysis("./testdata/in/grammar/variable_1.txt", "./testdata/out/grammar/output_1.txt")
}

// 函数定义
func TestGrammarAnalysis3(t *testing.T) {
	GrammarAnalysis("./testdata/in/grammar/const_1.txt", "./testdata/out/grammar/output_3.txt")
	fmt.Println("Finished ", len(TokenSequence), " lines! ")
	CompareTwoFile("./testdata/out/grammar/output_3.txt", "./testdata/ans/grammar_ans.txt")
}
