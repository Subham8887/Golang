package main

import ("fmt")

func main(){
	x:=[]int{10,20,30}
	x=append(x, 40,50,60)
	y:=[]int{100,200,300}
	x=append(x,y...)
	fmt.Println(x)

	//Deleting

	x=append(x[:2], x[4:]...)
	fmt.Println(x)
}