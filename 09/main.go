package main

import (
	"advent-of-code-2022/functions"
	"fmt"
)

func main() {
	input := functions.Get(9)

	motions := parse(input)
	start := Position{0, 0}
	fmt.Println("[Day 09, part one] Tail positions:", applyMotions(motions, start, []Position{start}))
	fmt.Println("[Day 09, part two] Tail positions:", applyMotions(motions, start, []Position{start}))
}

func applyMotions(motions []Motion, head Position, tail []Position) int {
	for _, motion := range motions {
		fmt.Println("==", motion, "==")
		head, tail = applyMotion(motion, head, tail)
	}
	printTail(tail)
	return countUnique(tail)
}

func applyMotion(motion Motion, head Position, tail []Position) (Position, []Position) {
	for i := 0; i < motion.Length; i++ {
		switch motion.Direction {
		case "U":
			head.X--
		case "D":
			head.X++
		case "L":
			head.Y--
		case "R":
			head.Y++
		}
		tail = moveTailIfNeeded(head, tail)
	}

	return head, tail
}
