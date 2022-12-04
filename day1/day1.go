// Package day1 solves the day1 puzzle.
package day1

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"

	"github.com/nerdatmath/aoc2022/aoc"
)

type elf struct {
	food []int
}

func (e elf) calories() int {
	return aoc.Sum(e.food)
}

func parseElf(s []byte) (elf, error) {
	food, err := aoc.ParseLines(s, func(s []byte) (int, error) { return strconv.Atoi(string(s)) })
	return elf{food: food}, err
}

type solution struct {
	elves []elf
}

func (sol *solution) Parse(s []byte) error {
	elves, err := aoc.ParseDelimited(s, parseElf, []byte("\n\n"))
	sol.elves = elves
	return err
}

func (sol solution) Part1() {
	fmt.Println("Part 1", aoc.Max(aoc.Map(elf.calories, sol.elves)))
}

func (sol solution) Part2() {
	cs := aoc.Map(elf.calories, sol.elves)
	slices.SortFunc(cs, func(a, b int) bool { return a > b })
	fmt.Println("Part 2", aoc.Sum(cs[0:3]))
}

func init() {
	aoc.RegisterSolution("1", &solution{})
}
