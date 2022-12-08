package main

import (
	"fmt"
)

func main() {
	fmt.Println("[Day 05, part one] Top crates:", partOne(parse()))
	fmt.Println("[Day 05, part two] Top crates:", partTwo(parse()))
}

func partOne(actions []string, stacks [][]rune) string {
	for _, action := range actions {
		count, fromIdx, toIdx := parseAction(action)

		for item := count; item > 0; item-- {
			from := stacks[fromIdx]
			move := from[len(from)-1]
			stacks[fromIdx] = from[:len(from)-1]
			stacks[toIdx] = append(stacks[toIdx], move)
		}
	}

	return getTopCrates(stacks)
}

func partTwo(actions []string, stacks [][]rune) string {
	for _, action := range actions {
		count, fromIdx, toIdx := parseAction(action)
		from := stacks[fromIdx]
		to := stacks[toIdx]

		items := from[len(from)-count:]
		stacks[fromIdx] = from[0 : len(from)-count]
		stacks[toIdx] = append(to, items...)
	}

	return getTopCrates(stacks)
}
