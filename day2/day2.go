// Package day2 solves the day2 puzzle.
package day2

import (
	"errors"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type entry struct {
	abc int // 0..2
	xyz int // 0..2
}

func parseEntry(s []byte) (entry, error) {
	if len(s) != 3 || s[1] != ' ' {
		return entry{}, errors.New("invalid line")
	}
	abc, xyz := int(s[0]-'A'), int(s[2]-'X')
	switch abc {
	case 0, 1, 2:
	default:
		return entry{}, errors.New("invalid line")
	}
	switch xyz {
	case 0, 1, 2:
	default:
		return entry{}, errors.New("invalid line")
	}
	return entry{abc, xyz}, nil
}

// rock=0, paper, scissors
// loss=0, draw, win

func scoreV1(e entry) int {
	return e.xyz + 1 + (e.xyz-e.abc+4)%3*3
}

func scoreV2(e entry) int {
	return (e.abc+e.xyz+2)%3 + 1 + e.xyz*3
}

type solution struct {
	entries []entry
}

func (sol *solution) Parse(s []byte) error {
	es, err := aoc.ParseLines(s, parseEntry)
	sol.entries = es
	return err
}

func (sol solution) Part1() {
	for i := 0; i < 9; i++ {
		abc := i % 3
		xyz := i / 3
		fmt.Println(abc, xyz, scoreV1(entry{abc, xyz}))
	}
	scores := aoc.Map(scoreV1, sol.entries)
	fmt.Println("Part 1", aoc.Sum(scores))
}

func (sol solution) Part2() {
	scores := aoc.Map(scoreV2, sol.entries)
	fmt.Println("Part 2", aoc.Sum(scores))
}

func init() {
	aoc.RegisterSolution("2", &solution{})
}
