package main
import (
	"fmt"
)

func main(){
	var amount, total_amount float64
	fmt.Scanf("%f %f",&amount,&total_amount)
	if amount+0.5<total_amount && int(amount)%5==0 {
		fmt.Println(total_amount-amount-0.5)
	} else {
		fmt.Printf("%0.2f",total_amount)
	}

}