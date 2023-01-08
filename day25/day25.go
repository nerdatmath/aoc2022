// Package day25 solves the day25 puzzle.
package day25

import (
	"bytes"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

// snafu digits are in little-endian order.
type snafu []int8

func addDigits(a, b, c int8) (sum int8, carry int8) {
	sum = a + b + c
	if sum < -2 {
		carry = -1
		sum += 5
	}
	if sum > 2 {
		carry = 1
		sum -= 5
	}
	return sum, carry
}

func add(x, y snafu, c int8) snafu {
	if len(x) == 0 && len(y) == 0 && c == 0 {
		return nil
	}
	a := int8(0)
	if len(x) != 0 {
		a = x[0]
		x = x[1:]
	}
	b := int8(0)
	if len(y) != 0 {
		b = y[0]
		y = y[1:]
	}
	sum, c := addDigits(a, b, c)
	return append(snafu{sum}, add(x, y, c)...)
}

func (s *snafu) Parse(p []byte) error {
	*s = nil
	for i := len(p) - 1; i >= 0; i-- {
		d, ok := map[byte]int8{
			'=': -2,
			'-': -1,
			'0': 0,
			'1': +1,
			'2': +2,
		}[p[i]]
		if !ok {
			return fmt.Errorf("invalid digit %c", p[i])
		}
		*s = append(*s, d)
	}
	return nil
}

func (s snafu) String() string {
	out := []byte(nil)
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, map[int8]byte{
			-2: '=',
			-1: '-',
			0:  '0',
			+1: '1',
			+2: '2',
		}[s[i]])
	}
	return string(out)
}

type solution struct {
	numbers []snafu
}

func (sol *solution) Parse(s []byte) error {
	for _, line := range bytes.Fields(s) {
		var n snafu
		if err := n.Parse(line); err != nil {
			return err
		}
		sol.numbers = append(sol.numbers, n)
	}
	return nil
}

func (sol solution) Part1() {
	sum := snafu{}
	for _, n := range sol.numbers {
		sum = add(sum, n, 0)
	}
	fmt.Println("Part 1", sum)
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("25", &solution{})
}
