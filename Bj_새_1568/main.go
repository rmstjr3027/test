package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanln(r, &N)
	k := 1
	count := 0

	for N != 0 {
		if k > N {
			k = 1
		}
		N -= k
		k++
		count++
	}
	fmt.Fprintln(w, "abc" > "aba")
	fmt.Fprintln(w, count)
}
