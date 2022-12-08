package main

import (
	"advent-of-code-2022/functions"
)

func main() {
	input := functions.Get(6)[0]
	findFirstUnique(input, 4)
	findFirstUnique(input, 14)
}
