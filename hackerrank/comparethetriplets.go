// https://www.hackerrank.com/challenges/compare-the-triplet

package main

import (
    "fmt"
    "bufio"    
    "os"
    "strconv"
)

func main() {
    var (
        alice [3]int
        bob [3]int
        aliceTotal int
        bobTotal int
    )
    
    aliceTotal = 0
    bobTotal = 0
   
    io := bufio.NewReader(os.Stdin)
    
    for i := 0; i < 3; i++ {
        fmt.Fscan(io, &alice[i])
    }
    
    for i := 0; i < 3; i++ {
        fmt.Fscan(io, &bob[i])
    }
    
    for i := 0; i < 3; i++ {
        if alice[i] > bob[i] {
            aliceTotal += 1
        } else if bob[i] > alice[i] {
            bobTotal += 1
        } 
    }
    
    fmt.Println(strconv.Itoa(aliceTotal) + " " + strconv.Itoa(bobTotal))
}
