package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stackSequence(nums []int, n int) string {
	count := 1
	stackTop := 0
	var stack = make([]int, n)
	var result []string
	for _, num := range nums {
		for count <= num {
			stack[stackTop] = count
			result = append(result, "+")
			count++
			stackTop++
		}
		if stack[stackTop-1] == num {
			stack[stackTop-1] = 0
			result = append(result, "-")
			stackTop--
		} else {
			return "NO"
		}
	}
	return strings.Join(result, "\n")
}

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

//func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
//func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	var n, num int
	count := 1
	stackTop := 0
	var result bytes.Buffer
	s.Scan()
	n, _ = strconv.Atoi(s.Text())

	var stack = make([]int, n)

	for i := 0; i < n; i++ {
		s.Scan()
		num, _ = strconv.Atoi(s.Text())
		for count <= num {
			stack[stackTop] = count
			result.WriteString("+\n")
			count++
			stackTop++
		}
		if stack[stackTop-1] == num {
			stack[stackTop-1] = 0
			result.WriteString("-\n")
			stackTop--
		} else {
			fmt.Print("NO")
			return
		}
	}
	fmt.Print(result.String())
}
