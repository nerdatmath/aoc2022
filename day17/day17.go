// Package day17 solves the day17 puzzle.
package day17

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
	aoc.RegisterSolution("17", &solution{})
}
