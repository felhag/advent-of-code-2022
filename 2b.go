package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"strings"
)

func main() {
	arr1 := functions.Get(2)
	total := 0

	for _, row := range arr1 {
		split := strings.Split(row, " ")

		opponent := []rune(split[0])[0] - 64
		strategy := split[1]
		me := findMyShape(strategy, opponent)
		result := opponent - me
		var score = int(me)

		if result == 0 {
			fmt.Println(split, "Score", score, " + 3 = ", total+score+3)
			score += 3
		} else if result == -1 || result == 2 {
			fmt.Println(split, "Score", score, " + 6 = ", total+score+6)
			score += 6
		} else {
			fmt.Println(split, "Score", score, " + 0 = ", total+score)
		}
		total += score
	}
	fmt.Println("Final score: ", total)
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
