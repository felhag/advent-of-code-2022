package main

import "advent-of-code-2022/functions"

func parseGrid(input []string) [][]int {
	var grid [][]int
	for _, row := range input {
		var gridRow []int
		for _, col := range row {
			gridRow = append(gridRow, functions.ParseInt(string(col)))
		}
		grid = append(grid, gridRow) // Append to the outer slice
	}
	return grid
}

func getColumn(grid [][]int, idx int) []int {
	var result []int
	for _, row := range grid {
		result = append(result, row[idx])
	}
	return result
}

func allLowerTrees(trees []int, height int) bool {
	for _, tree := range trees {
		if tree >= height {
			return false
		}
	}
	return true
}

func countVisibleTrees(trees []int, height int) int {
	result := 0
	for _, tree := range trees {
		result++
		if tree >= height {
			break
		}
	}
	return result
}

func reverseInts(input []int) []int {
	var result []int
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}
	return result
}
