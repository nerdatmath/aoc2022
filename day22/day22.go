// Package day22 solves the day22 puzzle.
package day22

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/nerdatmath/aoc2022/aoc"
)

const (
	wall  byte = '#'
	blank byte = ' '
	open  byte = '.'
)

type position struct {
	row, col int
}

type self struct {
	pos position
	d   position
}

type field [][]byte

func (f field) start() position {
	p := position{}
	for ; f.at(p) == blank; p.col++ {
	}
	return p
}

func (f field) at(p position) byte {
	if p.row < 0 || p.row >= len(f) || p.col < 0 || p.col >= len(f[p.row]) {
		return blank
	}
	return f[p.row][p.col]
}

func (f field) step(p, d position) position {
	p = p.add(d)
	for f.at(p) == blank {
		if p.row >= len(f) {
			p.row = 0
		} else if p.row < 0 {
			p.row = len(f) - 1
		} else if p.col >= len(f[p.row]) {
			p.col = 0
		} else if p.col < 0 {
			p.col = len(f[p.row]) - 1
		} else {
			p = p.add(d)
		}
	}
	return p
}

func (p position) add(d position) position {
	return position{row: p.row + d.row, col: p.col + d.col}
}

func (s *self) forward(n int, f field) {
	for ; n > 0; n-- {
		next := f.step(s.pos, s.d)
		if f.at(next) == wall {
			// fmt.Println("wall at", next)
			break
		}
		// fmt.Println(next)
		s.pos = next
	}
}

func (s *self) turn(dir byte) {
	switch dir {
	case 'L':
		// fmt.Println("turn left")
		s.d.row, s.d.col = -s.d.col, s.d.row
	case 'R':
		// fmt.Println("turn right")
		s.d.row, s.d.col = s.d.col, -s.d.row
	}
}

func (s *self) location() (x, y int) {
	return s.pos.col + 1, s.pos.row + 1
}

func (s *self) facing() int {
	if s.d.col == 0 {
		return 2 - s.d.row // 3 or 1
	}
	return 1 - s.d.col // 2 or 0
}

func (s *self) password() int {
	x, y := s.location()
	facing := s.facing()
	// fmt.Println("x", x, "y", y, "facing", facing)
	return 1000*y + 4*x + facing
}

func parseField(s []byte) (field, error) {
	return bytes.Split(s, []byte{'\n'}), nil
}

type instruction func(*self, field)

func parseInstructions(s []byte) ([]instruction, error) {
	inst := []instruction{}
	for len(s) != 0 {
		switch s[0] {
		case 'L', 'R':
			dir := s[0]
			inst = append(inst, func(me *self, _ field) { me.turn(dir) })
			s = s[1:]
			continue
		}
		i := bytes.IndexAny(s, "LR")
		if i == -1 {
			i = len(s)
		}
		n, err := strconv.Atoi(string(s[:i]))
		if err != nil {
			return nil, err
		}
		inst = append(inst, func(s *self, f field) { s.forward(n, f) })
		s = s[i:]
	}
	return inst, nil
}

type solution struct {
	f            field
	instructions []instruction
}

func (sol *solution) Parse(s []byte) error {
	m, i, _ := bytes.Cut(s, []byte("\n\n"))
	f, err := parseField(m)
	if err != nil {
		return err
	}
	instructions, err := parseInstructions(i)
	if err != nil {
		return err
	}
	sol.f = f
	sol.instructions = instructions
	return nil
}

func (sol solution) Part1() {
	s := self{pos: sol.f.start(), d: position{col: 1}}
	for _, f := range sol.instructions {
		f(&s, sol.f)
	}
	fmt.Println("Part 1", s.password())
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("22", &solution{})
}
