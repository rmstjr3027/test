package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//vars *bufio.Scanner = bufio.NewScanner(os.Stdin)
var s *bufio.Reader = bufio.NewReader(os.Stdin)
var (
	w = bufio.NewWriter(os.Stdout)
)

//func scanInt() int {
//	s.Scan()
//
//}
func maxValue(queue *[]int) int {
	maxValue := 0
	for _, value := range *queue {
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	defer w.Flush()
	var testCase, n, m int

	str, _ := s.ReadString('\n')
	testCase, _ = strconv.Atoi(strings.TrimSuffix(str, "\n"))

	for i := 0; i < testCase; i++ {
		str, _ = s.ReadString('\n')
		nm := strings.Split(strings.TrimSuffix(str, "\n"), " ")
		n, _ = strconv.Atoi(nm[0])
		m, _ = strconv.Atoi(nm[1])

		str, _ = s.ReadString('\n')
		queueStrings := strings.Split(strings.TrimSuffix(str, "\n"), " ")
		queue := make([]int, n)
		for j, queueStr := range queueStrings {
			queue[j], _ = strconv.Atoi(queueStr)
		}

		count := 0
		index := 0
		max := 0
		for count < n {
			first := queue[index]
			if max == 0 {
				max = maxValue(&queue)
			}
			if first == max {
				max = 0
				count++
				queue[index] = 0
				if index == m {
					fmt.Fprintln(w, count)
					break
				}
			}
			index++
			index = index % n
		}
	}
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var reader *bufio.Reader = bufio.NewReader(os.Stdin)
//var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
//
//func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
//func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
//
//func main() {
//	defer writer.Flush()
//	var t int
//	scanf("%d\n", &t)
//
//	var n, q int
//loop:
//	for i := 0; i < t; i++ {
//		scanf("%d %d\n", &n, &q)
//		x := make([]int, n)
//		z := make([]int, 10)
//		for i := 0; i < n; i++ {
//			scanf("%d", &x[i])
//			z[x[i]]++
//		}
//		scanf("\n")
//
//		for cnt := 0; q >= 0; {
//			y := x[0]
//			ok := true
//			for i := y + 1; i < 10; i++ {
//				if z[i] > 0 {
//					ok = false
//					break
//				}
//			}
//
//			if ok {
//				cnt++
//				if q == 0 {
//					printf("%d\n", cnt)
//					continue loop
//				}
//				z[y]--
//				x = x[1:]
//				q--
//			} else {
//				x = append(x[1:], y)
//				if q == 0 {
//					q = len(x) - 1
//				} else {
//					q--
//				}
//			}
//		}
//	}
//}
