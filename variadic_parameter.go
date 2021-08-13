package main

import (
	"fmt"
)

func main(){
	func1(1,2,3,4,5,6,7)
}

func func1(x ...int){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	sum:=0
	for _,v := range(x){
		sum+=v
	}

	fmt.Println("sum=",sum)
}