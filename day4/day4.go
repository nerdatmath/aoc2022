// Package day4 solves the day4 puzzle.
package day4

import (
	"errors"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type assignment struct {
	s, e int
}

func parseAssignment(p []byte) (assignment, error) {
	var a assignment
	if _, err := fmt.Sscanf(string(p), "%d-%d", &a.s, &a.e); err != nil {
		return assignment{}, err
	}
	return a, nil
}

type pair struct {
	a, b assignment
}

func parsePair(s []byte) (pair, error) {
	as, err := aoc.ParseDelimited(s, parseAssignment, []byte{','})
	if err == nil && len(as) != 2 {
		err = errors.New("invalid pair")
	}
	if err != nil {
		return pair{}, err
	}
	return pair{as[0], as[1]}, nil
}

func contains(a, b assignment) bool {
	return a.s <= b.s && a.e >= b.e
}

func overlaps(a, b assignment) bool {
	return a.s <= b.e && a.e >= b.s
}

func (p pair) fullyContained() bool {
	return contains(p.a, p.b) || contains(p.b, p.a)
}

func (p pair) overlapping() bool {
	return overlaps(p.a, p.b)
}

type solution struct {
	pairs []pair
}

func (sol *solution) Parse(s []byte) error {
	pairs, err := aoc.ParseLines(s, parsePair)
	sol.pairs = pairs
	return err
}

func (sol solution) Part1() {
	fmt.Println("Part 1", len(aoc.Filter(pair.fullyContained, sol.pairs)))
}

func (sol solution) Part2() {
	fmt.Println("Part 2", len(aoc.Filter(pair.overlapping, sol.pairs)))
}

func init() {
	aoc.RegisterSolution("4", &solution{})
}
