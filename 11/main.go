package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("[Day 11, part two] Total sum: %d\n", mostInspected(20, false))
	fmt.Printf("[Day 11, part two] Total sum: %d\n", mostInspected(10000, true))
}

func mostInspected(rounds int, useModulo bool) int {
	monkeys := parse(functions.Get(11))
	execute(monkeys, rounds, useModulo)

	printMonkeys(monkeys)

	sort.Slice(monkeys, func(i int, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	return monkeys[0].Inspected * monkeys[1].Inspected
}

func execute(monkeys []Monkey, rounds int, useModulo bool) {
	// Thanks reddit
	bigMod := int64(1)
	for _, m := range monkeys {
		bigMod *= m.Divisible
	}

	for i := 0; i < rounds; i++ {
		for idx, monkey := range monkeys {
			for _, item := range monkey.Items {
				worrying := monkey.Operation(item)
				if useModulo {
					worrying %= bigMod
				} else {
					worrying = worrying / 3
				}

				if worrying%monkey.Divisible == 0 {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, worrying)
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, worrying)
				}
				monkeys[idx].Inspected++
			}
			monkeys[idx].Items = nil
		}
	}
}

func printMonkeys(monkeys []Monkey) {
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d: [%6d inspections] %s\n", i, monkey.Inspected,
			strings.Trim(strings.Replace(fmt.Sprint(monkey.Items), " ", ",", -1), "[]"))
	}
}
