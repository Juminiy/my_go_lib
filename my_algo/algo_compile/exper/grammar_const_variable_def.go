package exper

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
