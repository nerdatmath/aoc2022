// Package day24 solves the day24 puzzle.
package day24

import (
	"bytes"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type point struct {
	x, y int
}

type blizzard struct {
	pos   point
	delta point
}

type solution struct {
	start, end point
	blizzards  mapset.Set[blizzard]
}

func (sol *solution) Parse(s []byte) error {
	lines := bytes.Split(s, []byte{'\n'})
	i := bytes.IndexByte(lines[0], '.')
	sol.start = point{i, 0}
	i = bytes.IndexByte(lines[len(lines)-1], '.')
	sol.end = point{i, len(lines) - 1}
	sol.blizzards = mapset.NewSet[blizzard]()
	for i, line := range lines {
		for j, c := range line {
			pos := point{i, j}
			switch c {
			case '<':
				sol.blizzards.Add(blizzard{pos: pos, delta: point{-1, 0}})
			case '>':
				sol.blizzards.Add(blizzard{pos: pos, delta: point{1, 0}})
			case '^':
				sol.blizzards.Add(blizzard{pos: pos, delta: point{0, -1}})
			case 'v':
				sol.blizzards.Add(blizzard{pos: pos, delta: point{0, 1}})
			case '#': // walls are just non-moving blizzards
				sol.blizzards.Add(blizzard{pos: pos, delta: point{0, 0}})
			}
		}
	}
	return nil
}

type state struct {
	e         point
	blizzards mapset.Set[blizzard]
}

func edges(e point, blocked func(point) bool) []point {
	edges := []point{}
	for _, p := range []point{
		e, // wait
		{e.x - 1, e.y},
		{e.x + 1, e.y},
		{e.x, e.y - 1},
		{e.x, e.y + 1},
	} {
		if !blocked(p) {
			edges = append(edges, p)
		}
	}
	return edges
}

func (sol solution) Part1() {
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("24", &solution{})
}
