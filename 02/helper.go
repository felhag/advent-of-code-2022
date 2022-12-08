package main

import (
	"fmt"
	"strings"
)

func getScore(opponent rune, me rune) int {
	result := opponent - me
	var score = int(me)

	if result == 0 {
		score += 3
	} else if result == -1 || result == 2 {
		score += 6
	} else {
	}
	return score
}

func calculateScore(arr1 []string, part string, f func(strategy string, opponent rune) rune) {
	total := 0

	for _, row := range arr1 {
		split := strings.Split(row, " ")

		opponent := []rune(split[0])[0] - 64
		me := f(split[1], opponent)
		result := opponent - me
		var score = int(me)

		if result == 0 {
			fmt.Println(split, "Score", score, " + 3 = ", total+score+3)
			score += 3
		} else if result == -1 || result == 2 {
			fmt.Println(split, "Score", score, " + 6 = ", total+score+6)
			score += 6
		}
		total += getScore(opponent, me)
	}
	fmt.Printf("[Day 02, part %s] Final score: %d\n", part, total)
}

func findMyShape(strategy string, opponent rune) rune {
	var me rune
	if strategy == "Y" {
		me = opponent
	} else if strategy == "X" {
		if opponent == 1 {
			me = 3
		} else {
			me = opponent - 1
		}
	} else {
		if opponent == 3 {
			me = 1
		} else {
			me = opponent + 1
		}
	}
	return me
}
