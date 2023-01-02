package main

import (
	"advent-of-code-2022/functions"
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Index     int
	Items     []int64
	Operation func(old int64) int64
	Divisible int64
	IfTrue    int
	IfFalse   int
	Inspected int
}

func parse(input []string) []Monkey {
	var monkeys []Monkey
	for i := 0; i < len(input); i += 7 {
		monkeys = append(monkeys, Monkey{
			Index:     lineToInt(input[i], "Monkey ", 1),
			Items:     parseItems(input[i+1]),
			Operation: parseOperation(input[i+2]),
			Divisible: int64(lineToInt(input[i+3], "  Test: divisible by ", 0)),
			IfTrue:    lineToInt(input[i+4], "    If true: throw to monkey ", 0),
			IfFalse:   lineToInt(input[i+5], "    If false: throw to monkey ", 0),
		})
	}
	return monkeys
}

func parseItems(line string) []int64 {
	var items []int64
	stringItems := line[len("  Starting items: "):]
	splitted := strings.Split(stringItems, ", ")
	for _, item := range splitted {
		items = append(items, int64(functions.ParseInt(item)))
	}
	return items
}

func parseOperation(line string) func(old int64) int64 {
	stringOperation := line[len("  Operation: new = "):]
	splitted := strings.Split(stringOperation, " ")
	operator := splitted[1]

	return func(old int64) int64 {
		value1 := resolveValue(splitted[0], old)
		value2 := resolveValue(splitted[2], old)
		if operator == "+" {
			return value1 + value2
		} else if operator == "*" {
			return value1 * value2
		} else {
			fmt.Println("unknown operator: ", operator)
			return 1
		}
	}
}

func resolveValue(raw string, old int64) int64 {
	if raw == "old" {
		return old
	} else {
		n, _ := strconv.ParseInt(raw, 10, 64)
		return n
	}
}

func lineToInt(line string, prefix string, offset int) int {
	return functions.ParseInt(line[len(prefix) : len(line)-offset])
}
