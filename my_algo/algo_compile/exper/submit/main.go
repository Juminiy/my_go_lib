package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type MyStack struct {
	stack []interface{}
}

func (s *MyStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *MyStack) Push(value interface{}) {
	if value != nil {
		s.stack = append(s.stack, value)
	}
}

func (s *MyStack) Pop() error {
	if s.IsEmpty() {
		return errors.New("Index is out of bounds,stack len = 0! ")
	}
	s.stack = s.stack[:len(s.stack)-1]
	return nil
}

func (s *MyStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Index is out of bounds,stack len = 0! ")
	}
	value := s.stack[len(s.stack)-1]
	return value, nil
}

var SpeciesCode = map[string]string{
	"":           "",
	"integer":    "INTCON",
	"string":     "STRCON",
	"const":      "CONSTTK",
	"charc":      "CHARCON",
	"main":       "MAINTK",
	"identifier": "IDENFR",
	"return":     "RETURNTK",
	"do":         "DOTK",
	"void":       "VOIDTK",
	"for":        "FORTK",
	"scanf":      "SCANFTK",
	"int":        "INTTK",
	"else":       "ELSETK",
	"char":       "CHARTK",
	"while":      "WHILETK",
	"printf":     "PRINTFTK",
	"if":         "IFTK",
	"-":          "MINU",
	"+":          "PLUS",
	"*":          "MULT",
	"/":          "DIV",
	"<":          "LSS",
	"<=":         "LEQ",
	">":          "GRE",
	">=":         "GEQ",
	"==":         "EQL",
	"!=":         "NEQ",
	"=":          "ASSIGN",
	";":          "SEMICN",
	",":          "COMMA",
	"(":          "LPARENT",
	")":          "RPARENT",
	"[":          "LBRACK",
	"]":          "RBRACK",
	"{":          "LBRACE",
	"}":          "RBRACE",
}

func writeOutputFile(file *os.File, spec string, curStr string) error {
	w := bufio.NewWriter(file)
	_, err := w.WriteString(SpeciesCode[spec] + " " + curStr + "\n")
	w.Flush()
	if err != nil {
		return err
	}
	return nil
}
func doCurStr(file *os.File, curStr string) error {
	if curStr != "" {
		var tmpStr string
		if unicode.IsLetter(rune(curStr[0])) {
			tmpStr = strings.ToLower(curStr)
		} else {
			tmpStr = curStr
		}
		if SpeciesCode[tmpStr] != "" {
			TokenSequence = append(TokenSequence, Token{SpeciesCode[tmpStr], curStr})
			writeOutputFile(file, tmpStr, curStr)
		} else {
			if unicode.IsLetter(rune(curStr[0])) || curStr[0] == '_' {
				TokenSequence = append(TokenSequence, Token{SpeciesCode["identifier"], curStr})
				writeOutputFile(file, "identifier", curStr)
			} else if unicode.IsDigit(rune(curStr[0])) || curStr[0] == '+' || curStr[0] == '-' {
				TokenSequence = append(TokenSequence, Token{SpeciesCode["integer"], curStr})
				writeOutputFile(file, "integer", curStr)
			}
		}
	}
	return nil
}

type Token struct {
	SpeciesCode string
	Value       string
}

var (
	TokenSequence []Token
	CurPtr        = 0
)

