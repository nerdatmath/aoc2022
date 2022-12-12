// Package day10 solves the day10 puzzle.
package day10

import (
	"fmt"
	"strconv"

	"github.com/nerdatmath/aoc2022/aoc"
)

type crt struct {
	pos    int
	pixels [40]rune
}

type cpu struct {
	time       int
	x          int
	breakpoint int
	callback   func()
	crt        *crt
}

func newCpu() *cpu {
	return &cpu{x: 1, crt: &crt{}}
}

func (c *cpu) tick() {
	c.time++
	c.crt.update(c.x)
	if c.time == c.breakpoint {
		c.callback()
	}
}

type instruction interface {
	run(c *cpu)
}

type noop struct{}

func (noop) run(c *cpu) { c.tick() }

type addx struct{ v int }

func (i addx) run(c *cpu) { c.tick(); c.tick(); c.x += i.v }

func (c *crt) update(x int) {
	if x >= c.pos-1 && x <= c.pos+1 {
		c.pixels[c.pos] = '#'
	} else {
		c.pixels[c.pos] = '.'
	}
	c.pos++
	if c.pos >= len(c.pixels) {
		c.print()
		c.pos = 0
	}
}

func (c crt) print() {
	fmt.Println(string(c.pixels[:]))
}

func parseInstruction(s []byte) (instruction, error) {
	opcode := string(s[:4])
	switch opcode {
	case "noop":
		return noop{}, nil
	case "addx":
		v, err := strconv.Atoi(string(s[5:]))
		if err != nil {
			return nil, err
		}
		return addx{v: v}, nil
	default:
		return nil, fmt.Errorf("unknown instruction %q", string(s))
	}
}

type solution struct {
	program []instruction
}

func (sol *solution) Parse(s []byte) error {
	program, err := aoc.ParseLines(s, parseInstruction)
	sol.program = program
	return err
}

func signalStrength(c cpu) int {
	return c.time * c.x
}

func (sol solution) Part1() {
	c := newCpu()
	sum := 0
	c.breakpoint = 20
	c.callback = func() {
		ss := signalStrength(*c)
		fmt.Println("cycle", c.time, "x", c.x, "signal strength", ss)
		sum += signalStrength(*c)
		c.breakpoint += 40
	}
	for _, i := range sol.program {
		i.run(c)
	}
	fmt.Println("Part 1", sum)
}

func (sol solution) Part2() {
	c := newCpu()
	for _, i := range sol.program {
		i.run(c)
	}
}

func init() {
	aoc.RegisterSolution("10", &solution{})
}
