package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := parse(input)

	var sum int
	for _, p := range list {
		// for every p, determine if one or twos assignment is entirely within the others
		switch {
		case p.one.start >= p.two.start && p.one.end <= p.two.end:
			sum++
			continue
		case p.two.start >= p.one.start && p.two.end <= p.one.end:
			sum++
			continue
		}
	}

	// now to determine if any overlap at all
	var sum2 int
	for _, p := range list {
		switch {
		case p.one.start >= p.two.start && p.one.start <= p.two.end:
			sum2++
			continue
		case p.two.start >= p.one.start && p.two.start <= p.one.end:
			sum2++
			continue
		case p.one.end >= p.two.start && p.one.end <= p.two.end:
			sum2++
			continue
		}
	}

	part1, part2 := sum, sum2
	return part1, part2
}

type pair struct {
	one assignment
	two assignment
}

type assignment struct {
	start int
	end   int
}

func parse(s string) []pair {
	lines := strings.Split(s, "\n")

	var pairs []pair
	for _, line := range lines {
		var p pair
		a := strings.Split(line, ",")

		a1 := strings.Split(a[0], "-")
		a2 := strings.Split(a[1], "-")

		p.one = assignment{
			start: mustParseInt(a1[0]),
			end:   mustParseInt(a1[1]),
		}

		p.two = assignment{
			start: mustParseInt(a2[0]),
			end:   mustParseInt(a2[1]),
		}

		pairs = append(pairs, p)
	}

	return pairs
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
