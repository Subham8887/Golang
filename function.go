package main

import (
	"fmt"
)

func main(){
	func1()
	func2("subham")
	s:=func3("akash")
	fmt.Println(s)
	x,y:= func4(2,6)
	fmt.Println(x,y)
}

func func1(){
	fmt.Println("msg from func1")
}

func func2(name string){
	fmt.Println("message from func 2",name)
}

func func3(name string) string {
	return fmt.Sprint("msg from func 3 ", name)
}

func func4(x int,y int) (int,int){
	return y,x
}
