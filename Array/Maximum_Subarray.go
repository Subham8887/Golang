package main
import (
	"fmt")

func Max(x,y int) int{
	if x>y{
		return x
	} else {
		return y
	}
}

func maxSubArray(nums []int) int {
	l:=len(nums)
	if l==0{
		return 0
	}
	global_maximum:=-1000000000000
	temp_maximum:=-1000000000000
	for i:=0;i<l;i++{
		temp_maximum+=nums[i]
		temp_maximum= Max(temp_maximum,nums[i])
		global_maximum = Max(global_maximum,temp_maximum)
	}
	return global_maximum
}

func main(){
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}