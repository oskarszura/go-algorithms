// https://www.hackerrank.com/challenges/a-very-big-sum

package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var (
        n int
        sum int64
    )
    
    io := bufio.NewReader(os.Stdin)
    
    fmt.Fscan(io, &n)
    arr := make([]int64, n)
    
    for i := 0; i < n; i++ {
        fmt.Fscan(io, &arr[i])
        sum += arr[i]
    }
    
    fmt.Println(sum)
}
