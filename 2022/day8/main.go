package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	m := parse(input)

	// these solutions are pretty brute forcey... did this late haha

	var sum int
	for y := range m {
		for x := range m[y] {
			// determine if this tree is visible

			visible := 4

			// first look left (i--)
			for k := x - 1; k >= 0; k-- {
				if m[y][k] >= m[y][x] {
					visible--
					break
				}
			}

			// then look right (i++)
			for k := x + 1; k < len(m[y]); k++ {
				if m[y][k] >= m[y][x] {
					visible--
					break
				}
			}

			// now look down
			for k := y + 1; k < len(m); k++ {
				if m[k][x] >= m[y][x] {

					visible--
					break
				}
			}

			// now look up
			for k := y - 1; k >= 0; k-- {
				if m[k][x] >= m[y][x] {

					visible--
					break
				}
			}

			if visible > 0 {
				sum++
			}
		}
	}

	var sum2 int
	for y := range m {
		for x := range m[y] {

			var score int

			// first look left (i--)
			var left int = x
			for k := x - 1; k >= 0; k-- {
				if m[y][k] >= m[y][x] {
					left = x - k
					break
				}
			}

			// then look right (i++)
			var right int = len(m[y]) - 1 - x
			for k := x + 1; k < len(m[y]); k++ {
				if m[y][k] >= m[y][x] {
					right = k - x
					break
				}
			}

			// now look down
			var down int = len(m) - 1 - y
			for k := y + 1; k < len(m); k++ {
				if m[k][x] >= m[y][x] {
					down = k - y
					break
				}
			}

			// now look up
			var up int = y
			for k := y - 1; k >= 0; k-- {
				if m[k][x] >= m[y][x] {
					up = y - k
					break
				}
			}

			score = left * right * up * down

			if score > sum2 {
				sum2 = score
			}

		}
	}

	part1, part2 := sum, sum2

	return part1, part2
}

func parse(s string) [][]int {
	lines := strings.Split(s, "\n")

	var m [][]int
	for _, line := range lines {
		var row []int
		for _, c := range strings.Split(line, "") {
			i, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}

			row = append(row, i)
		}

		m = append(m, row)
	}

	return m
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
