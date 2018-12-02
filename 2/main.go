package main

import (
	"fmt"

	"github.com/malafreniere/aoc2017/2/first"
	"github.com/malafreniere/aoc2017/2/second"
	"github.com/malafreniere/aoc2017/2/spreadsheet"
)

func main() {
	ss := spreadsheet.ParseFile("input.txt")

	fmt.Printf("Puzzle #1: %d\n", first.Execute(ss))
	fmt.Printf("Puzzle #2: %d", second.Execute(ss))
}
