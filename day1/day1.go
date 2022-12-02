// Package day1 solves the day1 puzzle.
package day1

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"github.com/nerdatmath/aoc2022/io"
)

type elf struct {
	food []int
}

func (e elf) calories() int {
	sum := 0
	for _, c := range e.food {
		sum += c
	}
	return sum
}

func parseElves(b []byte) ([]elf, error) {
	groups := bytes.Split(b, []byte{'\n', '\n'})
	elves := []elf(nil)
	for _, g := range groups {
		e := elf{}
		for _, line := range bytes.Split(g, []byte{'\n'}) {
			cals, err := strconv.Atoi(string(line))
			if err != nil {
				return nil, err
			}
			e.food = append(e.food, cals)
		}
		elves = append(elves, e)
	}
	return elves, nil
}

func input() ([]elf, error) {
	b, err := io.OpenAndReadAll("day1/input.txt")
	if err != nil {
		return nil, err
	}
	return parseElves(b)
}

func Part1() error {
	elves, err := input()
	if err != nil {
		return err
	}
	max := 0
	for _, e := range elves {
		c := e.calories()
		if c > max {
			max = c
		}
	}
	fmt.Println("Part 1", max)
	return nil
}

func Part2() error {
	elves, err := input()
	if err != nil {
		return err
	}
	calories := []int{}
	for _, elf := range elves {
		calories = append(calories, elf.calories())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println("Part 2", calories[0]+calories[1]+calories[2])
	return nil
}
