package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := functions.Get(7)
	structure := parse(input)

	fmt.Println("[Day 07, part one] Summed size:", findSmallFolders(structure))
	fmt.Println("[Day 07, part two] Deleted directory size:", findFolderToDelete(structure))
}

func findSmallFolders(structure FolderStructure) int {
	sum := 0
	for s := range structure {
		if strings.Contains(s, "dir, size=") {
			size := getFolderSize(s)

			if size < 100000 {
				sum += size
			}

			sum += findSmallFolders(structure[s])
		}
	}
	return sum
}

func findFolderToDelete(structure FolderStructure) int {
	for root, value := range structure {
		size := getFolderSize(root)
		needed := 30000000 - (70000000 - size)
		fmt.Println("Needed: 30000000 - ( 70000000 -", size, ") =", needed)
		smallest := findRecursive(value, needed)
		return smallest
	}
	return 0
}

func findRecursive(structure FolderStructure, needed int) int {
	children := []int{}
	for s := range structure {
		if strings.Contains(s, "dir, size=") {
			size := getFolderSize(s)
			if size < needed {
				continue
			}

			childSize := findRecursive(structure[s], needed)
			children = append(children, size)
			if childSize > 0 {
				children = append(children, childSize)
			}
		}
	}
	sort.Ints(children)
	if len(children) > 0 {
		return children[0]
	} else {
		return 0
	}
}

func getFolderSize(folder string) int {
	s := strings.Split(folder, "=")[1]
	return functions.ParseInt(s[0 : len(s)-1])
}
