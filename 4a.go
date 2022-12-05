package main

import (
	"advent-of-code-2022/functions"
	"fmt"
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
		idx1a, _ := strconv.Atoi(split1[0])
		idx1b, _ := strconv.Atoi(split1[1])
		idx2a, _ := strconv.Atoi(split2[0])
		idx2b, _ := strconv.Atoi(split2[1])

		if (idx1a >= idx2a && idx1b <= idx2b) ||
			(idx1a <= idx2a && idx1b >= idx2b) {
			count++
		}
	}

	fmt.Println("Total overlapping", count)
}
