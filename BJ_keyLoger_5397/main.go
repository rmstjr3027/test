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
	var testCase int
	var inputWord []byte
	fmt.Fscanf(r, "%d\n", &testCase)

	for i := 0; i < testCase; i++ {
		inputWord = nil
		var left, right []byte
		fmt.Fscanf(r, "%s\n", &inputWord)
		for _, c := range inputWord {
			switch c {
			case '<':
				if len(left) == 0 {
					continue
				}
				right = append(right, left[len(left)-1])
				left = left[:len(left)-1]
			case '>':
				if len(right) == 0 {
					continue
				}
				left = append(left, right[len(right)-1])
				right = right[:len(right)-1]
			case '-':
				if len(left) == 0 {
					continue
				}
				left = left[:len(left)-1]
			default:
				left = append(left, c)
			}
		}
		for i, j := 0, len(right)-1; i < j; i, j = i+1, j-1 {
			right[i], right[j] = right[j], right[i]
		}
		fmt.Fprintf(w, "%s%s\n", left, right)
	}
}

//type Stack struct {
//	top int
//	items []interface{}
//}
//
//func (s *Stack) push (item interface{}) {
//	s.items = append(s.items, item)
//	s.top++
//}
//
//func (s *Stack) pop () interface{} {
//	if s.top == 0 {
//		return nil
//	}
//	s.top--
//	item := s.items[s.top]
//	s.items = s.items[:s.top]
//	return item
//}
//
//func reversString(s string) string {
//	runes := []rune(s)
//	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//	return string(runes)
//}
//
//func main() {
//	defer w.Flush()
//	var testCase int
//	var inputWord string
//	fmt.Fscanf(r, "%d\n", &testCase)
//
//	for i := 0; i < testCase; i++ {
//		inputWord = ""
//		fmt.Fscanf(r, "%s\n", &inputWord)
//		var passWordStack0, passWordStack1  Stack
//		var passWord bytes.Buffer
//		for i := 0; i < len(inputWord); i++ {
//			if string(inputWord[i]) == "<" {
//				if passWordStack0.top > 0 {
//					passWordStack1.push(passWordStack0.pop().(string))
//				}
//			} else if string(inputWord[i]) == ">" {
//				if passWordStack1.top > 0 {
//					passWordStack0.push(passWordStack1.pop().(string))
//				}
//			} else if string(inputWord[i]) == "-" {
//				passWordStack0.pop()
//			} else {
//				passWordStack0.push(string(inputWord[i]))
//			}
//		}
//		for passWordStack0.top > 0 {
//			passWord.WriteString(passWordStack0.pop().(string))
//		}
//		str := reversString(passWord.String())
//		passWord.Reset()
//		passWord.WriteString(str)
//		for passWordStack1.top > 0 {
//			passWord.WriteString(passWordStack1.pop().(string))
//		}
//		fmt.Fprintln(w, passWord.String())
//	}
//
//}
