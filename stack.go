package main
import "fmt"

type Stack struct{
	items []int
}

//Push
func (s *Stack) Push(i int){
	s.items=append(s.items,i)
}


//Pop
func (s *Stack) Pop() int {
	l:=len(s.items)-1
	removed_item:=s.items[l]
	s.items=s.items[:l]
	return removed_item
}


func main(){
	myStack:=Stack{}
	fmt.Println(myStack)
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}