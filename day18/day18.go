// Package day18 solves the day18 puzzle.
package day18

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type point struct {
	x, y, z int
}

func parsePoint(s []byte) (point, error) {
	p := point{}
	_, err := fmt.Sscanf(string(s), "%d,%d,%d", &p.x, &p.y, &p.z)
	return p, err
}

type solution struct {
	points []point
}

func (sol *solution) Parse(s []byte) error {
	points, err := aoc.ParseLines(s, parsePoint)
	sol.points = points
	return err
}

func neighbors(p point) []point {
	return []point{
		{p.x - 1, p.y, p.z},
		{p.x + 1, p.y, p.z},
		{p.x, p.y - 1, p.z},
		{p.x, p.y + 1, p.z},
		{p.x, p.y, p.z - 1},
		{p.x, p.y, p.z + 1},
	}
}

func (sol solution) Part1() {
	lava := mapset.NewSet(sol.points...)
	area := 0
	for _, p := range sol.points {
		for _, n := range neighbors(p) {
			if !lava.Contains(n) {
				area++
			}
		}
	}
	fmt.Println("Part 1", area)
}

// The entire space fits in a cube (0,0,0)...(20,20,20);
// To be safe we extend it by 5 in all directions.
func inBounds(p point) bool {
	return -5 <= p.x && p.x <= 25 && -5 <= p.y && p.y <= 25 && -5 <= p.z && p.z <= 25
}

func dfs(active mapset.Set[point], blocked mapset.Set[point], inBounds func(point) bool) mapset.Set[point] {
	visited := mapset.NewSet[point]()
	for active.Cardinality() != 0 {
		visited = visited.Union(active)
		next := []point{}
		for n := range active.Iter() {
			for _, p := range neighbors(n) {
				if !inBounds(p) {
					continue
				}
				next = append(next, p)
			}
		}
		active = mapset.NewSet(next...).Difference(visited).Difference(blocked)
	}
	return visited
}

func (sol solution) Part2() {
	// Flood fill to find all the exterior cubes.
	// Everything else is interior.
	// We know (0,0,0) is in the exterior.
	lava := mapset.NewSet(sol.points...)
	exterior := dfs(mapset.NewSet(point{0, 0, 0}), lava, inBounds)
	area := 0
	for _, p := range sol.points {
		for _, n := range neighbors(p) {
			if exterior.Contains(n) {
				area++
			}
		}
	}
	fmt.Println("Part 2", area)
}

func init() {
	aoc.RegisterSolution("18", &solution{})
}
