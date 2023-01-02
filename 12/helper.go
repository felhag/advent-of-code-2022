package main

import "syscall"

type Position struct {
	X int
	Y int
}

func parse(input []string) ([][]rune, Position, Position) {
	var start Position
	var end Position
	var grid [][]rune
	for rowIdx, row := range input {
		var gridRow []rune
		for colIdx, col := range row {
			if col == 'S' {
				start = Position{X: rowIdx, Y: colIdx}
				col = 'a'
			} else if col == 'E' {
				end = Position{X: rowIdx, Y: colIdx}
				col = 'z'
			}
			gridRow = append(gridRow, col)
		}
		grid = append(grid, gridRow)
	}
	return grid, start, end
}

func createSums(grid [][]rune) [][]int {
	flags := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		flags[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			flags[i][j] = syscall.INFINITE
		}
	}
	return flags
}
