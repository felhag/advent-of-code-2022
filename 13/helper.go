package main

import (
	"encoding/json"
)

func parse(rows []string) [][]any {
	pairs := [][]any{}

	for i := 0; i <= len(rows)-1; i = i + 3 {
		var a, b any
		json.Unmarshal([]byte(rows[i]), &a)
		json.Unmarshal([]byte(rows[i+1]), &b)
		pair := []any{a, b}
		pairs = append(pairs, pair)
	}

	return pairs
}

func compare(p1 any, p2 any) int {
	_, ok1 := p1.(float64)
	_, ok2 := p2.(float64)
	if ok1 && ok2 {
		return int(p1.(float64)) - int(p2.(float64))
	}

	if ok1 {
		p1 = []any{p1}
	}
	if ok2 {
		p2 = []any{p2}
	}

	if len(p1.([]any)) == 0 || len(p2.([]any)) == 0 {
		return len(p1.([]any)) - len(p2.([]any))
	}

	result := compare(p1.([]any)[0], p2.([]any)[0])
	if result == 0 {
		next1 := p1.([]any)[1:]
		next2 := p2.([]any)[1:]

		if len(next1) == 0 || len(next2) == 0 {
			return len(next1) - len(next2)
		}
		return compare(next1, next2)
	}

	return result
}
