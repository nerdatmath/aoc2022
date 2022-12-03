// Package day2 solves the day2 puzzle.
package day2

import (
	"bytes"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type rps int

const (
	rock rps = iota
	paper
	scissors
)

func (r rps) score() int {
	return map[rps]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}[r]
}

func (r rps) better() rps {
	return map[rps]rps{
		rock:     paper,
		paper:    scissors,
		scissors: rock,
	}[r]
}

func (r rps) worse() rps {
	return map[rps]rps{
		rock:     scissors,
		paper:    rock,
		scissors: paper,
	}[r]
}

type round struct {
	them, you rps
}

func (r round) score() int {
	return r.you.score() + r.outcome().score()
}

type roundV2 struct {
	them    rps
	outcome outcome
}

func (r roundV2) you() rps {
	switch r.outcome {
	case draw:
		return r.them
	case win:
		return r.them.better()
	case loss:
		return r.them.worse()
	}
	return 0
}

func (r roundV2) score() int {
	return r.you().score() + r.outcome.score()
}

type outcome int

const (
	loss outcome = iota
	draw
	win
)

func (r round) outcome() outcome {
	if r.them == r.you {
		return draw
	} else if r.them.better() == r.you {
		return win
	} else {
		return loss
	}
}

func (o outcome) score() int {
	return map[outcome]int{win: 6, draw: 3, loss: 0}[o]
}

func parseOutcome(s string) (outcome, error) {
	outcome, ok := map[string]outcome{"X": loss, "Y": draw, "Z": win}[s]
	if !ok {
		return 0, fmt.Errorf("unknown outcome: %q", s)
	}
	return outcome, nil
}

func parseTheirMove(s string) (rps, error) {
	rps, ok := map[string]rps{"A": rock, "B": paper, "C": scissors}[s]
	if !ok {
		return 0, fmt.Errorf("unknown move for them: %q", s)
	}
	return rps, nil
}

func parseYourMove(s string) (rps, error) {
	rps, ok := map[string]rps{"X": rock, "Y": paper, "Z": scissors}[s]
	if !ok {
		return 0, fmt.Errorf("unknown move for you: %q", s)
	}
	return rps, nil
}

func parseStrategy(b []byte) ([]round, error) {
	lines := bytes.Split(b, []byte("\n"))
	rounds := []round(nil)
	for _, line := range lines {
		them, err := parseTheirMove(string(line[0:1]))
		if err != nil {
			return nil, err
		}
		you, err := parseYourMove(string(line[2:3]))
		if err != nil {
			return nil, err
		}
		rounds = append(rounds, round{them: them, you: you})
	}
	return rounds, nil
}

func parseStrategyV2(b []byte) ([]roundV2, error) {
	lines := bytes.Split(b, []byte("\n"))
	rounds := []roundV2(nil)
	for _, line := range lines {
		them, err := parseTheirMove(string(line[0:1]))
		if err != nil {
			return nil, err
		}
		outcome, err := parseOutcome(string(line[2:3]))
		if err != nil {
			return nil, err
		}
		rounds = append(rounds, roundV2{them: them, outcome: outcome})
	}
	return rounds, nil
}

type solution struct{}

func (solution) Part1(p []byte) error {
	rounds, err := parseStrategy(p)
	if err != nil {
		return err
	}
	score := 0
	for _, r := range rounds {
		score += r.score()
	}
	fmt.Println("Part 1", score)
	return nil
}

func (solution) Part2(p []byte) error {
	rounds, err := parseStrategyV2(p)
	if err != nil {
		return err
	}
	score := 0
	for _, r := range rounds {
		score += r.score()
	}
	fmt.Println("Part 2", score)
	return nil
}

func init() {
	aoc.RegisterSolution("2", solution{})
}
