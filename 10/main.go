package main

import (
	"advent-of-code-2022/functions"
	"fmt"
)

func main() {
	input := functions.Get(10)

	fmt.Println("[Day 10, part one] Total sum:", executePart1(input))
	fmt.Println("[Day 10, part two]:")
	executePart2(input)
}

func executePart1(input []string) int {
	todos := input
	pending := 0
	x := 1
	cycle := 0
	total := 0
	for ok := true; ok; ok = len(todos) > 0 || pending != 0 {
		cycle++
		fmt.Printf("Starting cycle %d. Current x: %d, pending: %d\n", cycle, x, pending)

		if cycle%40 == 20 {
			fmt.Printf("Adding total: %d += %d * %d = %d\n", total, cycle, x, cycle*x)
			total += cycle * x
		}

		pending, x, todos = tick(pending, x, todos)

		fmt.Printf("Finished cycle %d. Current x: %d\n", cycle, x)
	}

	return total
}

func executePart2(input []string) {
	todos := input
	pending := 0
	x := 1
	cycle := 0
	for ok := true; ok; ok = len(todos) > 0 || pending != 0 {
		pos := cycle % 40
		cycle++

		if pos == 39 {
			fmt.Print("\n")
		} else if x-pos >= -1 && x-pos <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		pending, x, todos = tick(pending, x, todos)
	}
}
