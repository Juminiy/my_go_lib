package pl0

import (
	"fmt"
	"testing"
	"unicode"
)

func TestConst(t *testing.T) {
	fmt.Println(nul, ident, number)

}

func TestReadPCode(t *testing.T) {
	//ReadPCode("./testdata/p_code_1.pl0")
	//processLine("var m,n,r,q;",nil)
	//fmt.Printf("%d %d %d %d\n",' ','	','\t','\n')
	fmt.Println(unicode.IsLetter('L'))
}
