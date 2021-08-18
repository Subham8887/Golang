package main
import "fmt"

func generate(numRows int) [][]int {
    var matrix [][]int
	fmt.Println(matrix)
	for line:=1;line<=numRows;line++{
		c:=1
		var temp []int
		for i:=1;i<line+1;i++ {
			temp=append(temp,c)
			c = int(c * (line - i) / i)
		}
		matrix=append(matrix, temp)
	}
	return matrix
}

func main(){
	n:=5
	fmt.Println(generate(n))
}