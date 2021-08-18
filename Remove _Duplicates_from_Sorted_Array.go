package main
import "fmt"
func removeDuplicates(nums []int) int {
    if len(nums)==0{
        return 0
    }
    i:=0
    fmt.Println(nums)
    for j:=1;j<len(nums);j++{

        if nums[j] != nums[i] {
            i++
            nums[i]=nums[j]
        }
        fmt.Println(nums)
    }
    return i+1
}

func main(){
	arr:=[]int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(arr))
}
