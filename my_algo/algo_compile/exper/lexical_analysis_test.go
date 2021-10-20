package exper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestProcessLine(t *testing.T) {
	//processLine()
}

func TestAnalysisToken(t *testing.T) {
	tm := time.Now()
	for i := 1; i <= 5; i++ {
		AnalysisToken("./testdata/in/testfile.txt_0"+strconv.Itoa(i)+".in", "./testdata/out/lexical/hln/output.txt_0"+strconv.Itoa(i)+".out")
	}
	//for i:=1 ;i <= 5; i++ {
	//	CompareTwoFile("./testdata/out/lexical/output.txt_0"+strconv.Itoa(i)+".out","./testdata/out/lexical/yyz/output"+strconv.Itoa(i)+".txt")
	//}

	//AnalysisToken("./testdata/in/testfile.txt_0"+strconv.Itoa(2)+".in", "./testdata/out/lexical/output.txt_0"+strconv.Itoa(2)+".out")
	//CompareTwoFile("./testdata/out/lexical/output.txt_0"+strconv.Itoa(2)+".out","./testdata/out/lexical/yyz/output"+strconv.Itoa(1)+".txt")
	fmt.Println(time.Now().Sub(tm))
}

func TestRetry(t *testing.T) {
	file, _ := os.OpenFile("spec.txt", os.O_APPEND, 0666)
	br := bufio.NewWriter(file)
	for k, v := range SpeciesCode {
		br.WriteString("speciesCode.insert(\"" + k + "\" ," + "\"" + v + "\");\n")
		br.Flush()
	}
}
