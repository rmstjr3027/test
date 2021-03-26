package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var nums string
	fmt.Fscanln(r, &nums)
	var result bytes.Buffer
	numCount := make(map[rune]int)

	for _, c := range nums {
		numCount[c]++
	}
	for i := '9'; i >= '0'; i-- {
		for j := 0; j < numCount[i]; j++ {
			result.WriteString(string(i))
		}
	}

	//for i := 0; i < len(nums); i++ {
	//	tmp, _ := strconv.Atoi(string(nums[i]))
	//	numCount[tmp]++
	//}
	//for i := 9; i >= 0; i-- {
	//	for numCount[i] > 0 {
	//		numCount[i]--
	//		result.WriteString(strconv.Itoa(i))
	//	}
	//}
	fmt.Fprintln(w, result.String())
}
