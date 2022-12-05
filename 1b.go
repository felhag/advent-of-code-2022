package main

import (
	"fmt"
	"sort"
	"strconv"
)
import "go-test/functions"

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

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println("Max:", elves[0]+elves[1]+elves[2])
}
