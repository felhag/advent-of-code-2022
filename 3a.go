package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"unicode"
)
import "golang.org/x/exp/slices"

func main() {
	arr := functions.Get(3)
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

	fmt.Println("Total sum", sum)
}
