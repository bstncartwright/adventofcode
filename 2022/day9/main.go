package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	instructions := parse(input)

	part1 := func() int {
		hx, hy := 0, 0
		tx, ty := 0, 0

		var placesVisited []coordinate

		for _, i := range instructions {
			for j := 0; j < i.steps; j++ {
				switch i.direction {
				case "U":
					hy++
				case "D":
					hy--
				case "L":
					hx--
				case "R":
					hx++
				default:
					panic(i.direction)
				}

				// are they still touching?
				if Abs(hx-tx) > 1 || Abs(hy-ty) > 1 {
					// check if diagnol, for some reason my brain can't do this with math tonight
					switch {
					case hx != tx && hy != ty:
						tx += gsd(hx, tx)
						ty += gsd(hy, ty)
					case Abs(hx-tx) > 1:
						tx += gsd(hx, tx)
					case Abs(hy-ty) > 1:
						ty += gsd(hy, ty)
					}
				}

				if !coordinateContains(placesVisited, tx, ty) {
					placesVisited = append(placesVisited, coordinate{x: tx, y: ty})
				}

			}
		}
		return len(placesVisited)
	}()

	part2 := func() int {
		hx, hy := 0, 0

		sections := make([]coordinate, 9)
		var placesVisited []coordinate

		for _, i := range instructions {
			for j := 0; j < i.steps; j++ {
				switch i.direction {
				case "U":
					hy++
				case "D":
					hy--
				case "L":
					hx--
				case "R":
					hx++
				default:
					panic(i.direction)
				}

				for k, c := range sections {
					var x, y int
					if k == 0 {
						x, y = checkMove(hx, hy, c.x, c.y)
					} else {
						x, y = checkMove(sections[k-1].x, sections[k-1].y, c.x, c.y)
					}

					sections[k].x += x
					sections[k].y += y
				}

				if !coordinateContains(placesVisited, sections[len(sections)-1].x, sections[len(sections)-1].y) {
					placesVisited = append(placesVisited, coordinate{x: sections[len(sections)-1].x, y: sections[len(sections)-1].y})
				}

			}
		}

		return len(placesVisited)
	}()

	return part1, part2
}

func checkMove(hx, hy, tx, ty int) (int, int) {
	if Abs(hx-tx) > 1 || Abs(hy-ty) > 1 {
		// check if diagnol, for some reason my brain can't do this with math tonight
		switch {
		case hx != tx && hy != ty:
			return gsd(hx, tx), gsd(hy, ty)
		case Abs(hx-tx) > 1:
			return gsd(hx, tx), 0
		case Abs(hy-ty) > 1:
			return 0, gsd(hy, ty)
		}
	}
	return 0, 0
}

func gsd(x, y int) int {
	r := x - y
	if r < 0 {
		return -1
	}
	return 1
}

func coordinateContains(c []coordinate, x, y int) bool {
	for _, c2 := range c {
		if c2.x == x && c2.y == y {
			return true
		}
	}
	return false
}

type coordinate struct {
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type instruction struct {
	direction string
	steps     int
}

func parse(s string) []instruction {
	lines := strings.Split(s, "\n")

	instructions := make([]instruction, 0, len(lines))
	for _, line := range lines {
		b := strings.Split(line, " ")

		i, err := strconv.Atoi(b[1])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, instruction{
			direction: b[0],
			steps:     i,
		})
	}

	return instructions
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
