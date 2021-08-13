package main
import ("fmt")

func main(){
	var n,a,k int
	fmt.Scan(&n)
	
	for i:=0;i<n;i++ {
		fmt.Scan(&a)
		k=a/3
		if (k)+2*(k)==a{
			fmt.Println(k,k)
		} else {
			if (k+1)+2*k ==a{
				fmt.Println(k+1,k)
			} else {
				fmt.Println(k,k+1)
			}
		}
	}
}