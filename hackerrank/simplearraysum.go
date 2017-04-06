// https://www.hackerrank.com/challenges/simple-array-sum

package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var (
        n int
        sum int
    )
    
    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &n)
    
    arr := make([]int, n)
    sum = 0

    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
        sum += arr[i]
    }
    
    fmt.Println(sum)
}
