package main

import (
	"fmt"
	"math"

	"github.com/golang/geo/r3"
)

const input = 347991

func main() {
	fmt.Printf("Distance: %v\n", puzzle1(input))
	fmt.Printf("Value: %v\n", puzzle2(input))
}

func puzzle1(input int) int {
	pos := r3.Vector{}
	dir := r3.Vector{X: 1}
	current := 1

	for i := 0; ; i++ {
		mvt := i + 1
		for j := 1; j <= 2; j++ {
			if current <= input && (current+mvt) >= input {
				pos = pos.Add(dir.Mul(float64(input - current)))
				distance := math.Abs(pos.X) + math.Abs(pos.Y)
				return int(distance)
			} else {
				current += mvt
				pos = pos.Add(dir.Mul(float64(mvt)))
				dir = r3.Vector{X: -dir.Y, Y: dir.X}
			}
		}
	}
}

func puzzle2(input int) int {
	pos := r3.Vector{X: 2500, Y: 2500}
	dir := r3.Vector{X: 1}

	grid := [5000][5000]int{}

	grid[int(pos.X)][int(pos.Y)] = 1

	for i := 0; ; i++ {
		mvt := i + 1
		for j := 1; j <= 2; j++ {
			for k := 0; k < mvt; k++ {
				pos = pos.Add(dir)

				current := grid[int(pos.X+1)][int(pos.Y)]
				current += grid[int(pos.X+1)][int(pos.Y+1)]
				current += grid[int(pos.X+1)][int(pos.Y-1)]

				current += grid[int(pos.X-1)][int(pos.Y)]
				current += grid[int(pos.X-1)][int(pos.Y+1)]
				current += grid[int(pos.X-1)][int(pos.Y-1)]

				current += grid[int(pos.X)][int(pos.Y+1)]
				current += grid[int(pos.X)][int(pos.Y-1)]

				grid[int(pos.X)][int(pos.Y)] = current

				if current > input {
					return current
				}
			}

			dir = r3.Vector{X: -dir.Y, Y: dir.X}
		}
	}
}
