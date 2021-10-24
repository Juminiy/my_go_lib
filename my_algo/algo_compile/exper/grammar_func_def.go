package exper

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
//＜赋值语句＞   ::=  ＜标识符＞＝＜表达式＞  |  ＜标识符＞'['＜表达式＞']'=＜表达式＞
func exprSen() {

}
func assignSen(token *Token) {

}

//＜读语句＞    ::=  scanf '('＜标识符＞{,＜标识符＞}')'
//	"scanf" : "SCANFTK",
func scanfSentence(token *Token) (*Token, bool) {
	// scanf

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
func printfSentence(token *Token) {
	// printf
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
			if token.Value == "(" && funcTypeDefSignal == 0 {
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
	// |
	token = NextToken()
	WriteFileLine(token.String())
	// 解决复合语句的嵌套 逐层向上的规约
	token = NextToken()
	token = constStatement(token)
	if token != nil && (token.Value == "int" || token.Value == "char") {
		token = variableStatement(token)
	}
	for token != nil {
		switch token.SpeciesCode {
		case "PRINTFTK":
			{
				printfSentence(token)
			}
		case "SCANFTK":
			{
				scanfSentence(token)
			}
		case "FORTK":
			{

			}
		case "DOTK":
			{

			}
		case "WHILETK":
			{

			}
		case "IFTK":
			{

			}
		case "RETURNTK":
			{

			}
		default:
			{

			}
		}
		token = NextToken()
	}

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
