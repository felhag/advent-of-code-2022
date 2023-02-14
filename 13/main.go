package main

import (
	"advent-of-code-2022/functions"
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func main() {
	pairs := parse(functions.Get(13))
	fmt.Printf("[Day 13, part one]: %d\n", part1(pairs))
	fmt.Printf("[Day 13, part two]: %d\n", part2(pairs))
}

func part1(packets [][]any) int {
	result := 0
	for i, pair := range packets {
		if compare(pair[0], pair[1]) <= 0 {
			result += i + 1
		}
	}

	return result
}

func part2(packets [][]any) int {
	new := []any{}
	for _, pair := range packets {
		new = append(new, pair...)
	}

	var divider1 any
	err := json.Unmarshal([]byte("[[2]]"), &divider1)
	if err != nil {
		log.Fatal(err)
	}

	var divider2 any
	err = json.Unmarshal([]byte("[[6]]"), &divider2)
	if err != nil {
		log.Fatal(err)
	}

	new = append(new, []any{divider1, divider2}...)
	sort.Slice(new, func(i, j int) bool {
		return compare(new[i], new[j]) <= 0
	})

	result := 1
	for i, packet := range new {
		packet, err := json.Marshal(packet)
		if err != nil {
			log.Fatal(err)
		}
		packetString := string(packet)
		if packetString == "[[2]]" || packetString == "[[6]]" {
			result *= i + 1
		}
	}

	return result
}
