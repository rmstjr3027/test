package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var testCase int
	fmt.Fscanln(r, &testCase)
	nums := make([]int, testCase)

	for i := 0; i < testCase; i++ {
		fmt.Fscanln(r, &nums[i])
	}
	sort.Ints(nums)
	for _, num := range nums {
		fmt.Fprintln(w, num)
	}
}
