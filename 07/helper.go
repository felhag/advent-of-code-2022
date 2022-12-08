package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FolderStructure map[string]FolderStructure

func (structure FolderStructure) add(folder string) {
	nextTree, ok := structure[folder]
	if !ok {
		nextTree = FolderStructure{}
		structure[folder] = nextTree
	}
}

func (structure FolderStructure) print(depth int) {
	keys := make([]string, 0, len(structure))
	for k := range structure {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(strings.Repeat("  ", depth), key)
		structure[key].print(depth + 1)
	}
}

func parse(input []string) FolderStructure {
	tree := FolderStructure{"/": FolderStructure{}}
	parseRecursive(input, tree)
	tree.print(0)
	return tree
}

func parseRecursive(input []string, root FolderStructure) (handled int, size int) {
	handled = 0
	size = 0
	for i := 0; i < len(input); i++ {
		command := input[i]
		handled++

		if !strings.HasPrefix(command, "$") {
			splitted := strings.Split(command, " ")
			if strings.HasPrefix(command, "dir ") {
				root.add(splitted[1])
			} else {
				root.add(splitted[1] + " (file, size=" + splitted[0] + ")")
				size += functions.ParseInt(splitted[0])
			}
		} else if strings.HasPrefix(command, "$ cd ") {
			dir := command[5:]
			if dir == ".." {
				break
			} else {
				childrenAdded, childrenSize := parseRecursive(input[i+1:], root[dir])
				root[dir+" (dir, size="+strconv.Itoa(childrenSize)+")"] = root[dir]
				delete(root, dir)
				size += childrenSize
				handled += childrenAdded
				i += childrenAdded
			}
		}
	}
	return handled, size
}
