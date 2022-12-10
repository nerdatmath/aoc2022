// Package day8 solves the day8 puzzle.
package day8

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
	aoc.RegisterSolution("8", &solution{})
}
