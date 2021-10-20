package exper

import (
	"bufio"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/simple"
	"io"
	"log"
	"os"
)

// 文法是单词，词法是句子
// 单个token传递给语法分析器，如果语法分析出一个完整句子，就输出相应的成分即可
var (
	analysisStk       = &simple.MyStack{}
	deductionStage    = 1
	constDefSignal    = 0
	variableDefSignal = 0
	functionDefSignal = 0
	bw                = &bufio.Writer{}
	funcDef           = map[string]string{}
)

func WriteFileLine(str string) {
	bw.WriteString(str + "\n")
	bw.Flush()
}
func RemoveLeftToken(in, out string) {
	file, err := os.OpenFile(in, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	br := bufio.NewReader(file)
	of, err := os.OpenFile(out, os.O_APPEND, 0666)
	bw = bufio.NewWriter(of)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if len(line) > 0 && line[0] != '<' {
			WriteFileLine(string(line))
		}
	}
}

func GrammarAnalysis(input, output string) error {
	AnalysisToken(input, output+".tmp")
	file, err := os.OpenFile(output, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	bw = bufio.NewWriter(file)
	os.Truncate(output, 0)
	GrammarDeduction()
	return nil
}

/*
	tkSpecCode == "IDENFR" ||
	tkSpecCode == "INTCON" ||
	tkSpecCode == "CHARCON" ||
	tkSpecCode == "INTTK" ||
	tkSpecCode =="CHARTK" ||
	tkSpecCode == "ASSIGN" ||
	tkSpecCode == "COMMA"  ||
	tkSpecCode == "MINUS" ||
	tkSpecCode == "PLUS"
*/
//只输出，不检查语法，不分析语法错误

func constStatement(token *Token) *Token {
	WriteFileLine(token.String())
	// 判断常量说明 与否,可能没有
	// 判断变量说明 与否,可能没有

	//可能有多行,不能结束
	// 常量定义; 常量定义... --> 常量说明 结束
	/**
	const int const1 = 1, const2 = -100;
	const char const3 = '_';
	int change1;
	char change3;
	int gets1(int var1,int var2){
	*/
	for token != nil && token.Value == "const" {
		constDefSignal = 1
		// int|char
		token = NextToken()
		WriteFileLine(token.String())
		for token.Value != ";" {
			tkSpecCode := token.SpeciesCode
			if tkSpecCode == "INTCON" {
				WriteFileLine("<无符号整数>")
				WriteFileLine("<整数>")
			}
			token = NextToken()
			if token.Value == ";" {
				WriteFileLine("<常量定义>")
				WriteFileLine(token.String())
				break
			}
			WriteFileLine(token.String())
		}
		token = NextToken()
		if token == nil || token.Value != "const" {
			break
		}
		WriteFileLine(token.String())
	}
	if constDefSignal == 1 {
		WriteFileLine("<常量说明>")
	}
	return token
}
func IsConstDef(tkSpecCode string) bool {
	if tkSpecCode == "IDENFR" ||
		tkSpecCode == "INTCON" ||
		tkSpecCode == "CHARCON" ||
		tkSpecCode == "INTTK" ||
		tkSpecCode == "CHARTK" ||
		tkSpecCode == "ASSIGN" ||
		tkSpecCode == "COMMA" ||
		tkSpecCode == "MINUS" ||
		tkSpecCode == "PLUS" {
		return true
	} else {
		return false
	}
}
func variableStatement(token *Token) *Token {
	// 变量定义; 变量定义... --> 变量说明 结束
	// 声明头部 int|char 标识符(

	for token != nil {

	}
	return token
}
func functionDefinition(funcName, returnType string) {
	if funcDef == nil || len(funcDef) == 0 {
		funcDef = make(map[string]string, 0)
	}
	funcDef[funcName] = returnType
}

// ＜程序＞    ::= ［＜常量说明＞］［＜变量说明＞］{＜有返回值函数定义＞|＜无返回值函数定义＞}＜主函数＞
// 无主函数 分析错误
// 分段顺序分析

func GrammarDeduction() {
	token := NextToken()
	token = constStatement(token)
	token = variableStatement(token)

}
