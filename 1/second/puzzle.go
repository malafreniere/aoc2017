package second

import (
	"math"
)

func Execute(inputs []int) int {
	total := 0
	length := float64(len(inputs))
	delta := length / 2
	for i, v := range inputs {
		j := int(math.Mod(float64(i)+delta, length))

		if v == inputs[j] {
			total += v
		}
	}

	return total
}
