package main

import(
	"fmt"
)
var a = `akash"subham kumar"`
var b= `bawkaldkjj 
akdfj
ajkdfh`
var c int =34
var d string ="subham Kumar"
var z int
func main(){
	s:=4
	fmt.Printf("%T %v\n",s,s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(z)
	e:=fmt.Sprintf("%v\t%v",c,d)
	fmt.Println(e)	
}