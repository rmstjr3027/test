//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//type node struct {
//	node map[int]int
//}
//
//var (
//	r = bufio.NewReader(os.Stdin)
//	w = bufio.NewWriter(os.Stdout)
//	graph []node
//	//visited []bool
//	start, end int
//)
//
//func bfs(w int) bool {
//	queue := make([]int, 0, len(graph))
//	queue = append(queue, start)
//	visited := make([]bool, len(graph)+1)
//	visited[start] = true
//	for len(queue) != 0 {
//		cur := queue[0]
//		queue = queue[1:]
//		if cur == end { return true }
//
//		//현재 섬에서 갈 수 있는 node를 확인 map을 확인
//		for k, v := range graph[cur].node {
//			// k 섬을 방문하지 않았으면서 다리의 중량이 w를 견딜 수 있는경우
//			if !visited[k] && w <= v {
//				visited[k] = true
//				queue = append(queue, k)
//			}
//		}
//	}
//	return false
//}
//
//func binarySearch(min, max int) int {
//	var low, mid, high int
//	low, high = min, max
//	result := min
//
//	for low <= high {
//		//for i := 0; i < len(visited); i++ {
//		//	visited[i] = false
//		//}
//		mid = (low + high) / 2
//		if bfs(mid) {
//			result = mid
//			low = mid+1
//		} else {
//			high = mid-1
//		}
//	}
//	return result
//}
//
//func main() {
//	defer w.Flush()
//	var n, m, min, max int
//	fmt.Fscanf(r, "%d %d\n", &n, &m)
//
//	// 그래프 초기화
//	graph = make([]node, n+1)
//	//visited = make([]bool, n+1)
//	for i := 0; i <= n; i++ {
//		graph[i].node = make(map[int]int)
//	}
//	// 다리 정보 입력
//	for i := 0; i < m; i++ {
//		var x, y, w int
//		fmt.Fscanf(r, "%d %d %d\n", &x, &y, &w)
//		graph[x].node[y] = w
//		graph[y].node[x] = w
//		if i == 0 {
//			min, max = 0, w
//		} else {
//			if min > w {min = w}
//			if max < w {max = w}
//		}
//	}
//	// 시작, 도착 노드 정보
//	fmt.Fscanf(r, "%d %d\n", &start, &end)
//
//	fmt.Fprintln(w, binarySearch(min, max))
//}
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Fio struct {
	r *bufio.Scanner
	w *bufio.Writer
}

func (fio *Fio) init() {
	fio.r = bufio.NewScanner(os.Stdin)
	fio.r.Split(bufio.ScanWords)
	fio.w = bufio.NewWriter(os.Stdout)
}

func (fio *Fio) readInt() int {
	fio.r.Scan()
	ret, _ := strconv.Atoi(fio.r.Text())
	return ret
}

func (fio *Fio) read() string {
	fio.r.Scan()
	return fio.r.Text()
}

func (fio *Fio) writeInt(i int) {
	fio.w.WriteString(strconv.Itoa(i))
	fio.w.WriteByte('\n')
}

func (fio *Fio) write(s string) {
	fio.w.WriteString(s)
	fio.w.WriteByte('\n')
}

type Edge struct {
	v, c int
}

type EdgeInput struct {
	u, v, c int
}

func bfs(m, S, T int, adj [][]Edge) bool {
	q := make([]int, 0, len(adj))
	q = append(q, S)
	visit := make([]bool, len(adj))
	visit[S] = true
	u := 0
	for len(q) > 0 {
		u, q = q[0], q[1:]
		for _, e := range adj[u] {
			if e.v == T && e.c >= m {
				return false
			}
			if !visit[e.v] && e.c >= m {
				visit[e.v] = true
				q = append(q, e.v)
			}
		}
	}
	return true
}

func binSearch(S, T int, adj [][]Edge) int {
	lb, ub := 1, int(1e9+1)
	for lb+1 != ub {
		m := (lb + ub) / 2
		if bfs(m, S, T, adj) {
			ub = m
		} else {
			lb = m
		}
	}
	return lb
}

func main() {
	var n int64
	n = 13
	var bin []bool
	var cnt, mask, addBit int64
	checkOne := false
	for i := 64 - 1; i >= 0; i-- {
		mask = 1 << i
		addBit = n & mask
		if addBit > 0 {
			checkOne = true
		}
		if checkOne {
			if addBit > 0 {
				bin = append(bin, true)
			} else {
				bin = append(bin, false)
			}
		}
	}

	for i := 0; i < len(bin); i++ {
		for bin[i] {
			checkZero := true
			for j := i + 2; j < len(bin); j++ {
				if bin[i+1] == true && bin[j] == false {
					continue
				} else {
					checkZero = false
				}
			}
			if checkZero {
				bin[i] = false
			} else {
				bin[len(bin)-1] = !bin[len(bin)-1]
			}
		}
	}
	fmt.Print(cnt)

	//var nums []int
	//c := len(nums)
	//tmp := make([][2]int32, c)
	//tmp[0][1] = 10
	//var num, mask int32
	//num = 31
	//count := 0
	//for i := 32 - 1; i >= 0; i-- {
	//	mask = 1 << i
	//	aaa := num & mask
	//	if aaa > 0 {
	//		count++
	//	}
	//}

	fio := new(Fio)
	fio.init()
	N, M := fio.readInt(), fio.readInt()
	ein := make([]EdgeInput, M)
	adj := make([][]Edge, N+1)
	for i := 0; i < M; i++ {
		u, v, c := fio.readInt(), fio.readInt(), fio.readInt()
		if u > v {
			u, v = v, u
		}
		ein[i] = EdgeInput{u, v, c}
	}
	//str := "...as@%@%*&df..few..@#%.."

	str := "...asdfsalkjlksaflk."
	str2 := ""
	r, _ := regexp.Compile("[^a-z0-9A-Z-_.]")
	index := r.FindStringIndex(str)
	if index != nil {
		str2 = r.ReplaceAllString(str, "")
	}
	for strings.Index(str, "..") > -1 {
		//str2 = r.ReplaceAllString(str, "")
		str2 = strings.ReplaceAll(str, "..", ".")
		str = str2
	}
	sort.Slice(ein, func(i, j int) bool {
		return ein[i].u < ein[j].u || (ein[i].u == ein[j].u && ein[i].v < ein[j].v)
	})
	for i := 0; i < M; i++ {
		if i != M-1 && ein[i].u == ein[i+1].u && ein[i].v == ein[i+1].v {
			ein[i+1].c += ein[i].c
		} else {
			adj[ein[i].u] = append(adj[ein[i].u], Edge{ein[i].v, ein[i].c})
			adj[ein[i].v] = append(adj[ein[i].v], Edge{ein[i].u, ein[i].c})
		}
	}
	S, T := fio.readInt(), fio.readInt()
	fmt.Println(binSearch(S, T, adj))
}
