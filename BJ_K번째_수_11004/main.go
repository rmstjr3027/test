package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func scanInt() int {
	s.Scan()
	tmp, _ := strconv.Atoi(s.Text())
	return tmp
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	var testCase, target int
	s.Scan()
	testCase, _ = strconv.Atoi(s.Text())
	s.Scan()
	target, _ = strconv.Atoi(s.Text())
	nums := make([]int, testCase)

	for i := 0; i < testCase; i++ {
		s.Scan()
		nums[i], _ = strconv.Atoi(s.Text())
	}
	//sort.Ints(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Fprintln(w, nums[target-1])
}
