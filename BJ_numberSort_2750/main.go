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

//func merge(lists []int, left, mid, right int) {
//	i := left
//	j := mid
//	k := right
//
//}
//
//func mergeSort(lists []int, left, right int) {
//	var mid int
//	if left < right {
//		mid = (left + right) / 2
//		mergeSort(lists, left, mid)
//		mergeSort(lists, mid+1, right)
//		merge(lists, left, mid, right)
//	}
//}

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
