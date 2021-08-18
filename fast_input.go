package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {

	var n,k,ans,i,input int
	scanf("%d", &n)
	scanf("%d\n", &k)

	for i<n {
		scanf("%d\n", &input)
		if input%k==0 {
			ans++
		}
		i++
	}
	fmt.Println(ans)
}
