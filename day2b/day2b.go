// Package day2b solves the day2 puzzle.
package day2b

import (
	"fmt"
	"strings"

	"github.com/nerdatmath/aoc2022/aoc"
)

func makeMap(f func(y, t, o byte) string) map[string]int {
	// rock=0, paper, scissors
	// loss=0, draw, win
	m := map[string]int{}
	for y := byte(0); y < 3; y++ {
		for t := byte(0); t < 3; t++ {
			o := (y - t + 4) % 3
			m[f(y, t, o)] = int(o*3 + y + 1)
		}
	}
	return m
}

type solution struct {
	lines []string
}

func (sol *solution) Parse(s []byte) error {
	sol.lines = strings.Split(string(s), "\n")
	return nil
}

func (sol solution) Part1() {
	m := makeMap(func(y, t, o byte) string { return string('A'+t) + " " + string('X'+y) })
	scores := aoc.Map(func(k string) int { return m[k] }, sol.lines)
	fmt.Println("Part 1", aoc.Sum(scores))
}

func (sol solution) Part2() {
	m := makeMap(func(y, t, o byte) string { return string('A'+t) + " " + string('X'+o) })
	scores := aoc.Map(func(k string) int { return m[k] }, sol.lines)
	fmt.Println("Part 2", aoc.Sum(scores))
}

func init() {
	aoc.RegisterSolution("2b", &solution{})
}
