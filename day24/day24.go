// Package day24 solves the day24 puzzle.
package day24

import (
	"bytes"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type point struct {
	x, y int
}

type direction struct {
	dx, dy int
}

func (d direction) Mul(n int) direction {
	d.dx *= n
	d.dy *= n
	return d
}

func (p point) Add(d direction) point {
	return point{x: p.x + d.dx, y: p.y + d.dy}
}

func (p point) Wrap(offset, size direction) point {
	// defer func(old point) {
	// 	fmt.Printf("%v.Wrap(%v, %v) = %v\n", old, offset, size, p)
	// }(p)
	p = p.Add(offset.Mul(-1))
	p.x, p.y = p.x%size.dx, p.y%size.dy
	if p.x < 0 {
		p.x += size.dx
	}
	if p.y < 0 {
		p.y += size.dy
	}
	p = p.Add(offset)
	return p
}

type solution struct {
	field [][]byte
}

var directions = map[byte]direction{
	'#': {dx: 0, dy: 0},
	'<': {dx: -1, dy: 0},
	'>': {dx: +1, dy: 0},
	'^': {dx: 0, dy: -1},
	'v': {dx: 0, dy: +1},
}

func (sol *solution) Parse(s []byte) error {
	sol.field = bytes.Split(s, []byte{'\n'})
	return nil
}

func (s *solution) lookup(p point) byte {
	if p.y < 0 || p.y >= len(s.field) || p.x < 0 || p.x >= len(s.field[p.y]) {
		return '#'
	}
	return s.field[p.y][p.x]
}

func (s *solution) start() point {
	i := bytes.IndexByte(s.field[0], '.')
	return point{i, 0}
}

func (s *solution) end() point {
	i := bytes.IndexByte(s.field[len(s.field)-1], '.')
	return point{i, len(s.field) - 1}
}

func (s *solution) offset() direction {
	return direction{1, 1}
}

func (s *solution) size() direction {
	h := len(s.field) - 2
	w := len(s.field[0]) - 2
	return direction{dx: w, dy: h}
}

func (s *solution) blocked(p point, t int) bool {
	return s.marker(p, t) != '.'
}

func (s *solution) marker(p point, t int) rune {
	if s.lookup(p) == '#' {
		return '#'
	}
	if p.x < s.offset().dx || p.y < s.offset().dy {
		return '.'
	}
	if p.x >= s.offset().dx+s.size().dx || p.y >= s.offset().dy+s.size().dy {
		return '.'
	}
	out := '.'
	for arrow, d := range directions {
		if (d == direction{}) {
			continue
		}
		b := p.Add(d.Mul(-t)).Wrap(s.offset(), s.size())
		if s.lookup(b) == arrow {
			switch out {
			case '.':
				out = rune(arrow)
			case '<', '>', '^', 'v':
				out = '2'
			case '2':
				out = '3'
			case '3':
				out = '4'
			}
		}
	}
	return out
}

var printProgress = false

func (s *solution) printField(t int, active mapset.Set[point]) {
	if !printProgress {
		return
	}
	fmt.Println()
	fmt.Println("Time", t)
	for y := 0; y < s.size().dy+2; y++ {
		for x := 0; x < s.size().dx+2; x++ {
			p := point{x, y}
			if active.Contains(p) {
				fmt.Print("E")
			} else {
				fmt.Print(string(s.marker(point{x, y}, t)))
			}
		}
		fmt.Println()
	}
}

func (s *solution) seek(t int, start, end point) int {
	active := mapset.NewSet(start)
	visited := mapset.NewSet[point]()
	for ; ; t++ {
		visited = visited.Union(active)
		next := mapset.NewSet[point]()
		for n := range active.Iter() {
			for _, d := range directions {
				p := n.Add(d)
				if s.blocked(p, t) {
					continue
				}
				if p == end {
					return t
				}
				next.Add(p)
			}
		}
		active = next
		s.printField(t, active)
	}
}

func (sol solution) Part1() {
	start := sol.start()
	end := sol.end()
	t := sol.seek(0, start, end)
	fmt.Println("Part 1", t)
}

func (sol solution) Part2() {
	start := sol.start()
	end := sol.end()
	t := sol.seek(0, start, end)
	t = sol.seek(t, end, start)
	t = sol.seek(t, start, end)
	fmt.Println("Part 2", t)
}

func init() {
	aoc.RegisterSolution("24", &solution{})
}
