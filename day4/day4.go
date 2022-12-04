// Package day4 solves the day4 puzzle.
package day4

import (
	"bytes"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type assignment struct {
	s, e int
}

func parseAssignment(p []byte) ([2]assignment, error) {
	var a, b assignment
	if _, err := fmt.Sscanf(string(p), "%d-%d,%d-%d", &a.s, &a.e, &b.s, &b.e); err != nil {
		return [2]assignment{}, nil
	}
	return [2]assignment{a, b}, nil
}

func parseAssignments(p []byte) ([][2]assignment, error) {
	lines := bytes.Split(p, []byte{'\n'})
	as := [][2]assignment(nil)
	for _, line := range lines {
		a, err := parseAssignment(line)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, nil
}

type solution struct{}

func (solution) Part1(p []byte) error {
	as, err := parseAssignments(p)
	if err != nil {
		return err
	}
	count := 0
	for _, pair := range as {
		a, b := pair[0], pair[1]
		if a.s <= b.s && a.e >= b.e {
			count++
		} else if b.s <= a.s && b.e >= a.e {
			count++
		}
	}
	fmt.Println("Part 1", count)
	return nil
}

func (solution) Part2(p []byte) error {
	as, err := parseAssignments(p)
	if err != nil {
		return err
	}
	count := 0
	for _, pair := range as {
		a, b := pair[0], pair[1]
		if a.s <= b.e && a.e >= b.s {
			count++
		}
	}
	fmt.Println("Part 2", count)
	return nil
}

func init() {
	aoc.RegisterSolution("4", solution{})
}
