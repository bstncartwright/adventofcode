package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	m := parse(input)

	for _, instruction := range m.instructions {
		for i := 0; i < instruction.amount; i++ {
			from := instruction.from - 1
			to := instruction.to - 1

			// move the top item from instruction.from to instruction.to
			m.stacks[to].items = append(m.stacks[to].items, m.stacks[from].items[len(m.stacks[from].items)-1])
			// remove the top item from instruction.from
			m.stacks[from].items = m.stacks[from].items[:len(m.stacks[from].items)-1]
		}
	}

	// get top rune of all stacks
	var p string
	for _, s := range m.stacks {
		p += string(s.items[len(s.items)-1])
	}

	m2 := parse(input)
	for _, instruction := range m2.instructions {
		from := instruction.from - 1
		to := instruction.to - 1

		m2.stacks[to].items = append(
			m2.stacks[to].items,
			m2.stacks[from].items[len(m2.stacks[from].items)-(instruction.amount):]...,
		)

		m2.stacks[from].items = m2.stacks[from].items[:len(m2.stacks[from].items)-(instruction.amount)]
	}

	var p2 string
	for _, s := range m2.stacks {
		p2 += string(s.items[len(s.items)-1])
	}

	part1, part2 := p, p2

	return part1, part2
}

type Map struct {
	stacks       []Stack
	instructions []Instruction
}

type Instruction struct {
	amount, from, to int
}

func (i Instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", i.amount, i.from, i.to)
}

type Stack struct {
	items []rune
}

func (s Stack) String() string {
	return fmt.Sprintf("Stack(%d): [%s]", len(s.items), string(s.items))
}

func parse(s string) Map {
	lines := strings.Split(s, "\n")

	// first find where the stacks stop and the instructions start
	var (
		m                Map
		stackLines       []string
		instructionLines []string
	)
	for i, line := range lines {
		if line == "" {
			stackLines = lines[:i]
			instructionLines = lines[i+1:]
			break
		}
	}

	// the last line of the stack lines will have the number of stacks
	var (
		numStacks = len(strings.ReplaceAll(stackLines[len(stackLines)-1], " ", ""))
		stacks    = make([]Stack, numStacks)
	)
	for i := len(stackLines) - 2; i >= 0; i-- {
		line := stackLines[i]

		for j := 1; j < (numStacks*4)+2; j += 4 {
			if len(line) <= j {
				break
			}
			if line[j] != ' ' {
				stacks[(j-1)/4].items = append(stacks[(j-1)/4].items, rune(line[j]))
			}
		}

	}

	var (
		r            = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		instructions = make([]Instruction, 0, len(instructionLines))
	)
	for _, line := range instructionLines {
		s := r.FindStringSubmatch(line)

		instructions = append(instructions, Instruction{
			amount: mustParseInt(s[1]),
			to:     mustParseInt(s[3]),
			from:   mustParseInt(s[2]),
		})
	}

	m.stacks = stacks
	m.instructions = instructions

	return m
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
