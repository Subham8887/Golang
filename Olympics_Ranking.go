package main
import (
	"fmt"
)

func main(){
	
	var n,G1,S1,B1,G2,S2,B2 int
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Scan(&G1,&S1,&B1,&G2,&S2,&B2)
		if G1+S1+B1 > G2+S2+B2{
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
		
	}
	
}