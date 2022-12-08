package main

import (
	"strconv"
)

func findElves(input []string) []int {
	elves := []int{0}
	for _, element := range input {
		if element == "" {
			elves = append(elves, 0)
		} else {
			atoi, _ := strconv.Atoi(element)
			elves[len(elves)-1] += atoi
		}
	}
	return elves
}
