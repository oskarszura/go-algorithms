// https://www.hackerrank.com/challenges/game-of-two-stacks/problem - good algorithm but too slow (timeouts)

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type node struct {
    children    []node
    value       int32
}

func (n *node) addChild(c node) {
    n.children = append(n.children, c)
}

func addBranch(n *node, a []int32, b []int32, max int32, level int32, maxLevel *int32) {
    if len(a) > 0 {
        lcVal := n.value + a[0]

        if lcVal <= max  {
            lc := node{
                children: []node{},
                value: n.value + a[0],
            }
            n.addChild(lc)

            if level > *maxLevel {
                *maxLevel = level
            }

            addBranch(&lc, a[1:], b, max, level + 1, maxLevel)
        }
    }

    if len(b) > 0 {
        rcVal := n.value + b[0]
        if rcVal <= max {
            rc := node{
                children: []node{},
                value: n.value + b[0],
            }
            n.addChild(rc)

            if level > *maxLevel {
                *maxLevel = level
            }

            addBranch(&rc, a, b[1:], max, level + 1, maxLevel)
        }
    }
}

func SearchDeepes(a []int32, b []int32, max int32, maxLevel *int32) {
    root := node{
        children: []node{},
        value: 0,
    }

    addBranch(&root, a, b, max, 1, maxLevel)
}

/*
 * Complete the twoStacks function below.
 */
func twoStacks(x int32, a []int32, b []int32) int32 {
    var maxLevel int32
    SearchDeepes(a, b, x, &maxLevel)
    return maxLevel
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    gTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    g := int32(gTemp)

    for gItr := 0; gItr < int(g); gItr++ {
        nmx := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nmx[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        mTemp, err := strconv.ParseInt(nmx[1], 10, 64)
        checkError(err)
        m := int32(mTemp)

        xTemp, err := strconv.ParseInt(nmx[2], 10, 64)
        checkError(err)
        x := int32(xTemp)

        aTemp := strings.Split(readLine(reader), " ")

        var a []int32

        for aItr := 0; aItr < int(n); aItr++ {
            aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
            checkError(err)
            aItem := int32(aItemTemp)
            a = append(a, aItem)
        }

        bTemp := strings.Split(readLine(reader), " ")

        var b []int32

        for bItr := 0; bItr < int(m); bItr++ {
            bItemTemp, err := strconv.ParseInt(bTemp[bItr], 10, 64)
            checkError(err)
            bItem := int32(bItemTemp)
            b = append(b, bItem)
        }

        result := twoStacks(x, a, b)

        fmt.Fprintf(writer, "%d\n", result)
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
