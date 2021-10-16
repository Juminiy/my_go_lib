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

var SpeciesCode map[string]string

func InitSpeciesCode() {
	SpeciesCode = make(map[string]string, 40)
	SpeciesCode["identifier"] = "IDENFR"
	SpeciesCode["integer"] = "INTCON"
	SpeciesCode["charc"] = "CHARCON"
	SpeciesCode["string"] = "STRCON"
	SpeciesCode["const"] = "CONSTTK"
	SpeciesCode["int"] = "INTTK"
	SpeciesCode["char"] = "CHARTK"
	SpeciesCode["void"] = "VOIDTK"
	SpeciesCode["main"] = "MAINTK"
	SpeciesCode["if"] = "IFTK"
	SpeciesCode["else"] = "ELSETK"
	SpeciesCode["do"] = "DOTK"
	SpeciesCode["while"] = "WHILETK"
	SpeciesCode["for"] = "FORTK"
	SpeciesCode["scanf"] = "SCANFTK"
	SpeciesCode["printf"] = "PRINTFTK"
	SpeciesCode["return"] = "RETURNTK"
	SpeciesCode["+"] = "PLUS"
	SpeciesCode["-"] = "MINU"
	SpeciesCode["*"] = "MULT"
	SpeciesCode["/"] = "DIV"
	SpeciesCode["<"] = "LSS"
	SpeciesCode["<="] = "LEQ"
	SpeciesCode[">"] = "GRE"
	SpeciesCode[">="] = "GEQ"
	SpeciesCode["=="] = "EQL"
	SpeciesCode["!="] = "NEQ"
	SpeciesCode["="] = "ASSIGN"
	SpeciesCode[";"] = "SEMICN"
	SpeciesCode[","] = "COMMA"
	SpeciesCode["("] = "LPARENT"
	SpeciesCode[")"] = "RPARENT"
	SpeciesCode["["] = "LBRACK"
	SpeciesCode["]"] = "RBRACK"
	SpeciesCode["{"] = "LBRACE"
	SpeciesCode["}"] = "RBRACE"
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
			writeOutputFile(file, tmpStr, curStr)
		} else {
			if unicode.IsLetter(rune(curStr[0])) || curStr[0] == '_' {
				writeOutputFile(file, "identifier", curStr)
			} else if unicode.IsDigit(rune(curStr[0])) || curStr[0] == '+' || curStr[0] == '-' {
				writeOutputFile(file, "integer", curStr)
			}
		}
	}
	return nil
}
func processLine(file *os.File, line string) error {
	if len(line) == 0 {
		return nil
	}
	var curStr string
	for i := 0; i < len(line); {
		c := line[i]
		if c == ';' || c == ',' || c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}' || c == '-' || c == '+' || c == '*' || c == '/' {
			doCurStr(file, curStr)
			curStr = fmt.Sprintf("%c", c)
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
			writeOutputFile(file, "integer", curStr)
			curStr = ""
		} else if c == '\'' { // 多个字符则为错误, ''不匹配则为错误
			if i+2 < len(line) && line[i+1] >= 32 && line[i+1] <= 126 && line[i+2] == '\'' && line[i+1] != ';' {
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
				writeOutputFile(file, "string", curStr)
				curStr = ""
				i++
			} else {
				// ERROR
			}
		} else {
			if c == '<' {
				i++
				if line[i] == '=' {
					writeOutputFile(file, "<=", "<=")
				} else {
					writeOutputFile(file, "<", "<")
				}
			} else if c == '=' {
				i++
				if line[i] == '=' {
					writeOutputFile(file, "==", "==")
				} else {
					writeOutputFile(file, "=", "=")
				}
			} else if c == '>' {
				i++
				if line[i] == '=' {
					writeOutputFile(file, ">=", ">=")
				} else {
					writeOutputFile(file, ">", ">")
				}
			} else if c == '!' {
				i++
				if line[i] == '=' {
					writeOutputFile(file, "!=", "!=")
				} else {
					// ERROR
				}
			}
		}
	}
	return nil
}
func AnalysisToken(input, output string) error {
	InitSpeciesCode()
	file, err := os.OpenFile(input, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	buffReader := bufio.NewReader(file)
	if _, err := os.Stat(output); err != nil {
		if os.IsExist(err) {
			os.Create(output)
		}
	}
	outfile, err := os.OpenFile(output, os.O_APPEND, 0666)
	if err != nil {
		os.Create(output)
	}
	defer file.Close()
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

func CompareTwoFile(res, ans string) {
	file, err := os.OpenFile(res, os.O_APPEND, 0666)
	file_result, err := os.OpenFile(ans, os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	fr1, fr2 := bufio.NewReader(file), bufio.NewReader(file_result)
	var lineNumber = 0
	for {
		lineNumber++
		line1, _, err1 := fr1.ReadLine()
		line2, _, err2 := fr2.ReadLine()
		str1, str2 := string(line1), string(line2)
		if str1 != str2 {
			fmt.Println(lineNumber, "is wrong")
		}
		if err1 == io.EOF || err2 == io.EOF {
			break
		}
	}
}
