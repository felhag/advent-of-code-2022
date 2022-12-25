package main

import (
	"advent-of-code-2022/functions"
	"strings"
)

func tick(pending int, x int, todos []string) (int, int, []string) {
	if pending != 0 {
		x += pending
		pending = 0
	} else {
		cmd := todos[0]
		todos = todos[1:]
		if strings.HasPrefix(cmd, "addx") {
			pending = functions.ParseInt(cmd[5:])
		}
	}
	return pending, x, todos
}
