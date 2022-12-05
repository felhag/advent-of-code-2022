package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"strconv"
)

func main() {
	arr1 := functions.Get(1)
	elves := []int{0}

	for _, element := range arr1 {
		if element == "" {
			elves = append(elves, 0)
		} else {
			atoi, _ := strconv.Atoi(element)
			elves[len(elves)-1] += atoi
		}
	}

	max := 0

	for idx, elf := range elves {
		if elves[max] < elf {
			max = idx
		}
	}

	fmt.Println("Elf", max+1, "has the most calories:", elves[max])
}
