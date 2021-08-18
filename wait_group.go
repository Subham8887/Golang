package main

import(
	"fmt"
	// "runtime"
	"sync"
)

var wg sync.WaitGroup

func main(){
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("jhgkhgkhgkgkgkh")
	wg.Wait()
}

func foo(){
	for i:=0;i<=10;i++{
		fmt.Println("foo")
	}
	wg.Done()
}

func bar(){
	for i:=0;i<=10;i++{
		fmt.Println("bar")
	}
	
}