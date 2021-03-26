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

type point struct {
	x int
	y int
}

type points []point

func (p points) Len() int      { return len(p) }
func (p points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p points) Less(i, j int) bool {
	return p[i].x < p[j].x || (p[i].x == p[j].x && p[i].y < p[j].y)
}

func main() {
	defer w.Flush()
	var testCase int
	fmt.Fscanln(r, &testCase)
	p := make([]point, testCase)
	for i := 0; i < testCase; i++ {
		fmt.Fscanf(r, "%d %d\n", &p[i].x, &p[i].y)
	}
	sort.Sort(points(p))
	for i := 0; i < len(p); i++ {
		fmt.Fprintf(w, "%d %d\n", p[i].x, p[i].y)
	}
}
