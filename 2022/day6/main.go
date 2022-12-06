package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	var part1, part2 int

	for i := range input {
		if i-4 < 0 {
			continue
		}

		if noDuplicates(input[i-4 : i]) {
			part1 = i
			break
		}
	}

	for i := range input {
		if i-14 < 0 {
			continue
		}

		if noDuplicates(input[i-14 : i]) {
			part2 = i
			break
		}
	}

	return part1, part2
}

func noDuplicates(s string) bool {
	for i := range s {
		for j := range s {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
