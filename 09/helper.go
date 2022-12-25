package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

type Motion struct {
	Direction string
	Length    int
}

type Position struct {
	X int
	Y int
}

func parse(input []string) []Motion {
	var motions []Motion
	for _, line := range input {
		splitted := strings.Split(line, " ")
		motions = append(motions, Motion{splitted[0], functions.ParseInt(splitted[1])})
	}
	return motions
}

func moveTailIfNeeded(head Position, tails []Position) []Position {
	tail := tails[len(tails)-1]
	x := head.X - tail.X
	y := head.Y - tail.Y
	if x < -1 || x > 1 || y < -1 || y > 1 {
		newTail := determineNewTail(tail, x, y)
		fmt.Println("Tail needs to move for  :", head, "|", tail, "=>", newTail)
		//if !slices.Contains(tails, newTail) {
		tails = append(tails, newTail)
		//}
	} else {
		fmt.Println("Tail doesnt have to move:", head, "|", tail)
	}
	return tails
}

func determineNewTail(tail Position, x int, y int) Position {
	if x == 0 {
		return Position{X: tail.X, Y: tail.Y + y/2}
	} else if y == 0 {
		return Position{X: tail.X + x/2, Y: tail.Y}
	} else {
		if x <= -2 || x >= 2 {
			return Position{X: tail.X + x/2, Y: tail.Y + y}
		} else if y <= -2 || y >= 2 {
			return Position{X: tail.X + x, Y: tail.Y + y/2}
		}
	}
	fmt.Println("????")
	return tail
}

func printTail(tails []Position) {
	startX := 0
	startY := 0
	endX := 0
	endY := 0
	for _, tail := range tails {
		if tail.X < startX {
			startX = tail.X
		} else if tail.X > endX {
			endX = tail.X
		}

		if tail.Y < startY {
			startY = tail.Y
		} else if tail.Y > endY {
			endY = tail.Y
		}
	}

	for x := startX - 1; x <= endX+1; x++ {
		for y := startY - 1; y <= endY+1; y++ {
			if (slices.Contains(tails, Position{X: x, Y: y})) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func countUnique(positions []Position) int {
	var unique []Position
outer:
	for _, v := range positions {
		for i, u := range unique {
			if v.X == u.X && v.Y == u.Y {
				unique[i] = v
				continue outer
			}
		}
		unique = append(unique, v)
	}
	return len(unique)
}
