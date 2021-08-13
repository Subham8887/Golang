package main
import (
	"fmt"
)

func main(){
	m:=map[string]int{
		"subham":22,
		"kumar":23,
	}
	fmt.Println(m["subham"])
	fmt.Println(m)
	fmt.Println(m["ujjwal"]) //returns 0 if not present

	fmt.Println()

	v,ok:=m["subham"]

	fmt.Println(v,ok)

	delete(m,"subham")
	fmt.Println(m)
}