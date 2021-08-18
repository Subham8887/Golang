package IntStack

import "fmt"

type Stack []int

func (s *Stack) IsEmpty() bool{
	return len(*s)==0
}

func (s *Stack) Push(val int) {
	*s = append(*s , val)
} 

func (s *Stack) Pop() (int,bool) {
	if s.IsEmpty() {
		return 0,false
	} else {
		index := len(*s) -1
		elementPoped := (*s)[index]
		*s= (*s)[:index]
		return elementPoped,true
	}
}

func intStack() *Stack {
	s:=Stack{}
	return &s
}