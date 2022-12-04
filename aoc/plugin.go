// Package aoc provides a plugin mechanism for puzzle solutions.
package aoc

import (
	"fmt"
)

// Solution is the interface that puzzle solutions must implement.
type Solution interface {
	Parse(p []byte) error
	Part1()
	Part2()
}

var solutions = map[string]Solution{}

// RegisterSolution registers a solution to one of the puzzles.
// It is meant to be called from an init function.
func RegisterSolution(day string, solution Solution) {
	solutions[day] = solution
}

// Lookup looks up the solution for a given day.
func Lookup(day string) (Solution, error) {
	if s, ok := solutions[day]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("day %s solution is not implemented", day)
}
