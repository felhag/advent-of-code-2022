package main

import (
	"advent-of-code-2022/functions"
)

func main() {
	arr1 := functions.Get(2)
	calculateScore(arr1, "one", func(strategy string, opponent rune) rune {
		return []rune(strategy)[0] - 87
	})
	calculateScore(arr1, "two", func(strategy string, opponent rune) rune {
		return findMyShape(strategy, opponent)
	})
}
