// https://www.hackerrank.com/challenges/arrays-ds

package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)


func main() {
    var (
        n int
    )
    
    io := bufio.NewReader(os.Stdin)
    
    fmt.Fscan(io, &n)
    a := make([]int, n)
        
    for i := 0; i < n; i++ {
        fmt.Fscan(io, &a[i])
    }
    
    for i := n - 1; i >= 0; i-- {
        if i == n-1 {
            fmt.Print(strconv.Itoa(a[i]))
        } else {
            fmt.Print(" " + strconv.Itoa(a[i]))
        }
    }
}
