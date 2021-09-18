package grammar_analysis

import (
	"errors"
	"fmt"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/algo_basic/algo_string"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/complicated"
	"github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/simple"
	"github.com/Juminiy/my_go_lib/my_algo/algo_compile/finite_automata"
	"strings"
)

const (
	CharSharp    = '#'
	StringSharp  = "#"
	CharBlank    = ' '
	CharEpsilon  = 'ε'
	Epsilon      = "ε"
	EpsilonASCII = 949
	NonEpsilon   = "NOT_ε"
	DefaultK     = 1
)

type GrammarRegular struct {
	Header string
	Body   string
}
type GrammarRegulars struct {
	RegularMap map[string]*complicated.MySet
	VT         *complicated.MySet
	VN         *complicated.MySet
}

func (regulars *GrammarRegulars) Construct() {
	regulars.RegularMap = make(map[string]*complicated.MySet, 0)
	regulars.VT = &complicated.MySet{}
	regulars.VT.Construct()
	regulars.VN = &complicated.MySet{}
	regulars.VN.Construct()
}
func (regulars *GrammarRegulars) Add(header, body string) {
	if regulars.RegularMap[header] == nil {
		regulars.RegularMap[header] = &complicated.MySet{}
		regulars.RegularMap[header].Construct()
	}
	regulars.RegularMap[header].Insert(body)
}
func (regulars *GrammarRegulars) AddPair(regular *GrammarRegular) {
	if regular != nil && regular.Header != "" && regular.Body != "" {
		regulars.Add(regular.Header, regular.Body)
	}
}
func (regulars *GrammarRegulars) Delete(header, body string) {
	if regulars.RegularMap[header] != nil {
		regulars.RegularMap[header].Erase(body)
	}
}
func (regulars *GrammarRegulars) DeletePair(regular *GrammarRegular) {
	if regular != nil && regular.Header != "" && regular.Body != "" {
		regulars.Delete(regular.Header, regular.Body)
	}
}
func (regulars *GrammarRegulars) BatchAddPair(regularArr []GrammarRegular) {
	if regularArr != nil && len(regularArr) != 0 {
		for _, regular := range regularArr {
			regulars.AddPair(&regular)
		}
	}
}
func (regulars *GrammarRegulars) ClearAll() {
	if len(regulars.RegularMap) != 0 {
		for regularHeader, _ := range regulars.RegularMap {
			delete(regulars.RegularMap, regularHeader)
		}
	}
	regulars.VT.EraseAll()
	regulars.VN.EraseAll()
}
func CharIsDigit(char int32) bool {
	if char >= '0' && char <= '9' {
		return true
	} else {
		return false
	}
}

// CharIsTerminator 虽然 ε 是终结符 但是可以看成空 影响判断
func CharIsTerminator(char int32) bool {
	if char >= 'a' && char <= 'z' {
		return true
	} else {
		return false
	}
}
func CharIsNonTerminator(char int32) bool {
	if char >= 'A' && char <= 'Z' {
		return true
	} else {
		return false
	}
}

func IsEpsilon(epsilon string) string {
	for _, ele := range epsilon {
		if ele != CharEpsilon {
			return NonEpsilon
		}
	}
	return Epsilon
}

// GrammarAnalysis 输入的正规文法字符串转换为对应的产生式规则
func (regulars *GrammarRegulars) GrammarAnalysis(productions []string) error {
	if productions == nil || len(productions) == 0 {
		return errors.New("Input production strings is nil! ")
	}
	for _, production := range productions {
		if len(production) <= 0 {
			regulars.ClearAll()
			return errors.New("Input production string is blank! ")
		}
		header, body, isHeader := "", "", true
		for _, char := range production {
			charStr := fmt.Sprintf("%c", char)
			if char == '>' {
				isHeader = false
				body = ""
			} else if char == '|' {
				if body != "" {
					regulars.Add(header, body)
				}
				body = ""
			} else if CharIsTerminator(char) {
				regulars.VT.Insert(charStr)
				if isHeader {
					header += charStr
				} else {
					body += charStr
				}
			} else if char == 'ε' {
				regulars.VT.Insert(charStr)
				if isHeader {
					regulars.ClearAll()
					return errors.New("The 'ε' can't occur in the header! ")
				} else {
					if body == "" || IsEpsilon(body) == Epsilon {
						regulars.Add(header, Epsilon)
					} else {
						regulars.ClearAll()
						return errors.New("The 'ε' can't mixed with other V*! ")
					}
				}
			} else if CharIsNonTerminator(char) {
				regulars.VN.Insert(charStr)
				if isHeader {
					header += charStr
				} else {
					body += charStr
				}
			} else if char == CharBlank || char == '-' {

			} else {
				regulars.ClearAll()
				return errors.New("Input production string is illegal! ")
			}
		}
		if body != "" {
			regulars.Add(header, body)
		}
	}
	return nil
}

