// Package day12 solves the day12 puzzle.
package day12

import (
	"bytes"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type node struct{ x, y int }

type solution struct {
	start, end node
	heights    map[node]byte
}

func (sol *solution) edges(n node) []node {
	h := sol.heights[n]
	edges := []node{}
	for _, p := range []node{{n.x, n.y - 1}, {n.x, n.y + 1}, {n.x - 1, n.y}, {n.x + 1, n.y}} {
		if ph, ok := sol.heights[p]; ok && ph <= h+1 {
			edges = append(edges, p)
		}
	}
	return edges
}

func (sol *solution) Parse(s []byte) error {
	sol.heights = make(map[node]byte)
	lines := bytes.Split(s, []byte{'\n'})
	n := node{}
	for n.y = range lines {
		for n.x = range lines[n.y] {
			h := lines[n.y][n.x]
			switch h {
			case 'S':
				sol.start = n
				sol.heights[n] = 'a'
			case 'E':
				sol.end = n
				sol.heights[n] = 'z'
			default:
				sol.heights[n] = h
			}
		}
	}
	return nil
}

func seek(start mapset.Set[node], end node, edges func(node) []node) int {
	active := start
	visited := start.Clone()
	var steps int
	for steps = 0; !visited.Contains(end); steps++ {
		next := mapset.NewSet[node]()
		for n := range active.Iter() {
			next = next.Union(mapset.NewSet(edges(n)...))
		}
		active = next.Difference(visited)
		visited = visited.Union(active)
	}
	return steps
}

func (sol solution) Part1() {
	steps := seek(mapset.NewSet(sol.start), sol.end, sol.edges)
	fmt.Println("Part 1", steps)
}

func (sol solution) Part2() {
	start := mapset.NewSet[node]()
	for n, h := range sol.heights {
		if h == 'a' {
			start.Add(n)
		}
	}
	steps := seek(start, sol.end, sol.edges)
	fmt.Println("Part 2", steps)
}

func init() {
	aoc.RegisterSolution("12", &solution{})
}
