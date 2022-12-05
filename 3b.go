package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"unicode"
)
import "golang.org/x/exp/slices"

func main() {
	arr := functions.Get(3)
	badges := 0

	for i := 0; i <= len(arr)-1; i = i + 3 {
		rucksack1 := []rune(arr[i])
		rucksack2 := []rune(arr[i+1])
		rucksack3 := []rune(arr[i+2])

		badge := findBadge(rucksack1, rucksack2, rucksack3)
		if unicode.IsUpper(badge) {
			badges += int(badge) - 38
		} else {
			badges += int(badge) - 96
		}
	}

	fmt.Println("Total sum", badges)
}

func findBadge(rucksack1 []rune, rucksack2 []rune, rucksack3 []rune) rune {
	for _, item := range rucksack1 {
		if slices.Contains(rucksack2, item) && slices.Contains(rucksack3, item) {
			return item
		}
	}
	return 0
}
