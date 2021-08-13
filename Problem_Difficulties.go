package main
import (
	"fmt"
)
type Counter struct {
	count int
}
func (self Counter) currentValue() int {
	return self.count
}
func (self Counter) increment() {
	self.count++
}
func main(){
	
	var n,A1,A2,A3,A4,max int
	
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		max=0
		fmt.Scan(&A1,&A2,&A3,&A4)
		inputArray := []int{A1,A2,A3,A4}
		dict:= make(map[int]int)
		for _ , num :=  range inputArray {
			dict[num] = dict[num]+1
    	}
		for _,k:= range dict{
			if k>max{
				max=k
			}
		}
		if max==4{
			fmt.Println(0)
		} else if max ==3{
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
	
}

