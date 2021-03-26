package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MAX = 10000 + 1
)

var (
	//r      = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	var testCase int
	//max := 0
	s.Scan()
	testCase, _ = strconv.Atoi(s.Text())
	counts := make([]int, MAX)
	//var result bytes.Buffer

	for i := 0; i < testCase; i++ {
		s.Scan()
		num, _ := strconv.Atoi(s.Text())
		counts[num]++
		//if max < num {
		//	max = num
		//}
	}
	for i := 0; i < MAX; i++ {
		for j := 0; j < counts[i]; j++ {
			fmt.Fprintln(w, i)
			//result.WriteString(strconv.Itoa(i))
			//result.WriteString("\n")
		}
	}
	//fmt.Fprint(w, result.String())
}
