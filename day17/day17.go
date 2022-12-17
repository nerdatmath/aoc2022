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

func (r rock) collides(chamber []scanline) bool {
	for i, s := range r {
		if s&walls != 0 || len(chamber) > i && s&chamber[i] != 0 {
			return true
		}
	}
	return false
}

func (sol solution) simulate(count int) int {
	chamber := []scanline{}
	time := 0
	floor := 0
	type compactState struct {
		rocknum, dirpos int
		chamber         [4]scanline
	}
	type counts struct {
		rocks int
		floor int
	}
	seen := map[compactState]counts{}
	foundRecurrence := false
	for i := 0; i < count; i++ {
		r := rocks[i%len(rocks)]
		// rock appears 3 rows above the top
		chamber = append(chamber, 0, 0, 0)
		var pos int
		for pos = len(chamber); pos >= 0 && !r.collides(chamber[pos:]); pos-- {
			n := r.shift(sol.dir[time%len(sol.dir)])
			time++
			if !n.collides(chamber[pos:]) {
				r = n
			}
		}
		pos++
		for _, s := range r {
			if len(chamber) <= pos {
				chamber = append(chamber, 0)
			}
			chamber[pos] |= s
			if chamber[pos] == 0b0111111100000000 {
				floor += pos + 1
				chamber = chamber[pos+1:]
				pos = 0
				continue
			}
			pos++
		}
		for len(chamber) > 0 && chamber[len(chamber)-1] == 0 {
			chamber = chamber[:len(chamber)-1]
		}
		if foundRecurrence {
			continue
		}
		if len(chamber) < 4 {
			// Check to see if our complete state has been seen before.
			state := compactState{
				rocknum: i % len(rocks),
				dirpos:  time % len(sol.dir),
			}
			copy(state.chamber[:], chamber)
			new := counts{
				rocks: i,
				floor: floor,
			}
			old, ok := seen[state]
			if !ok {
				seen[state] = new
				continue
			}
			// we found a repeated state, so we can fast forward
			foundRecurrence = true
			diff := counts{
				rocks: new.rocks - old.rocks,
				floor: new.floor - old.floor,
			}
			jump := (count - i) / diff.rocks
			floor += jump * diff.floor
			i += jump * diff.rocks
		}
	}
	return floor + len(chamber)
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
