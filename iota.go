package main

import (
	"fmt"
)

const (
	a=iota
	b=iota
	c
	d
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

go rutine,interfaces, method