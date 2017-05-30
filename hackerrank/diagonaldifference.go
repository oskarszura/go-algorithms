// https://www.hackerrank.com/challenges/diagonal-difference

package main
import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func main() {
    var (
        n int
        primaryDiag int
        secondaryDiag int
    )
    
    primaryDiag = 0
    secondaryDiag = 0
    
    io := bufio.NewReader(os.Stdin)
    
    fmt.Fscan(io, &n)
    matrix := make([]int, n * n)
    
    for c := 0; c < n * n; c++ {
        fmt.Fscan(io, &matrix[c])
    }
    
    for c:= 0; c < n; c++ {
        pi := c
	    pj := c
	    si := n - c - 1
	    sj := c
	
	    pc := sj * n + si
	    sc := pj * n + pi
      
        primaryDiag += matrix[pc]
        secondaryDiag += matrix[sc]        
    }
    
    fmt.Println(math.Abs(float64(primaryDiag - secondaryDiag)))
}
