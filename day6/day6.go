// Package day6 solves the day6 puzzle.
package day6

import (
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type solution struct {
	sig []byte
}

func (sol *solution) Parse(s []byte) error {
	sol.sig = s
	return nil
}

func marker(sig []byte, l int) int {
	seen := map[byte]int{}
	for i, b := range sig {
		seen[b]++
		if i >= l {
			old := sig[i-l]
			seen[old]--
			if seen[old] == 0 {
				delete(seen, old)
			}
		}
		if len(seen) == l {
			return i
		}
	}
	return len(sig)
}

func (sol solution) Part1() {
	m := marker(sol.sig, 4)
	if m >= len(sol.sig) {
		panic("marker not found")
	}
	fmt.Println("Part 1", m+1)
}

func (sol solution) Part2() {
	m := marker(sol.sig, 14)
	if m >= len(sol.sig) {
		panic("marker not found")
	}
	fmt.Println("Part 2", m+1)
}

func init() {
	aoc.RegisterSolution("6", &solution{})
}
