package main
import("fmt")
var a int
type custom_data_type int
var b custom_data_type
func main(){
 a=30
 b=40
 fmt.Println(a)
 fmt.Printf("%T\n",a)
 fmt.Println(b)
 fmt.Printf("%T\n",b)
//  a=b  # cannot do this bcz types are diffrent
}