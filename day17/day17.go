// Package day17 solves the day17 puzzle.
package day17

import (
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
	"golang.org/x/exp/slices"
)

type direction byte

const (
	lt direction = '<'
	rt direction = '>'
)

type solution struct {
	dir []direction
}

func (sol *solution) Parse(s []byte) error {
	sol.dir = nil
	for _, b := range s {
		sol.dir = append(sol.dir, direction(b))
	}
	return nil
}

type scanline uint16

// slices of scanlines are ordered bottom to top.

type rock []scanline

var rocks = []rock{{
	0b0001111000000000,
}, {
	0b0000100000000000,
	0b0001110000000000,
	0b0000100000000000,
}, { // bottom to top
	0b0001110000000000,
	0b0000010000000000,
	0b0000010000000000,
}, {
	0b0001000000000000,
	0b0001000000000000,
	0b0001000000000000,
	0b0001000000000000,
}, {
	0b0001100000000000,
	0b0001100000000000,
}}

const walls scanline = 0b1000000010000000

func (r rock) shift(d direction) rock {
	out := slices.Clone(r)
	for i := range out {
		switch d {
		case lt:
			out[i] <<= 1
		case rt:
			out[i] >>= 1
		}
	}
	return out
}

func (r rock) collides(chamber []scanline, pos int) bool {
	if pos < 0 {
		return true
	}
	for i, s := range r {
		if s&walls != 0 || len(chamber) > i+pos && s&chamber[i+pos] != 0 {
			return true
		}
	}
	return false
}

type state struct {
	rockpos int
	dirpos  int
	floor   int
	chamber []scanline
}

func (sol solution) step(s *state) {
	r := rocks[s.rockpos]
	s.rockpos = (s.rockpos + 1) % len(rocks)
	pos := len(s.chamber) + 3
	for ; ; pos-- {
		n := r.shift(sol.dir[s.dirpos])
		s.dirpos = (s.dirpos + 1) % len(sol.dir)
		if !n.collides(s.chamber, pos) {
			r = n
		}
		if r.collides(s.chamber, pos-1) {
			break
		}
	}
	if pos+len(r) > len(s.chamber) {
		s.chamber = slices.Grow(s.chamber, pos+len(r)-len(s.chamber))[:pos+len(r)]
	}
	for _, sl := range r {
		s.chamber[pos] |= sl
		if s.chamber[pos] == 0b0111111100000000 {
			s.floor += pos
			s.chamber = s.chamber[pos:]
			pos = 0
		}
		pos++
	}
}

type compactState struct {
	rockpos int
	dirpos  int
	chamber [4]scanline
}

func (s *state) compactState() (compactState, bool) {
	if len(s.chamber) > 4 {
		return compactState{}, false
	}
	cs := compactState{
		rockpos: s.rockpos,
		dirpos:  s.dirpos,
	}
	copy(cs.chamber[:], s.chamber)
	return cs, true
}

func (sol solution) simulate(count int) int {
	s := state{}
	type counts struct {
		count int
		floor int
	}
	seen := map[compactState]counts{}
	for ; count > 0; count-- {
		sol.step(&s)
		if cs, ok := s.compactState(); ok {
			old, ok := seen[cs]
			if !ok {
				seen[cs] = counts{
					count: count,
					floor: s.floor,
				}
				continue
			}
			// we found a repeated state, so we can fast forward
			jump := count / (old.count - count)
			s.floor += jump * (s.floor - old.floor)
			count -= jump * (old.count - count)
			break
		}
	}
	for ; count > 0; count-- {
		sol.step(&s)
	}
	return s.floor + len(s.chamber)
}

func (sol solution) Part1() {
	fmt.Println("Part 1", sol.simulate(2022))
}

func (sol solution) Part2() {
	fmt.Println("Part 2", sol.simulate(1000000000000))
}

func init() {
	aoc.RegisterSolution("17", &solution{})
}
