// https://www.hackerrank.com/challenges/lilys-homework/problem
// status: 6/12 tests passes / 6/12 tests time out
// how to improve: check other algorithm (with less complexcity) or try to use concurrency
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the lilysHomework function below.
func lilysHomework(arr []int32) int32 {
    var minSwaps int32 = 0
    var maxSwaps int32 = 0

    maxArr := make([]int32, len(arr))
    copy(maxArr, arr)

    minArr := make([]int32, len(arr))
    copy(minArr, arr)

    for k := 0; k < len(arr) - 1; k++ {
        minIndex := k
        maxIndex := k
        minValue := minArr[k]
        maxValue := maxArr[k]

        for i := k + 1; i < len(arr); i++ {
            if minArr[i] < minValue {
                minIndex = i
                minValue = minArr[i]
            }
            if maxArr[i] > maxValue {
                maxIndex = i
                maxValue = maxArr[i]
            }
        }

        if minArr[k] != minArr[minIndex] {
            minArr[k], minArr[minIndex] = minArr[minIndex], minArr[k]
            minSwaps += 1
        }

        if maxArr[k] != maxArr[maxIndex] {
            maxArr[k], maxArr[maxIndex] = maxArr[maxIndex], maxArr[k]
            maxSwaps += 1
        }
    }
 
    if minSwaps < maxSwaps {
        return minSwaps
    }
    return maxSwaps
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

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := lilysHomework(arr)

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
