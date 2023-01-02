package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"syscall"
)

func main() {
	fmt.Printf("[Day 12, part one] Shortest path: %d\n", part1())
	fmt.Printf("[Day 12, part two] Shortest path: %d\n", part2())
}

func part1() int {
	grid, start, end := parse(functions.Get(12))
	sums := createSums(grid)
	step(grid, sums, 0, []Position{end}, start)
	return sums[start.X][start.Y]
}

func part2() int {
	grid, start, end := parse(functions.Get(12))
	sums := createSums(grid)
	step(grid, sums, 0, []Position{end}, start)

	min := syscall.INFINITE
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'a' && sums[i][j] < min {
				min = sums[i][j]
			}
		}
	}
	return min
}

func step(grid [][]rune, sums [][]int, sum int, path []Position, end Position) {
	pos := path[len(path)-1]
	sums[pos.X][pos.Y] = sum

	if pos.X-1 > -1 && grid[pos.X-1][pos.Y]-grid[pos.X][pos.Y] > -2 {
		if sums[pos.X-1][pos.Y]-1 > sum {
			step(grid, sums, sum+1, append(path, Position{X: pos.X - 1, Y: pos.Y}), end)
		}
	}

	if pos.X+1 < len(grid) && grid[pos.X+1][pos.Y]-grid[pos.X][pos.Y] > -2 {
		if sums[pos.X+1][pos.Y]-1 > sum {
			step(grid, sums, sum+1, append(path, Position{X: pos.X + 1, Y: pos.Y}), end)
		}
	}

	if pos.Y-1 > -1 && grid[pos.X][pos.Y-1]-grid[pos.X][pos.Y] > -2 {
		if sums[pos.X][pos.Y-1]-1 > sum {
			step(grid, sums, sum+1, append(path, Position{X: pos.X, Y: pos.Y - 1}), end)
		}
	}

	if pos.Y+1 < len(grid[pos.X]) && grid[pos.X][pos.Y+1]-grid[pos.X][pos.Y] > -2 {
		if sums[pos.X][pos.Y+1]-1 > sum {
			step(grid, sums, sum+1, append(path, Position{X: pos.X, Y: pos.Y + 1}), end)
		}
	}
}
