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

type people struct {
	age  int
	name string
}

type peoples []people

func (p peoples) Len() int { return len(p) }

func main() {
	defer w.Flush()
	var testCase int
	fmt.Fscanln(r, &testCase)
	p := make([]people, testCase)

	for i := 0; i < testCase; i++ {
		fmt.Fscanf(r, "%d %s\n", &p[i].age, &p[i].name)
	}
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].age < p[j].age
	})

	for i := range p {
		fmt.Fprintf(w, "%d %s\n", p[i].age, p[i].name)
	}
}

//
//type Person struct {
//	age int
//	name string
//	order int
//}
//
//type Persons []Person
//
//func (a Persons) Len() int { return len(a)}
//func (a Persons) Swap(i int, j int) { a[i], a[j] = a[j], a[i]}
//func (a Persons) Less(i int, j int) bool {
//	return a[i].age < a[j].age || a[i].age == a[j].age && a[i].order < a[j].order
//}
//
//func main() {
//	n := scanInt()
//	persons := make([]Person, n)
//	for i:=0; i<n; i++ {
//		persons[i] = Person{age: scanInt(), name: scanString(), order: i}
//	}
//	sort.Sort(Persons(persons))
//	for i:=0; i<n; i++ {
//		p := persons[i]
//		w.WriteString(strconv.Itoa(p.age))
//		w.WriteByte(' ')
//		w.WriteString(p.name)
//		w.WriteByte('\n')
//	}
//	w.Flush()
//}
