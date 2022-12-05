package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	arr := functions.Get(4)
	count := 0
	for _, assignments := range arr {
		splitted := strings.Split(assignments, ",")
		split1 := strings.Split(splitted[0], "-")
		split2 := strings.Split(splitted[1], "-")
		idx1a, _ := strconv.ParseFloat(split1[0], 8)
		idx1b, _ := strconv.ParseFloat(split1[1], 8)
		idx2a, _ := strconv.ParseFloat(split2[0], 8)
		idx2b, _ := strconv.ParseFloat(split2[1], 8)

		start := math.Max(idx1a, idx2a)
		end := math.Min(idx1b, idx2b)
		diff := end - start

		if diff >= 0 {
			count++
		}

		fmt.Printf("%-12s %2g - %2g = %g\n", assignments, start, end, diff)
	}

	fmt.Println("Total overlapping", count)
}