// DeducedEpsilon 从alpha经若干步骤推导出空边
func (regulars *GrammarRegulars) DeducedEpsilon(alpha string) bool {
	if len(alpha) == 0 {
		return false
	}
	productionQueue := &simple.MyQueue{}
	productionQueue.Push(alpha)
	for !productionQueue.IsEmpty() {
		productionInterface, _ := productionQueue.Front()
		productionQueue.Pop()
		productionStr := productionInterface.(string)
		if len(productionStr) > 0 {
			if productionStr == Epsilon {
				return true
			}
			for header, iSet := range regulars.RegularMap {
				match, index := algo_string.KMP(productionStr, header)
				if match && index == 0 {
					for body, _ := range iSet.ImmutableMap {
						generateStr := ""
						if body == Epsilon {
							generateStr = productionStr[len(header):len(productionStr)]
						} else {
							generateStr = strings.Replace(productionStr, header, body.(string), 1)
						}
						if generateStr == "" {
							generateStr = Epsilon
						}
						productionQueue.Push(generateStr)
					}
				}
			}
		}
	}
	return false
}

// FIRST calculate a terminal character set
// 从一条规约推导出的规约中第一个字符为终结符的其他规约，求这些终结符的集合，并判断是否是alpha的开始符号集
func (regulars *GrammarRegulars) FIRST(alpha string) (*finite_automata.ISet, bool) {
	if len(alpha) == 0 {
		return nil, false
	}
	aSet := &finite_automata.ISet{}
	aSet.Construct()
	isAlphaFirst := false
	productionQueue := &simple.MyQueue{}
	productionQueue.Push(alpha)
	for !productionQueue.IsEmpty() {
		productionInterface, _ := productionQueue.Front()
		productionQueue.Pop()
		productionStr := productionInterface.(string)
		if len(productionStr) > 0 {
			if CharIsTerminator(int32(productionStr[0])) {
				tChar := fmt.Sprintf("%c", productionStr[0])
				aSet.CharSet.Insert(tChar)
			} else if productionStr == Epsilon {
				aSet.CharSet.Insert(fmt.Sprintf("%c", CharEpsilon))
				isAlphaFirst = true
			} else {
				for header, iSet := range regulars.RegularMap {
					match, index := algo_string.KMP(productionStr, header)
					if match && index == 0 {
						for body, _ := range iSet.ImmutableMap {
							generateStr := ""
							if body == Epsilon {
								generateStr = productionStr[len(header):len(productionStr)]
							} else {
								generateStr = strings.Replace(productionStr, header, body.(string), 1)
							}
							if generateStr == "" {
								generateStr = Epsilon
							}
							productionQueue.Push(generateStr)
						}
					}
				}
			}
		}
	}
	return aSet, isAlphaFirst
}

// CalculateTerminatorsFIRST 求出所有产生式右部的FIRST集
func (regulars *GrammarRegulars) CalculateTerminatorsFIRST() map[string]*complicated.MySet {
	allFirst := make(map[string]*complicated.MySet, 0)
	for _, alphaSet := range regulars.RegularMap {
		for alpha, _ := range alphaSet.ImmutableMap {
			alphaStr := alpha.(string)
			alphaFirstSet, _ := regulars.FIRST(alphaStr)
			allFirst[alphaStr] = alphaFirstSet.CharSet
		}
	}
	return allFirst
}

// FOLLOW calculate a terminal character set
// 一条产生式的header的后跟终结符集合,经所有产生式推导出A后紧挨着的第一个非终结符
// 起始 S 经若干步推导到以A为末尾的产生式 #包含
func (regulars *GrammarRegulars) FOLLOW(A string) *finite_automata.ISet {
	aSet := &finite_automata.ISet{}
	aSet.Construct()

	return aSet
}

func (regulars *GrammarRegulars) SELECT(regular *GrammarRegular) *finite_automata.ISet {
	firstAlpha, deducedEpsilon := regulars.FIRST(regular.Header)
	if !deducedEpsilon {
		return firstAlpha
	} else {
		// epsilonSet := &complicated.MySet{} ; epsilonSet.Construct() ; epsilonSet.Insert(complicated.Epsilon)
		firstAlpha.CharSet.Erase(Epsilon)
		selectSet := &finite_automata.ISet{CharSet: firstAlpha.CharSet.Unite(regulars.FOLLOW(regular.Header).CharSet)}
		return selectSet
	}
}
func (regulars *GrammarRegulars) CalculateProductionsSELECT() map[GrammarRegular]*complicated.MySet {
	allSelect := make(map[GrammarRegular]*complicated.MySet, 0)

	return allSelect
}

// EliminateLeftRecursion 消除左递归
func (regulars *GrammarRegulars) EliminateLeftRecursion() {

}

// IsLLk simple
// k == 1 相同的header推导出SELECT集合不可相交，相交判断即为假
func (regulars *GrammarRegulars) IsLLk(k int) bool {
	if k == -1 {
		k = DefaultK
	}

	return true
}
