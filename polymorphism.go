package main 
import ("fmt")

func main(){
	x:=10
	y:=20
	func1(x)
	func1(x,y)
}

func func1(x,y int){
	fmt.Println("from 1",x,y)
}

func func1(x int){
	fmt.Println("from 2",x)
}