package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (any, any) {
	part1, part2 := 0, 0

	for _, line := range strings.Split(input, "\n") {
		var i int
		s := strings.Split(line, ":")
		pkg.MustScanf(s[0], "Game %d", &i)

		var blue, green, red int
		for _, c := range strings.Split(s[1], ";") {
			// for a single pull
			p := strings.Split(c, ", ")
			for _, v := range p {
				var n int
				var color string
				pkg.MustScanf(v, "%d %s", &n, &color)
				switch color {
				case "blue":
					if blue < n {
						blue = n
					}
				case "green":
					if green < n {
						green = n
					}
				case "red":
					if red < n {
						red = n
					}
				}
			}
		}
		// 12 red, 13 green, 14 blue

		// fmt.Printf("Game %d: %d red, %d green, %d blue\n", i, red, green, blue)
		if red <= 12 && green <= 13 && blue <= 14 {
			// fmt.Printf("its good\n")
			part1 += i
		}
	}

	for _, line := range strings.Split(input, "\n") {
		var i int
		s := strings.Split(line, ":")
		pkg.MustScanf(s[0], "Game %d", &i)

		var blue, green, red int
		for _, c := range strings.Split(s[1], ";") {
			// for a single pull
			p := strings.Split(c, ", ")
			for _, v := range p {
				var n int
				var color string
				pkg.MustScanf(v, "%d %s", &n, &color)
				switch color {
				case "blue":
					if blue < n {
						blue = n
					}
				case "green":
					if green < n {
						green = n
					}
				case "red":
					if red < n {
						red = n
					}
				}
			}
		}
		pow := red * green * blue
		fmt.Printf("Game %d: %d red, %d green, %d blue, %d\n", i, red, green, blue, pow)
		part2 += pow
	}

	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
