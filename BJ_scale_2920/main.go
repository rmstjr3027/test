package main

import "fmt"

func scale(nums []int, n int) string {
	var result string
	ascending := true
	descending := true
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			descending = false
		} else if nums[i] < nums[i-1] {
			ascending = false
		}
	}

	if ascending {
		result = "ascending"
	} else if descending {
		result = "descending"
	} else {
		result = "mixed"
	}
	return result
}

func main() {
	num := make([]int, 8)
	for i := range num {
		fmt.Scan(&num[i])
	}
	fmt.Print(scale(num, 8))
}
