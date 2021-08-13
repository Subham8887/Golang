package main

import (
	"fmt"
)

func main(){
	p1:=struct{
		first string
		last string
		age int
	}{
		first: "subham",
		last: "Kumar",
		age: 22,
	}
	fmt.Println(p1)
}