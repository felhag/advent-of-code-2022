package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"sort"
)

func main() {
	elves := findElves(functions.Get(1))
	max := functions.FindMax(elves)
	fmt.Println("[Day 01, part one]: Elf", max+1, "has the most calories:", elves[max])

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println("[Day 01, part two]: Max:", elves[0]+elves[1]+elves[2])
}
