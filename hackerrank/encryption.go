// https://www.hackerrank.com/challenges/encryption/problem
// status: 100% passes

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "math"
)

// Complete the encryption function below.
func encryption(s string) string {
    stripped := strings.Replace(s, " ", "", -1)
    L := len(stripped)
    sqrtL := math.Sqrt(float64(L))
    lL := int(math.Floor(sqrtL))
    rL := int(math.Ceil(sqrtL))

    capacity := lL * rL
    cols := rL
    rows := rL

    if capacity >= L {
        cols = rL
        rows = lL
    }

    counter := 0
    grid := make([][]string, rows)
    for y, _ := range grid {
        grid[y] = make([]string, cols)

        for x, _ := range grid[y] {
            if counter < L {
                grid[y][x] = string([]rune(s)[counter])
                counter += 1
            } else {
                break;
            }
        }
    }

    var output string
    for x := 0; x < cols; x++ {
        for y := 0; y < rows; y++ {
            if x + y < L {
                output += grid[y][x]
            } else {
                break;
            }
        }
        output += " "
    }

    return output
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
