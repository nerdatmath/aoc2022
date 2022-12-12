// Package day9 solves the day9 puzzle.
package day9

import (
	"fmt"
	"strconv"

	"github.com/nerdatmath/aoc2022/aoc"
)

type point struct {
	x, y int
}

func add(a, b point) point {
	return point{a.x + b.x, a.y + b.y}
}

func sub(a, b point) point {
	return point{a.x - b.x, a.y - b.y}
}

type direction byte

const (
	R direction = iota
	L
	U
	D
)

func parseDirection(s byte) (direction, error) {
	switch s {
	case 'R':
		return R, nil
	case 'L':
		return L, nil
	case 'U':
		return U, nil
	case 'D':
		return D, nil
	default:
		return 0, fmt.Errorf("invalid direction '%c'", s)
	}
}

func (d direction) delta() point {
	switch d {
	case R:
		return point{1, 0}
	case L:
		return point{-1, 0}
	case U:
		return point{0, -1}
	case D:
		return point{0, 1}
	default:
		panic("invalid direction")
	}
}

type move struct {
	dir  direction
	dist int
}

func parseMove(s []byte) (move, error) {
	dir, err := parseDirection(s[0])
	if err != nil {
		return move{}, err
	}
	dist, err := strconv.Atoi(string(s[2:]))
	if err != nil {
		return move{}, err
	}
	return move{dir, dist}, nil
}

type solution struct {
	moves []move
}

func (sol *solution) Parse(s []byte) error {
	moves, err := aoc.ParseLines(s, parseMove)
	if err != nil {
		return err
	}
	sol.moves = moves
	return nil
}

func (t *point) follow(h point) {
	d := sub(h, *t)
	switch d {
	case point{2, 0}:
		*t = add(*t, point{1, 0})
	case point{-2, 0}:
		*t = add(*t, point{-1, 0})
	case point{0, 2}:
		*t = add(*t, point{0, 1})
	case point{0, -2}:
		*t = add(*t, point{0, -1})
	case point{2, 2}, point{2, 1}, point{1, 2}:
		*t = add(*t, point{1, 1})
	case point{-2, 2}, point{-2, 1}, point{-1, 2}:
		*t = add(*t, point{-1, 1})
	case point{2, -2}, point{2, -1}, point{1, -2}:
		*t = add(*t, point{1, -1})
	case point{-2, -2}, point{-2, -1}, point{-1, -2}:
		*t = add(*t, point{-1, -1})
	default:
		if d.x < -1 || d.x > 1 || d.y < -1 || d.y > 1 {
			panic(fmt.Sprint("incomplete switch", d))
		}
	}
}

func visits(l int, moves []move) map[point]struct{} {
	s := make([]point, l)
	visited := map[point]struct{}{{}: {}}
	for _, m := range moves {
		for i := 0; i < m.dist; i++ {
			s[0] = add(s[0], m.dir.delta())
			for i := 1; i < len(s); i++ {
				s[i].follow(s[i-1])
			}
			visited[s[len(s)-1]] = struct{}{}
		}
	}
	return visited
}

func (sol solution) Part1() {
	fmt.Println("Part 1", len(visits(2, sol.moves)))
}

func (sol solution) Part2() {
	fmt.Println("Part 2", len(visits(10, sol.moves)))
}

func init() {
	aoc.RegisterSolution("9", &solution{})
}
