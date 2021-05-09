package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func Z(N, r, c int) int {
	var sum int
	x := int(math.Pow(2, float64(N)) / 2)
	y := x

	for i := N - 1; i >= 0; i-- {
		// half는 다음 사분면에서의 중앙점을 찍기 위한 값
		half := int(math.Pow(2, float64(i)) / 2)
		skip := int(math.Pow(math.Pow(2, float64(i+1)), 2)) / 4

		if c < x && r < y {
			// 1
			x -= half
			y -= half
		} else if c >= x && r < y {
			// 2
			x += half
			y -= half
			sum += skip
		} else if c < x && r >= y {
			// 3
			x -= half
			y += half
			sum += skip * 2
		} else {
			// r > x && c > y
			// 4
			x += half
			y += half
			sum += skip * 3
		}
	}
	return sum
}

func main() {
	defer writer.Flush()
	var N, r, c int
	fmt.Fscanf(reader, "%d %d %d\n", &N, &r, &c)

	fmt.Fprintln(writer, Z(N, r, c))
}
