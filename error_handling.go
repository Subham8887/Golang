package main
import (
	"fmt"
	"errors"
)

func main(){
	total,err:=sum(10,9)
	if err!=nil{
		fmt.Println(err)
	} else {
		fmt.Println(total)
	}
}


func sum(start,end int) (int,error){
	if start > end {
		return 0,errors.New("Start is greater then end")
	}
	total:=0
	for i:=start;i<=end;i++{
		total+=i
	}
	return total,nil
}