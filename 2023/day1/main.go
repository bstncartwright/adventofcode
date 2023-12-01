package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var nums = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := parse(input)

	part1, part2 := 0, 0

	for _, v := range list {
		// v := "qbptwoneqcdkqqtnmfjrpplseven13one"
		numIndexes := make(map[int]string)
		for _, s := range nums {
			start := 0
			for {
				i := strings.Index(v[start:], s)
				if i == -1 {
					break
				}

				numIndexes[start+i] = s
				start += i + len(s)
			}
		}

		// go through map to find lowest index
		var (
			lowest  = len(v) + 100
			highest = -1
		)
		for k := range numIndexes {
			if k < lowest {
				lowest = k
			}

			if k > highest {
				highest = k
			}
		}

		var d1, d2 int

		switch numIndexes[lowest] {
		case "zero":
			d1 = 0
		case "one":
			d1 = 1
		case "two":
			d1 = 2
		case "three":
			d1 = 3
		case "four":
			d1 = 4
		case "five":
			d1 = 5
		case "six":
			d1 = 6
		case "seven":
			d1 = 7
		case "eight":
			d1 = 8
		case "nine":
			d1 = 9
		default:
			d1 = pkg.MustAtoi(numIndexes[lowest])
		}

		switch numIndexes[highest] {
		case "zero":
			d2 = 0
		case "one":
			d2 = 1
		case "two":
			d2 = 2
		case "three":
			d2 = 3
		case "four":
			d2 = 4
		case "five":
			d2 = 5
		case "six":
			d2 = 6
		case "seven":
			d2 = 7
		case "eight":
			d2 = 8
		case "nine":
			d2 = 9
		default:
			d2 = pkg.MustAtoi(numIndexes[highest])
		}

		part1 += pkg.MustAtoi(fmt.Sprintf("%d%d", d1, d2))

	}

	return part1, part2
}

func parse(s string) []string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		pkg.MustScanf(line, "")
	}
	return lines
}

func main() {
	execute.Run(run, nil, puzzle2, true)
}
