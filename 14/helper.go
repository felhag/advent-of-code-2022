package main

import (
	"advent-of-code-2022/functions"
	"strings"
)

type Position struct {
	X int
	Y int
}

func parse(rows []string) ([][]Position, int) {
	var positions [][]Position
	maxY := 0
	for _, row := range rows {
		var gridRow []Position
		for _, position := range strings.Split(row, " -> ") {
			pos := strings.Split(position, ",")
			x := functions.ParseInt(pos[0])
			y := functions.ParseInt(pos[1])

			if y > maxY {
				maxY = y
			}

			gridRow = append(gridRow, Position{X: x, Y: y})
		}

		positions = append(positions, gridRow) // Append to the outer slice
	}
	return positions, maxY + 3
}

func draw(rows []string) [][]rune {
	positions, maxY := parse(rows)
	var grid [][]rune
	for i := 0; i < maxY; i++ {
		grid = append(grid, []rune(strings.Repeat(string('.'), 1000)))
	}

	grid[0][500] = '+'
	for _, row := range positions {
		for i := 0; i <= len(row)-2; i++ {
			start := row[i]
			end := row[i+1]
			if start.X < end.X {
				drawLineX(grid[start.Y], start.X, end.X)
			} else if start.X > end.X {
				drawLineX(grid[start.Y], end.X, start.X)
			} else if start.Y < end.Y {
				drawLineY(grid, start.X, start.Y, end.Y)
			} else if start.Y > end.Y {
				drawLineY(grid, start.X, end.Y, start.Y)
			}
		}
	}

	return grid
}

func drawLineX(line []rune, start int, end int) {
	for i := start; i <= end; i++ {
		line[i] = '#'
	}
}

func drawLineY(grid [][]rune, x int, start int, end int) {
	for i := start; i <= end; i++ {
		grid[i][x] = '#'
	}
}

func dropSand(grid [][]rune) bool {
	return moveSand(grid, 500, 0)
}

func moveSand(grid [][]rune, x int, y int) bool {
	if y >= len(grid)-1 {
		return true
	}
	nextRow := grid[y+1]
	if nextRow[x] == '.' {
		return moveSand(grid, x, y+1)
	} else if nextRow[x-1] == '.' {
		return moveSand(grid, x-1, y+1)
	} else if nextRow[x+1] == '.' {
		return moveSand(grid, x+1, y+1)
	} else {
		grid[y][x] = 'o'
		return false
	}
}
