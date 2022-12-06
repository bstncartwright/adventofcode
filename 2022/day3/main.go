package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

func run(input string) (interface{}, interface{}) {
	list := parse(input)

	var sum int
	for _, b := range list {
		for _, v := range b.both {
			i := runeToIntScore(v)
			sum += i
		}
	}

	var sum2 int
	split := splitIntoGroupsOfThree(list)
	for _, b := range split {
		var common []rune
		for _, r := range b[0].all {
			if strings.ContainsRune(b[1].all, r) && strings.ContainsRune(b[2].all, r) {
				common = append(common, r)
			}
		}

		common = dedupeRunes(common)
		for _, r := range common {
			sum2 += runeToIntScore(r)
		}
	}

	part1, part2 := sum, sum2

	return part1, part2
}

func splitIntoGroupsOfThree(bags []bag) [][]bag {
	// split into groups of 3
	v := make([][]bag, 0, len(bags)/3)

	for i := 0; i < len(bags); i += 3 {
		b := bags[i : i+3]
		v = append(v, b)
	}

	return v
}

// a - 97
// b - 98
// A - 65
// B - 66

func runeToIntScore(r rune) int {
	switch {
	case 'a' <= r && r <= 'z':
		return int(r) - 96
	case 'A' <= r && r <= 'Z':
		return int(r) - 38
	default:
		panic("invalid rune")
	}
}

type bag struct {
	all   string
	left  string
	right string
	both  []rune
}

func parse(s string) []bag {
	lines := strings.Split(s, "\n")

	bags := make([]bag, 0, len(lines))

	for _, line := range lines {
		var b bag

		b.all = line
		b.left = line[:len(line)/2]
		b.right = line[len(line)/2:]

		for _, l := range b.left {
			for _, r := range b.right {
				if l == r {
					b.both = append(b.both, l)
					break
				}
			}
		}

		b.both = dedupeRunes(b.both)

		bags = append(bags, b)
	}

	return bags
}

// remove duplicates from a slice of runes
func dedupeRunes(r []rune) []rune {
	var out []rune
	m := make(map[rune]bool)

	for _, v := range r {
		if !m[v] {
			out = append(out, v)
			m[v] = true
		}
	}

	return out
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
