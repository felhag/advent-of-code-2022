package main

import (
	"advent-of-code-2022/functions"
	"fmt"
)

func main() {
	input := functions.Get(8)

	grid := parseGrid(input)
	fmt.Println("[Day 08, part one] Visible trees:", findVisibleTrees(grid))
	fmt.Println("[Day 08, part two] Scenic tree:", findScenicTree(grid))
}

func findVisibleTrees(grid [][]int) int {
	count := 0
	for rowIdx, row := range grid {
		for colIdx, tree := range row {
			if allLowerTrees(row[0:colIdx], tree) || allLowerTrees(row[colIdx+1:], tree) {
				count++
			} else {
				column := getColumn(grid, colIdx)
				if allLowerTrees(column[0:rowIdx], tree) || allLowerTrees(column[rowIdx+1:], tree) {
					count++
				}
			}
		}
	}
	return count
}

func findScenicTree(grid [][]int) int {
	result := 0
	for rowIdx, row := range grid {
		for colIdx, tree := range row {
			column := getColumn(grid, colIdx)
			visible := countVisibleTrees(reverseInts(row[0:colIdx]), tree) *
				countVisibleTrees(row[colIdx+1:], tree) *
				countVisibleTrees(reverseInts(column[0:rowIdx]), tree) *
				countVisibleTrees(column[rowIdx+1:], tree)

			fmt.Printf("Visible from %d [%d,%d]: %d\n", tree, rowIdx, colIdx, visible)

			if visible > result {
				result = visible
			}
		}
	}
	return result
}
