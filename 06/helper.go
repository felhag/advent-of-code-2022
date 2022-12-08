package main

import "fmt"

func findFirstUnique(input string, length int) {
	for i := length; i < len(input); i++ {
		chars := input[i-length : i]
		if onlyUnique(chars) {
			fmt.Printf("Unique index found: %d, [%s]\n", i, chars)
			break
		}
	}
}

func onlyUnique(chars string) bool {
	for outerIdx, outerChar := range chars {
		for idx, char := range chars {
			if idx != outerIdx && outerChar == char {
				return false
			}
		}
	}
	return true
}
