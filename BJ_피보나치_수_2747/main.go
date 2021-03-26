package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MAX = 45 + 1
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	s.Scan()
	var a, b int
	num, _ := strconv.Atoi(s.Text())
	a, b = 0, 1
	if num == 0 {
		a = num
	} else if num == 1 {
		a = num
	} else {
		for i := 0; i < num; i++ {
			a, b = b, a+b
		}
	}
	fmt.Fprintln(w, a)
}
