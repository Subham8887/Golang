package main
import (
	"fmt"
	// "math"
)
func Max(x,y int) int{
	if x>y {
		return x
	} else {
		return y
	} 
}

func Abs(x int) int{
	if x<0{
		return -1*x
	}
	return x
}

func maxProfit(prices []int) int {
	l:=len(prices)
	m:=0
    for i:=0;i<l-1;i++{
		for j:=i+1;j<l;j++{
			m=Max(m,Abs(prices[i]-prices[j]))
		}
	}
	return m
}

func main(){
	arr:=[]int{7,1,5,3,6,4}
	fmt.Println(maxProfit(arr))
}

