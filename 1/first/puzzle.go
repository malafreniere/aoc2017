package first

import (
	"math"
)

func Execute(inputs []int) int {
	total := 0
	length := float64(len(inputs))

	for i, v := range inputs {
		j := int(math.Mod(float64(i+1), length))

		if v == inputs[j] {
			total += v
		}
	}

	return total
}
