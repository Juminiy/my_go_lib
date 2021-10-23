package exper

import (
	"bufio"
	"fmt"
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
	funcTypeDefSignal = 0
	bw                = &bufio.Writer{}
	// 						funcName : returnType
	funcDef  = map[string]string{}
	funcMain = false
	nilStr   = ""
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
	of, err := os.OpenFile(out, os.O_RDWR, 0666)
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

// ＜程序＞    ::= ［＜常量说明＞］［＜变量说明＞］{＜有返回值函数定义＞|＜无返回值函数定义＞}＜主函数＞
// 无主函数 分析错误
// 分段顺序分析
// 函数名重复出错

func GrammarDeduction() {
	token := NextToken()
	if token.Value == "const" {
		token = constStatement(token)
	}
	fmt.Println("left token : ", token)
	ntoken := GetNextNThToken(1)
	if ntoken.Value != "(" {
		token = variableStatement(token)
	}
	fmt.Println("left token : ", token)
	for token != nil {
		token = functionDefinition(token)
		funcTypeDefSignal = 0
	}
	fmt.Println(token)
	if funcMain == false {
		// TODO error
	}
}

func GrammarAnalysis(input, output string) error {
	AnalysisToken(input, output+".tmp")
	os.Chmod(input, 0777)
	os.Remove(output)
	os.Create(output)
	file, err := os.OpenFile(output, os.O_RDWR, 0777)
	if err != nil {
		os.Create(output)
	}
	defer file.Close()
	bw = bufio.NewWriter(file)
	os.Truncate(output, 0)
	GrammarDeduction()
	return nil
}
