package main

import (
	"fmt"
)

func main() {
	arrayToSort := []int{5, 4, 5, 6, 2, 1, 543, 31, 123, 453, 12}
	swaps := 1
	for swaps > 0 {
		swaps = 0
		for i := 0; i < len(arrayToSort) - 1; i++ {
			buff := arrayToSort[i+1]
			curr := arrayToSort[i]
			if curr > buff {
				arrayToSort[i] = buff
				arrayToSort[i+1] = curr
				swaps = swaps + 1
			}
		}

	}
	fmt.Println(arrayToSort)
}