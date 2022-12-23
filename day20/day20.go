// Package day20 solves the day20 puzzle.
package day20

import (
	"fmt"
	"strconv"

	"github.com/nerdatmath/aoc2022/aoc"
	"golang.org/x/exp/slices"
)

type solution struct {
	n []int
}

func (sol *solution) Parse(s []byte) error {
	n, err := aoc.ParseLines(s, func(s []byte) (int, error) { return strconv.Atoi(string(s)) })
	sol.n = n
	return err
}

func mod(n, d int) int {
	n = n % d
	if n < 0 {
		n += d
	}
	return n
}

type ring struct {
	l    int
	next []int
	prev []int
}

func (r ring) link(a, b int) {
	r.next[a] = b
	r.prev[b] = a
}

func newRing(l int) ring {
	r := ring{
		l:    l,
		next: make([]int, l),
		prev: make([]int, l),
	}
	for i := 0; i < l; i++ {
		r.link(i, (i+1)%l)
	}
	return r
}

func (r ring) forward(pos, steps int) int {
	steps = mod(steps, r.l)
	for j := 0; j < steps; j++ {
		pos = r.next[pos]
	}
	return pos
}

func mix(r ring, src []int) {
	for i := range src {
		pos := r.forward(i, mod(src[i], r.l-1))
		if pos == i {
			continue
		}
		n := r.next[pos]
		r.link(r.prev[i], r.next[i])
		r.link(pos, i)
		r.link(i, n)
	}
}

func findZero(src []int) int {
	for i := range src {
		if src[i] == 0 {
			return i
		}
	}
	panic("zero not found")
}

func (sol solution) Part1() {
	r := newRing(len(sol.n))
	zeroPos := findZero(sol.n)
	mix(r, sol.n)
	x := sol.n[r.forward(zeroPos, 1000)]
	y := sol.n[r.forward(zeroPos, 2000)]
	z := sol.n[r.forward(zeroPos, 3000)]
	fmt.Println("Part 1", x+y+z)
}

func (sol solution) Part2() {
	n := slices.Clone(sol.n)
	for i := range n {
		n[i] *= 811589153
	}
	r := newRing(len(n))
	zeroPos := findZero(n)
	for i := 0; i < 10; i++ {
		mix(r, n)
	}
	x := n[r.forward(zeroPos, 1000)]
	y := n[r.forward(zeroPos, 2000)]
	z := n[r.forward(zeroPos, 3000)]
	fmt.Println("Part 2", x+y+z)
}

func init() {
	aoc.RegisterSolution("20", &solution{})
}
