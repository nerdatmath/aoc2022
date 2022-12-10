// Package day10 solves the day10 puzzle.
package day10

import (
	"errors"

	"github.com/nerdatmath/aoc2022/aoc"
)

type solution struct{}

func (sol *solution) Parse(s []byte) error {
	return errors.New("notimplemented")
}

func (sol solution) Part1() {
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("10", &solution{})
}
