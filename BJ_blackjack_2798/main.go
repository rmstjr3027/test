package main

import (
	"fmt"
)

func blackjack(cards []int, n, m int) int {
	result := 0
	sum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sum = cards[i] + cards[j] + cards[k]
				if sum <= m {
					result = func(result, sum int) int {
						if result > sum {
							return result
						} else {
							return sum
						}
					}(result, sum)
				}
			}
		}
	}
	return result
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	cards := make([]int, n)
	for i := range cards {
		fmt.Scan(&cards[i])
	}
	fmt.Print(blackjack(cards, n, m))

}
