package main

import (
	"fmt"
)

func main() {
	fmt.Println(puzzle())
}

func puzzle() (int, int) {
	input := [...]int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	// input := [...]int{0, 2, 7, 0}
	seen := make(map[string]int)
	redistribution := 0

	for {
		i, m := findMaxIndex(input[:])

		input[i] = 0

		for j := i + 1; m > 0; j++ {
			if j == len(input) {
				j = -1
				continue
			}

			m--
			input[j]++
		}

		redistribution++
		f := fmt.Sprintf("%v", input)

		if cycles, ok := seen[f]; ok {
			return redistribution, (redistribution - cycles)
		} else if !ok {
			seen[f] = redistribution
		}
	}
}

func findMaxIndex(blocks []int) (i int, m int) {
	i = 0
	m = -1

	for index, val := range blocks {
		if val > m {
			i = index
			m = val
		}
	}

	return
}
