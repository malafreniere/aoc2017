package main

import (
	"fmt"

	"github.com/malafreniere/aoc2017/1/captcha"
	"github.com/malafreniere/aoc2017/1/first"
	"github.com/malafreniere/aoc2017/1/second"
)

func main() {
	inputs := captcha.ParseFile("input.txt")

	fmt.Printf("Puzzle #1: %d\n", first.Execute(inputs))
	fmt.Printf("Puzzle #2: %d", second.Execute(inputs))
}
