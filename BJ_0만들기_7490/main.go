package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	nums []int
)

func printNum(c []string) {
	for i := 0; i < len(nums); i++ {
		fmt.Fprint(w, nums[i])
		if i != len(nums)-1 {
			fmt.Fprint(w, c[i])
		}
	}
	fmt.Fprint(w, "\n")
}

func makeZero(idx, sum, num, sign int, c []string) {
	if idx == len(nums) {
		sum += sign * num
		if sum == 0 {
			printNum(c)
		}
		return
	}
	// " "
	c[idx-1] = " "
	makeZero(idx+1, sum, num*10+nums[idx], sign, c)
	// "+"
	c[idx-1] = "+"
	makeZero(idx+1, sum+(sign*num), nums[idx], 1, c)
	// "-"
	c[idx-1] = "-"
	makeZero(idx+1, sum+(sign*num), nums[idx], -1, c)
}

func main() {
	defer w.Flush()
	var testCase, num int
	var c []string
	fmt.Fscanln(r, &testCase)

	for i := 0; i < testCase; i++ {
		fmt.Fscanln(r, &num)
		nums = nil
		c = nil
		c = make([]string, num-1)
		for j := 1; j <= num; j++ {
			nums = append(nums, j)
		}
		makeZero(1, 0, 1, 1, c)
		fmt.Fprint(w, "\n")
	}

}
