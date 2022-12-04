// Package day4 solves the day4 puzzle.
package day4

import (
	"bytes"
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

func parsePair(p []byte) (pair, error) {
	parts := bytes.Split(p, []byte{','})
	if len(parts) != 2 {
		return pair{}, errors.New("invalid pair")
	}
	a, err := parseAssignment(parts[0])
	if err != nil {
		return pair{}, err
	}
	b, err := parseAssignment(parts[1])
	if err != nil {
		return pair{}, err
	}
	return pair{a, b}, nil
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

func parsePairs(p []byte) ([]pair, error) {
	lines := bytes.Split(p, []byte{'\n'})
	ps := []pair(nil)
	for _, line := range lines {
		pr, err := parsePair(line)
		if err != nil {
			return nil, err
		}
		ps = append(ps, pr)
	}
	return ps, nil
}

type solution struct{}

func (solution) Part1(p []byte) error {
	ps, err := parsePairs(p)
	if err != nil {
		return err
	}
	count := 0
	for _, pair := range ps {
		if pair.fullyContained() {
			count++
		}
	}
	fmt.Println("Part 1", count)
	return nil
}

func (solution) Part2(p []byte) error {
	ps, err := parsePairs(p)
	if err != nil {
		return err
	}
	count := 0
	for _, pair := range ps {
		if pair.overlapping() {
			count++
		}
	}
	fmt.Println("Part 1", count)
	return nil
}

func init() {
	aoc.RegisterSolution("4", solution{})
}
