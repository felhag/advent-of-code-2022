package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Get(Day int) []string {
	var file = "0" + strconv.Itoa(Day) + "/input.txt"
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func FindMax(input []int) int {
	max := 0
	for idx, elf := range input {
		if input[max] < elf {
			max = idx
		}
	}
	return max
}
