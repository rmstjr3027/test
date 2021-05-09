package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var str, target string
	s.Split(bufio.ScanLines)
	s.Scan()
	str = s.Text()
	s.Scan()
	target = s.Text()
	index := 0
	count := 0

	for index < len(str) {
		find := strings.Index(str[index:], target)
		if find < 0 {
			break
		}
		index += len(target) + find
		count++
	}
	fmt.Fprintln(w, count)
}
