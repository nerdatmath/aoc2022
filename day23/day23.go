// Package day23 solves the day23 puzzle.
package day23

import (
	"bytes"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type position struct {
	x, y int
}

func (p position) add(d direction) position {
	return position{x: p.x + d.dx, y: p.y + d.dy}
}

type elf struct {
	p position
}

func newElf(p position) elf {
	return elf{p: p}
}

type elves struct {
	mapset.Set[elf]
}

func newElves() elves {
	return elves{mapset.NewSet[elf]()}
}

func (es elves) blocked(p position) bool {
	return es.Contains(newElf(p))
}

func (es elves) Clone() elves {
	return elves{es.Set.Clone()}
}

type direction struct {
	dx, dy int
}

var (
	N  = direction{dx: 0, dy: -1}
	S  = direction{dx: 0, dy: +1}
	W  = direction{dx: -1, dy: 0}
	E  = direction{dx: +1, dy: 0}
	NW = direction{dx: -1, dy: -1}
	NE = direction{dx: +1, dy: -1}
	SW = direction{dx: -1, dy: +1}
	SE = direction{dx: +1, dy: +1}
)

type solution struct {
	elves elves
}

func (sol *solution) Parse(s []byte) error {
	elves := newElves()
	for y, row := range bytes.Split(s, []byte{'\n'}) {
		for x, c := range row {
			if c == '#' {
				elves.Add(newElf(position{x: x, y: y}))
			}
		}
	}
	sol.elves = elves
	return nil
}

func propose(e elf, es elves, preferred []direction) position {
	blocks := map[direction][]direction{
		N: {NW, N, NE},
		S: {SW, S, SE},
		W: {NW, W, SW},
		E: {NE, E, SE},
	}
	blocked := mapset.NewSet[direction]()
	for d, bs := range blocks {
		for _, b := range bs {
			if es.blocked(e.p.add(b)) {
				blocked.Add(d)
			}
		}
	}
	if blocked.Cardinality() == 0 {
		return e.p
	}
	for _, d := range preferred {
		if !blocked.Contains(d) {
			return e.p.add(d)
		}
	}
	return e.p
}

func step(es elves, preferred []direction) elves {
	proposed := map[elf]position{}
	counts := map[position]int{}
	for e := range es.Iter() {
		p := propose(e, es, preferred)
		proposed[e] = p
		counts[p]++
	}
	next := newElves()
	for e := range es.Iter() {
		p := proposed[e]
		if counts[p] == 1 {
			next.Add(newElf(p))
		} else {
			next.Add(e)
		}
	}
	return next
}

func box(es elves) (tl, br position) {
	es = es.Clone()
	e, _ := es.Pop()
	tl, br = e.p, e.p
	for es.Cardinality() > 0 {
		e, _ := es.Pop()
		if e.p.x < tl.x {
			tl.x = e.p.x
		}
		if e.p.x > br.x {
			br.x = e.p.x
		}
		if e.p.y < tl.y {
			tl.y = e.p.y
		}
		if e.p.y > br.y {
			br.y = e.p.y
		}
	}
	return tl, br
}

func (sol solution) Part1() {
	var preferred = []direction{N, S, W, E}
	es := sol.elves.Clone()
	for i := 0; i < 10; i++ {
		es = step(es, preferred)
		preferred = append(preferred[1:], preferred[0])
	}
	tl, br := box(es)
	area := (br.x - tl.x + 1) * (br.y - tl.y + 1)
	fmt.Println("Part 1", area-es.Cardinality())
}

func (sol solution) Part2() {
	var preferred = []direction{N, S, W, E}
	es := sol.elves.Clone()
	for i := 0; ; i++ {
		next := step(es, preferred)
		if next.Equal(es.Set) {
			fmt.Println("Part 2", i+1)
			return
		}
		es = next
		preferred = append(preferred[1:], preferred[0])
	}
}

func init() {
	aoc.RegisterSolution("23", &solution{})
}
