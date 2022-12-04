// Package day3 solves the day3 puzzle.
package day3

import (
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type item byte

func (i item) priority() int {
	if i >= 'a' {
		return int(i) - 'a' + 1
	}
	return int(i) - 'A' + 27
}

type rucksack struct {
	compartments [2][]item
}

func (r rucksack) items() []item {
	return append(r.compartments[0], r.compartments[1]...)
}

func (r rucksack) dupe() item {
	i := intersect(r.compartments[0], r.compartments[1])
	if len(i) != 1 {
		panic("no dupe or too many")
	}
	return i[0]
}

func intersect(a, b []item) []item {
	m := map[item]struct{}{}
	for _, i := range a {
		m[i] = struct{}{}
	}
	intersection := []item(nil)
	for _, i := range b {
		if _, ok := m[i]; ok {
			intersection = append(intersection, i)
			delete(m, i)
		}
	}
	return intersection
}

func parseRucksack(p []byte) (rucksack, error) {
	items := []item(nil)
	for _, b := range p {
		items = append(items, item(b))
	}
	l := len(items) / 2
	return rucksack{
		compartments: [2][]item{items[:l], items[l:]},
	}, nil
}

type solution struct {
	elves []rucksack
}

func (sol *solution) Parse(p []byte) error {
	elves, err := aoc.ParseLines(p, parseRucksack)
	sol.elves = elves
	return err
}

func (sol solution) Part1() {
	sum := aoc.Sum(aoc.Map(func(r rucksack) int { return r.dupe().priority() }, sol.elves))
	fmt.Println("Part 1", sum)
}

func (sol solution) Part2() {
	sum := 0
	for i := 0; i < len(sol.elves); i += 3 {
		v := intersect(intersect(sol.elves[i].items(), sol.elves[i+1].items()), sol.elves[i+2].items())
		if len(v) != 1 {
			panic("too few or too many")
		}
		sum += v[0].priority()
	}
	fmt.Println("Part 2", sum)
}

func init() {
	aoc.RegisterSolution("3", &solution{})
}
