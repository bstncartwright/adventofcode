package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := parse(input)

	part1 := func() int {
		var (
			x         = 1
			stepToSum = 20
			sum       int
			step      int
		)
		for _, i := range list {
			for j := 0; j < i.steps; j++ {
				step++

				if step == stepToSum {
					sum += (x * step)
					stepToSum += 40
				}
			}

			x += i.add
		}
		return sum
	}()

	part2 := func() int {
		// initialize output
		// 40 wide
		// 6 high
		crt := make([][]bool, 6)
		for r := range crt {
			crt[r] = make([]bool, 40)
		}

		// correcty populate
		var (
			reg  = 1
			step int
		)
		for _, i := range list {
			for j := 0; j < i.steps; j++ {

				// get crt position from step
				y := step / 40
				x := step - (40 * y)

				if reg == x || (reg-1) == x || (reg+1) == x {
					crt[y][x] = true
				}

				step++

			}

			reg += i.add
		}

		// print output
		var sb strings.Builder
		for _, row := range crt {
			for _, col := range row {
				if col {
					sb.WriteString("#")
				} else {
					sb.WriteString(".")
				}
			}
			sb.WriteString("\n")
		}
		fmt.Println(sb.String())

		return 0
	}()

	return part1, part2
}

type instruction struct {
	add   int
	steps int
}

func parse(s string) []instruction {
	lines := strings.Split(s, "\n")
	i := make([]instruction, 0, len(lines))
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "addx"):
			l := strings.Split(line, " ")

			n, err := strconv.Atoi(l[1])
			if err != nil {
				panic(err)
			}

			i = append(i, instruction{add: n, steps: 2})
		case line == "noop":
			i = append(i, instruction{add: 0, steps: 1})
		}
	}
	return i
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
