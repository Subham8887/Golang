package main

import (
    "os"
    "bufio"
    "fmt"
)

func main(){
    var n, k, c int
    
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    
    fmt.Sscanf(scanner.Text(), "%d %d", &n, &k)
    
    for ;n > 0; n-- {
        scanner.Scan()
        if toInt(scanner.Bytes())%k == 0 {
            c++
        }
    }
    
    fmt.Println(c)
}

func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}