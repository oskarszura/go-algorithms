// https://www.hackerrank.com/challenges/maxsubarray/problem
// 4/5 passes - 1 is timeouted - complexity can be reduced with Kadane's algorithm

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the maxSubarray function below.
func maxSubarray(arr []int32) []int32 {
    resp := []int32{0, 0}
    var maxEl int32 = -10
    var prev int32
    var acc int32 = 0
    var reduced []int32

    for i := 0; i < len(arr); i++ {
        if arr[i] > maxEl {
            maxEl = arr[i]
        }

        if i > 0 {
            if prev >= 0 && arr[i] < 0 || prev < 0 && arr[i] > 0 {
                reduced = append(reduced, acc)
                acc = 0
            }
        }

        prev = arr[i]
        acc += arr[i]

        if i == len(arr) - 1 {
            reduced = append(reduced, acc)
        }
    }

    if maxEl < 0 {
        resp[0] = maxEl
        resp[1] = maxEl
        return resp
    }

    var maxArr int32 = 0

    for i := 0; i < len(reduced); i++ {
        acc = 0
        for j := i; j < len(reduced); j++ {
            acc += reduced[j]

            if acc > maxArr {
                maxArr = acc
            }
        }
    }
    resp[0] = maxArr

    for i := 0; i < len(arr); i++ {
        if maxEl > 0 && arr[i] > 0 {
            resp[1] += arr[i]
        }
    }

    return resp
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        arrTemp := strings.Split(readLine(reader), " ")

        var arr []int32

        for i := 0; i < int(n); i++ {
            arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arr = append(arr, arrItem)
        }

        result := maxSubarray(arr)

        for i, resultItem := range result {
            fmt.Fprintf(writer, "%d", resultItem)

            if i != len(result) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        fmt.Fprintf(writer, "\n")
    }

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
