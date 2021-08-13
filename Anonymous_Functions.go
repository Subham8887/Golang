package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	//Type 1
	fmt.Println(area(20, 30))

	//Type 2
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	//Type 3
	fmt.Printf("100 (°F) = %.2f (°C)\n",
					func(f float64) float64 {
						return (f - 32.0) * (5.0 / 9.0)
					}(100),
				)
}