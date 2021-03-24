package main

import (
	"bytes"
	"fmt"
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
func main() {
	var n, num int
	count := 1
	stackTop := 0
	var result bytes.Buffer
	fmt.Scan(&n)

	var stack = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		for count <= num {
			stack[stackTop] = count
			result.WriteString("+\n")
			//result = append(result, "+")
			count++
			stackTop++
		}
		if stack[stackTop-1] == num {
			stack[stackTop-1] = 0
			result.WriteString("-\n")
			//result = append(result, "-")
			stackTop--
		} else {
			fmt.Print("NO")
			return
		}
	}
	//result.
	fmt.Print(result.String())
}

//var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
//
//func nextInt() (r int) {
//	sc.Scan()
//	r = 0
//	for _, c := range sc.Bytes() {
//		r *= 10
//		r += int(c - '0')
//	}
//	return
//}
//func main() {
//	sc.Split(bufio.ScanWords)
//	wt := bufio.NewWriter(os.Stdout)
//	defer wt.Flush()
