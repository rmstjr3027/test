package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX = 200000 + 1
)

var (
	r      = bufio.NewReader(os.Stdin)
	w      = bufio.NewWriter(os.Stdout)
	parent [MAX]int
	set    [MAX]int
)

func find(idx int) int {
	if idx == parent[idx] {
		return idx
	}
	parent[idx] = find(parent[idx])
	return parent[idx]
}

func union(x, y int) int {
	left := find(x)
	right := find(y)
	if left != right {
		if left < right {
			left, right = right, left
		}
		parent[right] = left
		set[left] += set[right]
		set[right] = 0
	}
	return set[left]
}

func main() {
	defer w.Flush()
	var testCase int
	fmt.Fscanf(r, "%d\n", &testCase)

	for i := 0; i < testCase; i++ {
		var relationshipCount int
		idx := 1
		friendMap := make(map[string]int)
		fmt.Fscanf(r, "%d\n", &relationshipCount)
		for j := 0; j < MAX; j++ {
			parent[j] = j
			set[j] = 1
		}
		for j := 0; j < relationshipCount; j++ {
			var name1, name2 string
			fmt.Fscanf(r, "%s %s\n", &name1, &name2)
			_, exists := friendMap[name1]
			if exists == false {
				friendMap[name1] = idx
				idx++
			}
			_, exists = friendMap[name2]
			if exists == false {
				friendMap[name2] = idx
				idx++
			}

			x := find(friendMap[name1])
			y := find(friendMap[name2])
			fmt.Fprintf(w, "%d\n", union(x, y))
		}
	}
}
