package main

import (
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
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	fmt.Print(stackSequence(nums, n))
}
