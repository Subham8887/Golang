package main

import (
	"fmt"
)

func main(){
	var x=[]int{8,2,3,4,5}
	// var y =[]int// wrong expression
	fmt.Println(x)
	for _,b:=range x{
		fmt.Println(b)
	}
}