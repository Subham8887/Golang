package main

import ("fmt")

func main(){
	arr:=[]int{1,2,3,4,5,6,7,8,9}
	s:=sum(arr...)
	fmt.Println("sum = ",s)
}

func sum(x ...int) int {
	sum:=0
	for _, v:=range(x){
		sum+=v
	}
	return sum
}