//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strings"
//)
//
//var (
//	r = bufio.NewReader(os.Stdin)
//	w = bufio.NewWriter(os.Stdout)
//)
//
//type people struct {
//	age  int
//	name string
//}
//
//type peoples []people
//
//func (p peoples) Len() int { return len(p) }
//
//func main() {
//
//	str := ".123.123.123"
//	strs := strings.Split(str, ".")
//	fmt.Println(strs)
//
//	defer w.Flush()
//	var testCase int
//	fmt.Fscanln(r, &testCase)
//	p := make([]people, testCase)
//
//	for i := 0; i < testCase; i++ {
//		fmt.Fscanf(r, "%d %s\n", &p[i].age, &p[i].name)
//	}
//	sort.SliceStable(p, func(i, j int) bool {
//		return p[i].age < p[j].age
//	})
//
//	for i := range p {
//		fmt.Fprintf(w, "%d %s\n", p[i].age, p[i].name)
//	}
//}
//
////
////type Person struct {
////	age int
////	name string
////	order int
////}
////
////type Persons []Person
////
////func (a Persons) Len() int { return len(a)}
////func (a Persons) Swap(i int, j int) { a[i], a[j] = a[j], a[i]}
////func (a Persons) Less(i int, j int) bool {
////	return a[i].age < a[j].age || a[i].age == a[j].age && a[i].order < a[j].order
////}
////
////func main() {
////	n := scanInt()
////	persons := make([]Person, n)
////	for i:=0; i<n; i++ {
////		persons[i] = Person{age: scanInt(), name: scanString(), order: i}
////	}
////	sort.Sort(Persons(persons))
////	for i:=0; i<n; i++ {
////		p := persons[i]
////		w.WriteString(strconv.Itoa(p.age))
////		w.WriteByte(' ')
////		w.WriteString(p.name)
////		w.WriteByte('\n')
////	}
////	w.Flush()
////}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'validateAddresses' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY addresses as parameter.
 */

func validateAddresses(addresses []string) []string {
	// Write your code here
	var result []string
	for i := range addresses {
		tmp := addresses[i]
		if strings.Contains(tmp, ".") {
			ipv4 := strings.Split(tmp, ".")
			if len(ipv4) == 4 {
				check := false
				for _, num := range ipv4 {
					n, _ := strconv.Atoi(num)
					if (len(num) < 4 && len(num) > 0) && n < 256 {
						if (num[0] == '0' && n < 8) || num[0] != '0' {
							check = true
							continue
						}
						check = false
					} else {
						check = false
					}
					if !check {
						break
					}
				}
				if check {
					result = append(result, "IPv4")
				} else {
					result = append(result, "Neither")
				}
			} else {
				result = append(result, "Neither")
				continue
			}
		} else if strings.Contains(tmp, ":") {
			if strings.Count(tmp, "::") > 1 {
				result = append(result, "Neither")
				continue
			}
			ipv6 := strings.Split(tmp, "::")
			tmpIPv6 := strings.Join(ipv6, ":")
			ipv6 = strings.Split(tmpIPv6, ":")
			if len(ipv6) <= 8 {
				check := false
				for _, num := range ipv6 {
					for j := 0; j < len(num); j++ {
						if (num[j] >= '0' && num[j] <= '9') || (num[j] >= 'a' && num[j] <= 'z') {
							check = true
						} else {
							check = false
						}
					}
				}
				if check {
					result = append(result, "IPv6")
				} else {
					result = append(result, "Neither")
				}
			} else {
				result = append(result, "Neither")
				continue
			}

		} else {
			result = append(result, "Neither")
			continue
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	addressesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var addresses []string

	for i := 0; i < int(addressesCount); i++ {
		addressesItem := readLine(reader)
		addresses = append(addresses, addressesItem)
	}

	result := validateAddresses(addresses)
	fmt.Print(result)

	//for i, resultItem := range result {
	//	fmt.Fprintf(writer, "%s", resultItem)
	//
	//	if i != len(result) - 1 {
	//		fmt.Fprintf(writer, "\n")
	//	}
	//}
	//
	//fmt.Fprintf(writer, "\n")
	//
	//writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