func InitTokenSequence() {
	TokenSequence = make([]Token, 0)

}
func (token *Token) String() string {
	return token.SpeciesCode + " " + token.Value
}
func processLine(file *os.File, line string) error {
	if len(line) == 0 {
		return nil
	}
	var curStr string
	for i := 0; i < len(line); {
		c := line[i]
		if c == ';' || c == ',' || c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}' || c == '-' || c == '+' || c == '*' || c == '/' || (c == '=' && line[i+1] != '=') {
			doCurStr(file, curStr)
			curStr = fmt.Sprintf("%c", c)
			TokenSequence = append(TokenSequence, Token{SpeciesCode[curStr], curStr})
			writeOutputFile(file, curStr, curStr)
			curStr = ""
			i++
		} else if c == ' ' || c == '\t' {
			doCurStr(file, curStr)
			i++
			curStr = ""
			// 标识符必须字母/_开头 , 数字是数字开头
		} else if unicode.IsLetter(rune(c)) || c == '_' {
			for i < len(line) && (unicode.IsLetter(rune(line[i])) || line[i] == '_' || unicode.IsDigit(rune(line[i]))) {
				curStr = curStr + fmt.Sprintf("%c", line[i])
				i++
			}
			doCurStr(file, curStr)
			curStr = ""
		} else if unicode.IsDigit(rune(c)) { // 出现多个-/+ 视为错误
			curStr = curStr + fmt.Sprintf("%c", c)
			i++
			for i < len(line) && unicode.IsDigit(rune(line[i])) {
				curStr = curStr + fmt.Sprintf("%c", line[i])
				i++
			}
			TokenSequence = append(TokenSequence, Token{SpeciesCode["integer"], curStr})
			writeOutputFile(file, "integer", curStr)
			curStr = ""
		} else if c == '\'' { // 多个字符则为错误, ''不匹配则为错误
			if i+2 < len(line) && line[i+1] >= 32 && line[i+1] <= 126 && line[i+2] == '\'' && line[i+1] != ';' {
				TokenSequence = append(TokenSequence, Token{SpeciesCode["charc"], fmt.Sprintf("%c", line[i+1])})
				writeOutputFile(file, "charc", fmt.Sprintf("%c", line[i+1]))
			} else {
				// ERROR
			}
			i += 3
		} else if c == '"' { // 无字符则为错误
			i++
			for i < len(line) && line[i] >= 32 && line[i] <= 126 && line[i] != '"' {
				curStr = curStr + fmt.Sprintf("%c", line[i])
				i++
			}
			if i < len(line) && line[i] == '"' {
				TokenSequence = append(TokenSequence, Token{SpeciesCode["string"], curStr})
				writeOutputFile(file, "string", curStr)
				curStr = ""
				i++
			} else {
				// ERROR
			}
		} else {
			specs, curs := "", ""
			if c == '<' {
				i++
				if line[i] == '=' {
					specs, curs = "<=", "<="
				} else {
					specs, curs = "<", "<"
				}
			} else if c == '=' {
				i++
				if line[i] == '=' {
					specs, curs = "==", "=="
				} else {
					specs, curs = "=", "="
				}
			} else if c == '>' {
				i++
				if line[i] == '=' {
					specs, curs = ">=", ">="
				} else {
					specs, curs = ">", ">"
				}
			} else if c == '!' {
				i++
				if line[i] == '=' {
					specs, curs = "!=", "!="
				} else {
					// ERROR
				}
			}
			TokenSequence = append(TokenSequence, Token{SpeciesCode[specs], curs})
			writeOutputFile(file, specs, curs)
			i++
		}
	}
	return nil
}
func AnalysisToken(input, output string) error {
	file, err := os.OpenFile(input, os.O_RDWR, 0777)
	os.Chmod(input, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	buffReader := bufio.NewReader(file)
	os.Remove(output)
	os.Create(output)
	outfile, err := os.OpenFile(output, os.O_RDWR, 0777)
	if err != nil {
		os.Create(output)
	}
	defer file.Close()
	InitTokenSequence()
	os.Truncate(output, 0)
	for {
		line, _, err := buffReader.ReadLine()
		strLine := string(line)
		processLine(outfile, strLine)
		if err == io.EOF {
			break
		}
	}
	return nil
}
func GetNextNThToken(n int) *Token {
	if CurPtr+n > len(TokenSequence) {
		return nil
	} else {
		return &TokenSequence[CurPtr+n]
	}
}
func GetNextToken() *Token {
	return GetNextNThToken(1)
}
func NextToken() *Token {
	cur := CurPtr
	CurPtr++
	if CurPtr > len(TokenSequence) {
		return nil
	} else {
		return &TokenSequence[cur]
	}
}

func CompareTwoFile(res, ans string) {
	file, err := os.OpenFile(res, os.O_APPEND, 0666)
	file_result, err := os.OpenFile(ans, os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	fr1, fr2 := bufio.NewReader(file), bufio.NewReader(file_result)
	var lineNumber = 0
	var cnt = 0
	for {
		lineNumber++
		line1, _, err1 := fr1.ReadLine()
		line2, _, err2 := fr2.ReadLine()
		str1, str2 := string(line1), string(line2)
		if str1 != str2 {
			fmt.Println(lineNumber, "is different")
			cnt++
		}
		if err1 == io.EOF || err2 == io.EOF {
			break
		}
	}
	if cnt == 0 {
		fmt.Println("All lines same,success! ")
	} else {
		fmt.Println(cnt, " lines diff")
	}
}

// 文法是单词，词法是句子
// 单个token传递给语法分析器，如果语法分析出一个完整句子，就输出相应的成分即可
var (
	analysisStk       = &MyStack{}
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
// 检查标识符重复错误
//只输出，不检查语法，不分析语法错误
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
func variableStatement(token *Token) *Token {
	// 变量定义; 变量定义... --> 变量说明 结束
	// 区别于 声明头部 int|char 标识符(
	// ＜变量定义＞  ::= ＜类型标识符＞(＜标识符＞|＜标识符＞'['＜无符号整数＞']'){,(＜标识符＞|＜标识符＞'['＜无符号整数＞']' )} //＜无符号整数＞表示数组元素的个数，其值需大于0
	WriteFileLine(token.String())
	/*
		int change1,change2,sss;
		char change3;
	*/
	for token != nil {
		ntoken := GetNextNThToken(2)
		// 非变量定义，直接转向函数定义
		if ntoken != nil && ntoken.Value == "(" {
			break
		} else {
			variableDefSignal = 1
		}
		token = NextToken()
		if token != nil {
			WriteFileLine(token.String())
		}
		for token != nil && token.Value != ";" {
			token = NextToken()
			if token.Value == ";" {
				WriteFileLine("<变量定义>")
				WriteFileLine(token.String())
				break
			} else {
				WriteFileLine(token.String())
			}
		}
	}
	if variableDefSignal == 1 {
		WriteFileLine("<变量说明>")
	}
	return token
}
func addDefFunction(funcName, returnType string) {
	if funcDef == nil || len(funcDef) == 0 {
		funcDef = make(map[string]string, 0)
	}
	if funcDef[funcName] == nilStr {
		funcDef[funcName] = returnType
	} else {
		// TODO error
	}

}
func WalkPrintToken(n int) {
	var token *Token
	for i := 1; i <= n; i++ {
		token = NextToken()
		if token != nil {
			WriteFileLine(token.String())
		}
	}
}

//＜因子＞    ::= ＜标识符＞｜＜标识符＞'['＜表达式＞']'|'('＜表达式＞')'｜＜整数＞|＜字符＞｜＜有返回值函数调用语句＞
//＜项＞     ::= ＜因子＞{＜乘法运算符＞＜因子＞}
//＜表达式＞    ::= ［＋｜－］＜项＞{＜加法运算符＞＜项＞}   //[+|-]只作用于第一个<项>
//＜赋值语句＞   ::=  ＜标识符＞＝＜表达式＞|＜标识符＞'['＜表达式＞']'=＜表达式＞
func exprSen() {

}

//＜读语句＞    ::=  scanf '('＜标识符＞{,＜标识符＞}')'
//	"scanf" : "SCANFTK",
func scanfSentence() (*Token, bool) {
	// scanf
	token := NextToken()
	if token != nil && token.Value == "scanf" {
		WriteFileLine(token.String())
		for token != nil {
			token = NextToken()
			if token.Value == ")" {
				WriteFileLine(token.String())
				WriteFileLine("<读语句>")
				break
			} else {
				WriteFileLine(token.String())
			}
		}
		// ;
		token = NextToken()
		WriteFileLine(token.String())
		WriteFileLine("<语句>")
		return nil, true
	} else {
		return token, false
	}
}

//	＜写语句＞    ::=
//	printf '(' ＜字符串＞,＜表达式＞ ')'
//	printf '('＜字符串＞ ')'
//	printf '('＜表达式＞')'
//    "printf" : "PRINTFTK",
func printfSentence() {
	// printf
	token := NextToken()
	if token != nil && token.Value == "printf" {
		WriteFileLine(token.String())
		for token != nil {
			token = NextToken()
			switch token.SpeciesCode {
			case "(":
			case ",":
				{
					WriteFileLine(token.String())
				}
			case "STRCON":
				{
					WriteFileLine(token.String())
					WriteFileLine("<字符串>")
				}
			case ")":
				{
					WriteFileLine(token.String())
					WriteFileLine("<写语句>")
					break
				}
			default:
				{
					exprSen()
				}
			}
		}
	}
	// ;
	token = NextToken()
	WriteFileLine(token.String())
	WriteFileLine("<语句>")
}

//＜值参数表＞   ::= ＜表达式＞{,＜表达式＞}｜＜空＞
//    ＜有返回值函数调用语句＞ ::= ＜标识符＞'('＜值参数表＞')'
//    ＜无返回值函数调用语句＞ ::= ＜标识符＞'('＜值参数表＞')'
func funcCallSen() {
	token := NextToken()
	if token != nil && funcDef[token.Value] != "" {
		WriteFileLine(token.String())
		for token != nil {

		}
	}
}

// ＜条件语句＞  ::= if '('＜条件＞')'＜语句＞［else＜语句>]
//	"if" : "IFTK",
//    "else" : "ELSETK",
//    ＜条件＞    ::=  ＜表达式＞＜关系运算符＞＜表达式＞｜＜表达式＞    //整型表达式之间才能进行关系运算,表达式为整型，其值为0条件为假，值不为0时条件为真
func condSen() {

}

// 内可嵌套语句
//＜循环语句＞   ::= do＜语句＞while '('＜条件＞')'
//"do" : "DOTK",
//＜循环语句＞   ::=  while '('＜条件＞')'＜语句＞
//"while" : "WHILETK",
//＜步长＞::= ＜无符号整数＞
//＜循环语句＞   ::= for'('＜标识符＞＝＜表达式＞;＜条件＞;＜标识符＞＝＜标识符＞(+|-)＜步长＞')'＜语句＞
//"for" : "FORTK",
func doWhileLoopSen() {

}
func whileLoopSen() {

}
func forLoopSen() {

}
func returnSen() {
	//token := NextToken()

}

// 函数里面的句型句法分析
// 识别一个函数
//＜程序＞    ::= ［＜常量说明＞］［＜变量说明＞］{＜有返回值函数定义＞|＜无返回值函数定义＞}＜主函数＞
func functionDefinition(token *Token) *Token {
	// int | void | char
	for token.Value == ";" {
		token = NextToken()
	}
	WriteFileLine(token.String())
	if token.Value != "void" {
		funcTypeDefSignal = 1
	}
	preStr := token.Value
	// funcName
	token = NextToken()
	WriteFileLine(token.String())
	// 主函数区别于其他函数
	if token.Value == "main" {
		if funcMain == true {
			// TODO error
		}
		funcMain = true
	} else {
		addDefFunction(token.Value, preStr)
	}
	// 处理参数表 ) 为止
	if funcMain {
		WalkPrintToken(2)
	} else {
		token = NextToken()
		for ; token != nil; token = NextToken() {
			if token.Value == "(" {
				WriteFileLine("<声明头部>")
			}
			if token.Value == ")" {
				WriteFileLine("<参数表>")
				WriteFileLine(token.String())
				break
			}
			WriteFileLine(token.String())
		}
	}
	// arguments table tst ok !
	// function body start
	// {
	token = NextToken()
	// 解决复合语句的嵌套 逐层向上的规约

	token = NextToken()
	// }
	// function body ended
	if funcMain {
		WriteFileLine("<主函数>")
		WriteFileLine("<程序>")
	} else {
		if funcTypeDefSignal == 1 {
			WriteFileLine("<有返回值函数定义>")
		} else {
			WriteFileLine("<无返回值函数定义>")
		}
	}
	return token
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

func main() {
	GrammarAnalysis("testfile.txt", "output.txt")
}
