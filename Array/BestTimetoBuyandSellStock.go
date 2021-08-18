package main
import (
	"fmt"
	"math"
)
func maxProfit(prices []int) int {
	l:=len(prices)
	minval:=math.MaxInt32
	maxprofit:=0
    for i:=0;i<l;i++{
		if minval>prices[i]{
			minval=prices[i]
		} else if prices[i]-minval>maxprofit {
			maxprofit=prices[i]-minval
		} 
	}
	return maxprofit
}
func main(){
	arr:=[]int{1,2}
	fmt.Println(maxProfit(arr))
}

