package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func parse() ([]string, [][]rune) {
	arr := functions.Get(5)
	idx := slices.IndexFunc(arr, func(c string) bool {
		_, err := strconv.Atoi(string(c[1]))
		return err == nil
	})
	stacks := parseStacks(arr[0:idx])
	actions := arr[idx+2:]
	printStacks(stacks)
	return actions, stacks
}

func parseStacks(input []string) [][]rune {
	stacks := (len(input[len(input)-1]) + 1) / 4
	var result [][]rune
	for i := 0; i < stacks; i++ {
		result = append(result, make([]rune, 0))
	}

	for rowIdx := len(input); rowIdx > 0; rowIdx-- {
		row := input[rowIdx-1]
		for i := 0; i <= len(row)-1; i = i + 4 {
			if row[i] == '[' {
				result[i/4] = append(result[i/4], rune(row[i+1]))
			}
		}
	}
	return result
}

func printStacks(stacks [][]rune) {
	max := 0
	for _, stack := range stacks {
		if len(stack) > max {
			max = len(stack)
		}
	}

	for height := max - 1; height >= 0; height-- {
		for _, stack := range stacks {
			if len(stack) > height {
				fmt.Printf("[%s] ", string(stack[height]))
			} else {
				fmt.Print(strings.Repeat(" ", 4))
			}
		}
		fmt.Print("\n")
	}
	fmt.Println(strings.Repeat("=", len(stacks)*4-1))
}

func parseAction(action string) (int, int, int) {
	splitted := strings.Split(action, " ")
	count := functions.ParseInt(splitted[1])
	fromIdx := functions.ParseInt(splitted[3]) - 1
	toIdx := functions.ParseInt(splitted[5]) - 1
	return count, fromIdx, toIdx
}

func getTopCrates(stacks [][]rune) string {
	tops := ""
	for _, stack := range stacks {
		r := stack[len(stack)-1]
		tops = tops + string(r)
	}
	return tops
}
