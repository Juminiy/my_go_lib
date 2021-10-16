package pl0

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const (
	nul       = iota
	ident     = iota
	number    = iota
	plus      = iota
	minus     = iota
	times     = iota
	slash     = iota
	oddsym    = iota
	eql       = iota
	neq       = iota
	lss       = iota
	leq       = iota
	gtr       = iota
	geq       = iota
	lparen    = iota
	rparen    = iota
	comma     = iota
	semicolon = iota
	period    = iota
	becomes   = iota
	beginsym  = iota
	endsym    = iota
	ifsym     = iota
	thensym   = iota
	whilesym  = iota
	writesym  = iota
	readsym   = iota
	dosym     = iota
	callsym   = iota
	constsym  = iota
	varsym    = iota
	procsym   = iota
)

var SpeciesCode map[string]int

func init() {
	SpeciesCode = make(map[string]int, 64)
	SpeciesCode[""] = 0
	SpeciesCode["identifier"] = 1
	SpeciesCode["constantNumber"] = 2
	SpeciesCode["begin"] = 3
	SpeciesCode["end"] = 4
	SpeciesCode["if"] = 5
	SpeciesCode["then"] = 6
	SpeciesCode["while"] = 7
	SpeciesCode["do"] = 8
	SpeciesCode["const"] = 9
	SpeciesCode["var"] = 10
	SpeciesCode["call"] = 11
	SpeciesCode["procedure"] = 12
	SpeciesCode["odd"] = 13
	SpeciesCode["read"] = 14
	SpeciesCode["write"] = 15
	SpeciesCode["+"] = 16
	SpeciesCode["-"] = 17
	SpeciesCode["*"] = 18
	SpeciesCode["/"] = 19
	SpeciesCode["="] = 20
	SpeciesCode["#"] = 21
	SpeciesCode["<"] = 22
	SpeciesCode["<="] = 23
	SpeciesCode[">"] = 24
	SpeciesCode[">="] = 25
	SpeciesCode[":="] = 26
	SpeciesCode["("] = 27
	SpeciesCode[")"] = 28
	SpeciesCode[","] = 29
	SpeciesCode["."] = 30
	SpeciesCode[";"] = 31
	SpeciesCode["(*"] = 32
	SpeciesCode["*)"] = 33
	SpeciesCode["//"] = 34
}

type Token struct {
	Id    int
	Value string
}

// 解决字符冲突问题
// 标识符只允许字母数字下划线
func processLine(line string) ([]Token, error) {
	var curStr string
	tokens := make([]Token, 0)
	for i, c := range line {
		if i == len(line)-1 && (c == ';' || c == '.') {
			tokens = append(tokens, Token{SpeciesCode[fmt.Sprintf("%c", c)], curStr})
		} else if c == ' ' || c == '\t' || c == ',' { // 空格,制表符 略过
			if SpeciesCode[curStr] != 0 {
				tokens = append(tokens, Token{SpeciesCode[curStr], "-"})
			} else {
				if unicode.IsDigit(rune(curStr[0])) {
					tokens = append(tokens, Token{SpeciesCode["constantNumber"], curStr})
				} else {
					tokens = append(tokens, Token{SpeciesCode["identifier"], curStr})
				}
			}
			curStr = ""
		} else if unicode.IsLetter(c) || c == '_' { // 标识符，关键字
			curStr = curStr + fmt.Sprintf("%c", c)
		} else if unicode.IsDigit(c) || c == '.' {
			curStr = curStr + fmt.Sprintf("%c", c)
		} else {
			if c == ':' {
				curStr = ":"
			} else if c == '=' {
				if curStr == ":" {
					curStr = ":="
					tokens = append(tokens, Token{SpeciesCode[curStr], "-"})
				} else {
					return nil, errors.New("Signal \":\" must follow a \"=\" ")
				}
			} else if c == '+' || c == '-' || c == '*' {

			}
		}
	}
	return tokens, nil
}
func OutputToken(tokens map[string]int) error {
	if len(tokens) == 0 {
		return errors.New("Tokens are nil! ")
	}
	for str, id := range tokens {
		fmt.Println("<", str, ",", id, ">")
	}
	return nil
}
func ReadPCode(codePath string) (map[string]int, error) {
	if !strings.Contains(codePath, ".pl0") {
		return nil, errors.New("File is not \".pl0\" code,please check! ")
	}
	file, err := os.Open(codePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buffReader := bufio.NewReader(file)
	for {
		line, _, err := buffReader.ReadLine()
		strLine := string(line)
		fmt.Println(strLine)
		if err == io.EOF {
			break
		}
	}
	tokenResult := make(map[string]int, 1024)

	if len(tokenResult) == 0 {
		return nil, errors.New("Token is empty,maybe \".pl0\" file is empty! ")
	} else {
		return tokenResult, nil
	}
}
