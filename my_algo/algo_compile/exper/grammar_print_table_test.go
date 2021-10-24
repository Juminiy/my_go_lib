package exper

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

// WriteFileLine()
func TestGrammarDeduction(t *testing.T) {
	for i := 1; i <= 5; i++ {
		file, err := os.OpenFile("./submit/table_"+strconv.Itoa(i)+".txt", os.O_RDWR, 0777)
		if err != nil {
			log.Fatalln(err)
		}
		br := bufio.NewReader(file)
		file1, err := os.OpenFile("./submit/res/res_"+strconv.Itoa(i)+".txt", os.O_RDWR, 0777)
		os.Truncate("./submit/res/res_"+strconv.Itoa(i)+".txt", 0)
		if err != nil {
			log.Fatalln(err)
		}
		bw := bufio.NewWriter(file1)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			bw.WriteString("WriteFileLine(\"" + string(line) + "\")\n")
			bw.Flush()
		}
	}

}
func TestGrammarAnalysisX(t *testing.T) {
	fmt.Println("\\n")
}
