package main

import (
	"strconv"
	"strings"
)

func parseAssignment(assignments string) (float64, float64, float64, float64) {
	splitted := strings.Split(assignments, ",")
	split1 := strings.Split(splitted[0], "-")
	split2 := strings.Split(splitted[1], "-")
	idx1a, _ := strconv.ParseFloat(split1[0], 8)
	idx1b, _ := strconv.ParseFloat(split1[1], 8)
	idx2a, _ := strconv.ParseFloat(split2[0], 8)
	idx2b, _ := strconv.ParseFloat(split2[1], 8)
	return idx1a, idx1b, idx2a, idx2b
}
