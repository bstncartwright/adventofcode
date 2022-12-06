package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	rounds := parse(input)

	var part1 int

	for _, round := range rounds {
		win := calculateWin(round.Opponent, round.Player)

		part1 += (int(win) + int(round.Player))
	}

	var part2 int

	for _, round := range rounds {
		var player Action
		switch round.Player {
		case Rock:
			// we need to lose
			switch round.Opponent {
			case Rock:
				player = Scissors
			case Paper:
				player = Rock
			case Scissors:
				player = Paper
			}
		case Paper:
			// we need to draw
			player = round.Opponent
		case Scissors:
			// we need to win
			switch round.Opponent {
			case Rock:
				player = Paper
			case Paper:
				player = Scissors
			case Scissors:
				player = Rock
			}
		}

		win := calculateWin(round.Opponent, player)

		part2 += (int(win) + int(player))

	}

	return part1, part2
}

type Action int

const (
	Rock     Action = 1
	Paper    Action = 2
	Scissors Action = 3
)

type Winner int

const (
	OpponentWins Winner = 0
	Draw         Winner = 3
	PlayerWins   Winner = 6
)

func calculateWin(opponent Action, player Action) Winner {
	switch {
	case opponent == player:
		return Draw
	case opponent == Rock && player == Scissors:
		return OpponentWins
	case opponent == Paper && player == Rock:
		return OpponentWins
	case opponent == Scissors && player == Paper:
		return OpponentWins
	case player == Rock && opponent == Scissors:
		return PlayerWins
	case player == Paper && opponent == Rock:
		return PlayerWins
	case player == Scissors && opponent == Paper:
		return PlayerWins
	default:
		panic("unreachable")
	}
}

func parseAction(a string) Action {
	switch strings.ToLower(a) {
	case "a", "x":
		return Rock
	case "b", "y":
		return Paper
	case "c", "z":
		return Scissors
	default:
		panic(fmt.Sprintf("unknown action: %s", a))
	}
}

type Round struct {
	Opponent Action
	Player   Action
}

func parse(s string) []Round {
	lines := strings.Split(s, "\n")

	var rounds []Round

	for _, line := range lines {
		actions := strings.Split(line, " ")
		if len(actions) != 2 {
			panic(fmt.Sprintf("invalid line: %s", line))
		}

		rounds = append(rounds, Round{
			Opponent: parseAction(actions[0]),
			Player:   parseAction(actions[1]),
		})
	}

	return rounds
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
