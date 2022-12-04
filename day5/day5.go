// Package day5 solves the day5 puzzle.
package day5

import (
	"errors"

	"github.com/nerdatmath/aoc2022/aoc"
)

type solution struct{}

func (solution) Part1(p []byte) error {
	return errors.New("notimplemented")
}

func (solution) Part2(p []byte) error {
	return errors.New("notimplemented")
}

func init() {
	aoc.RegisterSolution("5", solution{})
}
