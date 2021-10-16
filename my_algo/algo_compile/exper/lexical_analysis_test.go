package exper

import (
	"strconv"
	"testing"
)

func TestProcessLine(t *testing.T) {
	InitSpeciesCode()
	//processLine()
}

func TestAnalysisToken(t *testing.T) {
	for i := 1; i <= 5; i++ {
		AnalysisToken("./testdata/in/testfile.txt_0"+strconv.Itoa(i)+".in", "./testdata/out/lexical/output.txt_0"+strconv.Itoa(i)+".out")
	}

	CompareTwoFile("./testdata/out/output.txt", "./testdata/ans/lexical_ans.txt")
}
