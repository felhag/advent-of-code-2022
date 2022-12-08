package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"math"
)

func main() {
	arr := functions.Get(4)
	fmt.Println("[Day 04, part one]:", partOne(arr))
	fmt.Println("[Day 04, part two]:", partTwo(arr))
}

func partOne(arr []string) int {
	count := 0
	for _, assignments := range arr {
		idx1a, idx1b, idx2a, idx2b := parseAssignment(assignments)

		if (idx1a >= idx2a && idx1b <= idx2b) ||
			(idx1a <= idx2a && idx1b >= idx2b) {
			count++
		}
	}
	return count
}

func partTwo(arr []string) int {
	count := 0
	for _, assignments := range arr {
		idx1a, idx1b, idx2a, idx2b := parseAssignment(assignments)

		start := math.Max(idx1a, idx2a)
		end := math.Min(idx1b, idx2b)
		diff := end - start

		if diff >= 0 {
			count++
		}
	}
	return count
}
