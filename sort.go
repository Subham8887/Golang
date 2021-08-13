package main
import (
	"fmt"
	"sort"
)

func main(){
	arr:=[]int{4,2,6,1,7,4,2,8,10}
	sort.Ints(arr)
	fmt.Println("sorted ",arr )
	sort.Sort(sort.Reverse(sort.Ints(arr)))
	fmt.Println()


	arr2:=[]string{"subham","shukla","saloni"}
	sort.Strings(arr2)
	fmt.Println(arr2)
}