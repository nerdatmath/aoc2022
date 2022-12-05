// Package day5 solves the day5 puzzle.
package day5

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type item byte

type stackName byte

type stacks map[stackName][]item

func parseStacks(s []byte) (stacks, error) {
	lines := bytes.Split(s, []byte("\n"))
	stacks := stacks{}
	nameLine := lines[len(lines)-1]
	for j := 1; j < len(nameLine); j += 4 {
		name := stackName(nameLine[j])
		for i := len(lines) - 2; i >= 0; i-- {
			b := lines[i][j]
			if b == ' ' {
				break
			}
			stacks.push(name, item(b))
		}
	}
	return stacks, nil
}

func (s stacks) push(n stackName, i ...item) {
	s[n] = append(s[n], i...)
}

func (s stacks) pop(n stackName, c int) []item {
	l := len(s[n]) - c
	i := (s[n])[l:]
	s[n] = (s[n])[:l]
	return i
}

func (s stacks) top(n stackName) item {
	l := len(s[n]) - 1
	return (s[n])[l]
}

func (s stacks) copy() stacks {
	out := stacks{}
	for n, st := range s {
		out[n] = append([]item{}, st...)
	}
	return out
}

func (s stacks) move(from, to stackName, c int) {
	s.push(to, s.pop(from, c)...)
}

type move struct {
	count    int
	from, to stackName
}

func parseMove(s []byte) (move, error) {
	var m move
	_, err := fmt.Sscanf(string(s), "move %d from %c to %c", &m.count, &m.from, &m.to)
	return m, err
}

type solution struct {
	stacks stacks
	moves  []move
}

func (sol *solution) Parse(s []byte) error {
	parts := bytes.Split(s, []byte("\n\n"))
	if len(parts) != 2 {
		return errors.New("parse error")
	}
	stacks, err := parseStacks(parts[0])
	if err != nil {
		return err
	}
	moves, err := aoc.ParseLines(parts[1], parseMove)
	if err != nil {
		return err
	}
	sol.stacks = stacks
	sol.moves = moves
	return nil
}

func tops(s stacks) string {
	word := ""
	for n := stackName('1'); n <= '9'; n++ {
		word = word + string(s.top(n))
	}
	return word
}

func (sol solution) Part1() {
	s := sol.stacks.copy()
	for _, m := range sol.moves {
		for i := 0; i < m.count; i++ {
			s.move(m.from, m.to, 1)
		}
	}
	fmt.Println("Part 1", tops(s))
}

func (sol solution) Part2() {
	s := sol.stacks.copy()
	for _, m := range sol.moves {
		s.move(m.from, m.to, m.count)
	}
	fmt.Println("Part 2", tops(s))
}

func init() {
	aoc.RegisterSolution("5", &solution{})
}
