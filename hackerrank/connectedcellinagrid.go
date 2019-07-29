// https://www.hackerrank.com/challenges/connected-cell-in-a-grid/problem
// 100 Solved!
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func walk(x int, y int, matrix [][]int32, visited [][]bool) int32 {
    var nearArea int32 = 0

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            dy := y + i
            dx := x + j

            if dx >= 0 && dy >= 0 && dy < len(matrix) && dx < len(matrix[dy]) {
                if visited[dy][dx] == false {
                    visited[dy][dx] = true

                    if matrix[dy][dx] == 1 {
                        nearArea += 1
                        nearArea += walk(dx, dy, matrix, visited)
                    }
                }
            }
        }   
    }

    return nearArea
}

// Complete the connectedCell function below.
func connectedCell(matrix [][]int32) int32 {
    var maxArea int32 = 0
    visited := make([][]bool, len(matrix))
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, len(matrix[i]))
    
        for j := 0; j < len(visited[i]); j++ {
            visited[i][j] = false
        }
    }

    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if visited[i][j] == false {
                if matrix[i][j] == 1 {
                    areaSize := walk(j, i, matrix, visited)

                    if areaSize > maxArea {
                        maxArea = areaSize
                    }
                }
            }
        }   
    }

    return maxArea
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    m := int32(mTemp)

    var matrix [][]int32
    for i := 0; i < int(n); i++ {
        matrixRowTemp := strings.Split(readLine(reader), " ")

        var matrixRow []int32
        for _, matrixRowItem := range matrixRowTemp {
            matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
            checkError(err)
            matrixItem := int32(matrixItemTemp)
            matrixRow = append(matrixRow, matrixItem)
        }

        if len(matrixRow) != int(int(m)) {
            panic("Bad input")
        }

        matrix = append(matrix, matrixRow)
    }

    result := connectedCell(matrix)

    fmt.Fprintf(writer, "%d\n", result)

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
