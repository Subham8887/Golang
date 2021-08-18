package main
import (
	"fmt"
	"sort"
)

func kthLargestNumber(nums []string, k int) string {
	arr:=[]int{}
	for i:=0;i<len(nums);i++{
		arr = append(arr, int(nums[i]))
	}
    sort.Strings(arr)
	fmt.Println(nums)
    return nums[k-1]
}

func main(){
	arr:=[]string{"3","6","7","10"}
	k:= 4
	fmt.Println(kthLargestNumber(arr,k))

}