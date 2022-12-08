package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"golang.org/x/exp/slices"
	"unicode"
)

func main() {
	arr := functions.Get(3)
	fmt.Println("[Day 03, part one]:", partOne(arr))
	fmt.Println("[Day 03, part two]:", partTwo(arr))
}

func partOne(arr []string) int {
	sum := 0

	for _, rucksack := range arr {
		runes := []rune(rucksack)
		split := len(rucksack) / 2
		first := runes[0:split]
		second := runes[split:]
		var handled []rune
		items := 0

		for _, item := range first {
			if slices.Contains(second, item) && !slices.Contains(handled, item) {
				if unicode.IsUpper(item) {
					items += int(item) - 38
				} else {
					items += int(item) - 96
				}
				handled = append(handled, item)
			}
		}

		sum += items
	}

	return sum
}

func partTwo(arr []string) int {
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

	return badges
}

func findBadge(rucksack1 []rune, rucksack2 []rune, rucksack3 []rune) rune {
	for _, item := range rucksack1 {
		if slices.Contains(rucksack2, item) && slices.Contains(rucksack3, item) {
			return item
		}
	}
	return 0
}
