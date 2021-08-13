package main
import ("fmt")
func main(){
	arr:=[]int{1,2,3,4,5,6,7,8,9}
	s:=sum(arr...)
	fmt.Println(s)
	s2:=even(sum,arr...)
	fmt.Println(s2)
}

func sum(arr ...int) int {
	total:=0
	for _,val:= range arr{
		total +=val
	}
	return total
}

func even(f func(arr ...int) int, vi ...int) int {
	var yi []int
	
	for _,v :=range vi {
		if v%2==0{
			yi=append(yi,v)
		}
	}
	fmt.Println(yi) 
	return f(yi...)
}