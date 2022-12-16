// Package day14 solves the day14 puzzle.
package day14

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type point struct {
	x, y int
}

func parsePoint(s []byte) (point, error) {
	words := bytes.Split(s, []byte{','})
	if len(words) != 2 {
		return point{}, errors.New("bad point")
	}
	x, err := strconv.Atoi(string(words[0]))
	if err != nil {
		return point{}, err
	}
	y, err := strconv.Atoi(string(words[1]))
	if err != nil {
		return point{}, err
	}
	return point{x, y}, nil
}

func inLine(a, b, c point) bool {
	v := a.y <= b.y && b.y <= c.y || a.y >= b.y && b.y >= c.y
	h := a.x <= b.x && b.x <= c.x || a.x >= b.x && b.x >= c.x
	return v && h
}

type path []point

func parsePath(s []byte) (path, error) {
	ps, err := aoc.ParseDelimited(s, parsePoint, []byte(" -> "))
	return path(ps), err
}

func onPath(p point, pth path) bool {
	for i := 0; i < len(pth)-1; i++ {
		if inLine(pth[i], p, pth[i+1]) {
			return true
		}
	}
	return false
}

type solution struct {
	paths []path
}

func (sol *solution) Parse(s []byte) error {
	paths, err := aoc.ParseLines(s, parsePath)
	sol.paths = paths
	return err
}

func (sol solution) occupied(p point) bool {
	for _, pth := range sol.paths {
		if onPath(p, pth) {
			return true
		}
	}
	return false
}

func (sol solution) drop(p point, check, occupied func(point) bool) (point, bool) {
outer:
	for check(p) {
		for _, x := range []int{p.x, p.x - 1, p.x + 1} {
			n := point{x, p.y + 1}
			if !occupied(n) {
				p = n
				continue outer
			}
		}
		return p, true
	}
	return p, false
}

func (sol solution) Part1() {
	maxy := 0
	for _, pth := range sol.paths {
		for _, p := range pth {
			if p.y > maxy {
				maxy = p.y
			}
		}
	}
	start := point{500, 0}
	sand := mapset.NewSet[point]()
	occupied := func(p point) bool {
		return sand.Contains(p) || sol.occupied(p)
	}
	check := func(p point) bool {
		return p.y < maxy
	}
	for {
		p, ok := sol.drop(start, check, occupied)
		if !ok {
			break
		}
		sand.Add(p)
	}
	fmt.Println("Part 1", sand.Cardinality())
}

func (sol solution) Part2() {
	maxy := 0
	for _, pth := range sol.paths {
		for _, p := range pth {
			if p.y > maxy {
				maxy = p.y
			}
		}
	}
	maxy += 2
	start := point{500, 0}
	sand := mapset.NewSet[point]()
	occupied := func(p point) bool {
		return p.y == maxy || sand.Contains(p) || sol.occupied(p)
	}
	check := func(p point) bool {
		return true
	}
	for {
		p, ok := sol.drop(start, check, occupied)
		if !ok {
			break
		}
		sand.Add(p)
		if p == start {
			break
		}
	}
	fmt.Println("Part 2", sand.Cardinality())
}

func init() {
	aoc.RegisterSolution("14", &solution{})
}
