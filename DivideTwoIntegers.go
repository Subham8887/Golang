package main
import ("fmt")

func Abs(x int) int {
	if x<0 {
		return -1*x
	}
	return x
}


func divide(dividend int, divisor int) int {
	var sign bool
	if (dividend<0 || divisor<0) && !(dividend<0 && divisor<0) {
		sign=true
	} else {
		sign=false
	}

	dividend=Abs(dividend)
	divisor=Abs(divisor)
	ans:=0
	for dividend >= divisor {
		ans+=1
		dividend-=divisor
	}
	if sign == true{
		ans=-1*ans
	}
	return ans
	
}

func main(){
	dividend := -2147483648
	divisor := -1
	fmt.Println(divide(dividend,divisor))
}