package main
import (
	"fmt"
	"math"
	"sort"
)
func Min(a,b int) int {
	if a>b {
		return b
	} else {
		return a
	}
}

func minimumDifference(nums []int, k int) int {
    ans:=math.MaxInt32
	sort.Ints(nums)
	for i:=0;i<len(nums)-k+1;i++{
		ans=Min(ans, nums[i+k-1] - nums[i])
	}
	return ans
}
//1 4 7 9

func main(){
	arr:=[]int{9,4,1,7}
	k:=2
	fmt.Println(minimumDifference(arr,k))
}