package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Get(Day int) []string {
	var file = "input/" + strconv.Itoa(Day) + ".txt"
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
