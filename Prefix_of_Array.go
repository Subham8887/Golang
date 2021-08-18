package main
import "fmt"

func isPrefixString(s string, words []string) bool {
    var a string
	for _,v :=range words{
		a=a+v
	}
	if len(s)>len(a){
		return false
	} else {
		if s == a[:len(s)] {
			return true
		} else {
			return false
		}
	}
}

func main(){
	s:="iloveleetcode"
	words:=[]string{"i","love","leetcode","apples"}
	if isPrefixString(s,words){
		fmt.Println("true")
	} else{
		fmt.Println("flase")
	}
}