package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := parse(input)

	var totals []int
	for _, elf := range list {
		var sum int
		for _, count := range elf {
			sum += count
		}

		totals = append(totals, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	part1, part2 := totals[0], totals[0]+totals[1]+totals[2]

	return part1, part2
}

func parse(s string) [][]int {
	lines := strings.Split(s, "\n")
	// split lines into groups of ints separated by an empty line
	var (
		groups = make([][]int, 0)
		group  = make([]int, 0)
	)

	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = make([]int, 0)
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		group = append(group, i)
	}

	return groups
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
