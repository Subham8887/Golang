package main

import ("fmt")

func main(){
	s:=`Hello world`
	fmt.Println(s)

	bs:= []byte(s)    // it return [72 101 108 108 111 32 119 111 114 108 100]
	fmt.Println(bs)

	for i:=0;i<len(s);i++{
		fmt.Printf("%#U ",s[i])
	}
	print("\n")

	for i,v:=range s{
		fmt.Println(i,v)
	}
}