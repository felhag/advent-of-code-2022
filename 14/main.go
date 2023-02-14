package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("[Day 14, part one] Pieces of sand: %d\n", part1(draw(functions.Get(14))))
	fmt.Printf("[Day 14, part two] Pieces of sand: %d\n", part2(draw(functions.Get(14))))
}

func part1(grid [][]rune) int {
	done, count := false, 0
	for ok := true; ok; ok = !done {
		done = dropSand(grid)
		if !done {
			count++
		}
	}
	//drawGrid(grid)
	return count
}

func part2(grid [][]rune) int {
	grid[len(grid)-1] = []rune(strings.Repeat(string('#'), 1000))
	count := 0
	for ok := true; ok; ok = grid[0][500] != 'o' {
		dropSand(grid)
		count++
	}
	//drawGrid(grid)
	return count
}

func drawGrid(grid [][]rune) {
	minX := 500
	maxX := 500

	for _, row := range grid[0 : len(grid)-2] {
		for idx, letter := range row {
			if letter != '.' && minX > idx {
				minX = idx
			} else if letter != '.' && maxX < idx {
				maxX = idx
			}
		}
	}

	fmt.Println("\nGrid:")
	for _, row := range grid {
		fmt.Println(string(row[minX-3 : maxX+4]))
	}
}
