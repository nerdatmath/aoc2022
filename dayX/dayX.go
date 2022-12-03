// Package dayX solves the dayX puzzle.
package dayX

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
	aoc.RegisterSolution("X", solution{})
}
