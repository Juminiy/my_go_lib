package exper

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

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
	//defer file.Close()
	buffReader := bufio.NewReader(file)
	os.Remove(output)
	os.Create(output)
	outfile, err := os.OpenFile(output, os.O_RDWR, 0777)
	if err != nil {
		os.Create(output)
	}
	//defer file.Close()
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
