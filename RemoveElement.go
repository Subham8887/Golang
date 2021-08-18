package main
import "fmt"



func removeElement(nums []int, val int) int {
    if len(nums)==0{
		return 0
	}
	// fmt.Println(nums)
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[j]==val{
			nums = append(nums[:j],nums[j+1:]...)
			nums=append(nums, val)
		} else {
			j++
		}
		// fmt.Println(nums)
	}
	return j
}

func main(){
	arr := []int{0,1,2,2,3,0,4,2}
	val := 2
	fmt.Println(removeElement(arr,val))
	// fmt.Println(arr[:2])
}