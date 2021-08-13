package main

import("fmt")

func main(){
	s:=func1()
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
}

func func1() func() int {
	i:=0
	return func () int {
		i+=1
		return i
	}
}