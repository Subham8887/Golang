package main
import ("fmt")
func main(){
	// defer func1() ///prints aftewords
	// defer func2()
	// func3()
	x:=1
	y:=&x

	defer func4(y)
	defer func5(y)
	func6(y)
}

// func func1(){
// 	fmt.Println("from func 1")
// }

// func func2(){
// 	fmt.Println("from func 2")
// }
// func func3(){
// 	fmt.Println("from func 3")
// }

func func4(y *int){
	*y+=1
	fmt.Println("from func 4",*y)
}

func func5(y *int){
	*y+=1
	fmt.Println("from func 5",*y)
}
func func6(y *int){
	*y+=1
	fmt.Println("from func 6",*y)
}