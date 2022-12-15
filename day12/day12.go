// Package day12 solves the day12 puzzle.
package day12

import (
	"bytes"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type node [2]int

type solution struct {
	start, end node
	edges      map[node][]node
}

func (sol *solution) addEdge(f, t node) {
	sol.edges[f] = append(sol.edges[f], t)
}

func (sol *solution) Parse(s []byte) error {
	sol.edges = make(map[node][]node)
	lines := bytes.Split(s, []byte{'\n'})
	for y := range lines {
		for x := range lines[y] {
			n := node{x, y}
			switch lines[y][x] {
			case 'S':
				sol.start = n
				lines[y][x] = 'a'
			case 'E':
				sol.end = n
				lines[y][x] = 'z'
			}
		}
	}
	for y := range lines {
		for x := range lines[y] {
			h := lines[y][x]
			n := node{x, y}
			if y > 0 && lines[y-1][x] <= h+1 {
				sol.addEdge(n, node{x, y - 1})
			}
			if y > 0 && h <= lines[y-1][x]+1 {
				sol.addEdge(node{x, y - 1}, n)
			}
			if x > 0 && lines[y][x-1] <= h+1 {
				sol.addEdge(n, node{x - 1, y})
			}
			if x > 0 && h <= lines[y][x-1]+1 {
				sol.addEdge(node{x - 1, y}, n)
			}
		}
	}
	return nil
}

func (sol solution) Part1() {
	visited := map[node]struct{}{sol.start: {}}
	active := map[node]struct{}{sol.start: {}}
	steps := 0
	for {
		// fmt.Println("step", steps)
		if _, ok := visited[sol.end]; ok {
			break
		}
		next := map[node]struct{}{}
		for n := range active {
			for _, p := range sol.edges[n] {
				if _, ok := visited[p]; !ok {
					// fmt.Println("visiting", p)
					next[p] = struct{}{}
					visited[p] = struct{}{}
				}
			}
		}
		active = next
		steps++
	}
	fmt.Println("Part 1", steps)
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("12", &solution{})
}
