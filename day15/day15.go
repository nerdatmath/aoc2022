// Package day15 solves the day15 puzzle.
package day15

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
	"golang.org/x/exp/slices"
)

type point struct{ x, y int }

type sensor struct{ s, b point }

func parseSensor(s []byte) (sensor, error) {
	var sns sensor
	_, err := fmt.Sscanf(string(s), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sns.s.x, &sns.s.y, &sns.b.x, &sns.b.y)
	return sns, err
}

type solution struct {
	sensors []sensor
}

func (sol *solution) Parse(s []byte) error {
	sensors, err := aoc.ParseLines(s, parseSensor)
	sol.sensors = sensors
	return err
}

type interval struct{ start, end int }

func size(ivls []interval) int {
	slices.SortFunc(ivls, func(a, b interval) bool { return a.start < b.start })
	size := 0
	if len(ivls) == 0 {
		return 0
	}
	start, end := ivls[0].start, ivls[0].end
	ivls = ivls[1:]
	for _, ivl := range ivls {
		if ivl.start <= end {
			if ivl.end > end {
				end = ivl.end
			}
			continue
		}
		size += end - start + 1
		start, end = ivl.start, ivl.end
	}
	size += end - start + 1
	return size
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func manhattan(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func (sol solution) Part1() {
	row := 2000000
	ivls := []interval{}
	beacons := mapset.NewSet[int]()
	for _, s := range sol.sensors {
		if s.b.y == row {
			beacons.Add(s.b.x)
		}
		d := manhattan(s.s, s.b) - abs(s.s.y-row)
		if d < 0 {
			continue
		}
		ivl := interval{s.s.x - d, s.s.x + d}
		ivls = append(ivls, ivl)
	}
	s := size(ivls) - beacons.Cardinality()
	fmt.Println("Part 1", s)
}

func covered(r interval, ivls []interval) bool {
	slices.SortFunc(ivls, func(a, b interval) bool { return a.start < b.start })
	for _, ivl := range ivls {
		if ivl.start > r.start {
			break
		}
		if ivl.end >= r.end {
			return true
		}
		if ivl.end >= r.start {
			r.start = ivl.end + 1
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func in(x int, ivl interval) bool {
	return x >= ivl.start && x <= ivl.end
}

func intersect(a, b interval) (interval, bool) {
	i := interval{start: max(a.start, b.start), end: min(a.end, b.end)}
	return i, i.start <= i.end
}

func (sol solution) xcovered(x int, r interval) bool {
	ivls := []interval{}
	for _, s := range sol.sensors {
		if s.b.x == x && in(s.b.y, r) {
			ivls = append(ivls, interval{s.b.y, s.b.y})
		}
		d := manhattan(s.s, s.b) - abs(s.s.x-x)
		if d < 0 {
			continue
		}
		if ivl, ok := intersect(r, interval{s.s.y - d, s.s.y + d}); ok {
			ivls = append(ivls, ivl)
		}
	}
	return covered(r, ivls)
}

func (sol solution) ycovered(y int, r interval) bool {
	ivls := []interval{}
	for _, s := range sol.sensors {
		if s.b.y == y && in(s.b.x, r) {
			ivls = append(ivls, interval{s.b.x, s.b.x})
		}
		d := manhattan(s.s, s.b) - abs(s.s.y-y)
		if d < 0 {
			continue
		}
		if ivl, ok := intersect(r, interval{s.s.x - d, s.s.x + d}); ok {
			ivls = append(ivls, ivl)
		}
	}
	return covered(r, ivls)
}

func (sol solution) Part2() {
	max := 4000000
	r := interval{0, 4000000}
	b := point{}
	foundx, foundy := false, false
	for x := 0; x <= max; x++ {
		if !sol.xcovered(x, r) {
			b.x = x
			foundx = true
			break
		}
	}
	if !foundx {
		panic("couldn't find a beacon")
	}
	for y := 0; y <= max; y++ {
		if !sol.ycovered(y, r) {
			b.y = y
			foundy = true
			break
		}
	}
	if !foundy {
		panic("couldn't find a beacon")
	}
	fmt.Println("Part 2", b.x*4000000+b.y)
}

func init() {
	aoc.RegisterSolution("15", &solution{})
}
