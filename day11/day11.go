// Package day11 solves the day11 puzzle.
package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nerdatmath/aoc2022/aoc"
	"golang.org/x/exp/slices"
)

type item struct {
	worry int
}

func parseItem(s []byte) (item, error) {
	worry, err := strconv.Atoi(string(s))
	if err != nil {
		return item{}, err
	}
	return item{worry: worry}, nil
}

type operation func(item) item

func parseOperation(s []byte) (operation, error) {
	words := strings.Split(string(s), " ")
	if len(words) != 3 || words[0] != "old" {
		return nil, fmt.Errorf("unrecognized operation %q", s)
	}
	switch words[1] {
	case "+":
		n, err := strconv.Atoi(words[2])
		if err != nil {
			return nil, err
		}
		return func(i item) item {
			return item{worry: i.worry + n}
		}, nil
	case "*":
		if words[2] == "old" {
			return func(i item) item { return item{worry: i.worry * i.worry} }, nil
		}
		n, err := strconv.Atoi(words[2])
		if err != nil {
			return nil, err
		}
		return func(i item) item {
			return item{worry: i.worry * n}
		}, nil
	default:
		return nil, fmt.Errorf("unrecognized operation %q", s)
	}
}

type test func(item) bool

type monkey struct {
	items     []item
	op        operation
	test      test
	trueCase  int
	falseCase int
}

func parseMonkey(s []byte) (monkey, error) {
	lines := strings.Split(string(s), "\n")
	items, err := aoc.ParseDelimited([]byte(strings.TrimPrefix(lines[1], "  Starting items: ")), parseItem, []byte(", "))
	if err != nil {
		return monkey{}, err
	}
	op, err := parseOperation([]byte(strings.TrimPrefix(lines[2], "  Operation: new = ")))
	if err != nil {
		return monkey{}, err
	}
	testDiv, err := strconv.Atoi(strings.TrimPrefix(lines[3], "  Test: divisible by "))
	if err != nil {
		return monkey{}, err
	}
	trueMonkey, err := strconv.Atoi(strings.TrimPrefix(lines[4], "    If true: throw to monkey "))
	if err != nil {
		return monkey{}, err
	}
	falseMonkey, err := strconv.Atoi(strings.TrimPrefix(lines[5], "    If false: throw to monkey "))
	if err != nil {
		return monkey{}, err
	}
	return monkey{
		items:     items,
		op:        op,
		test:      func(i item) bool { return i.worry%testDiv == 0 },
		trueCase:  trueMonkey,
		falseCase: falseMonkey,
	}, nil
}

func (m *monkey) turn(manageWorry func(int) int, throw func(int, item)) {
	items := m.items
	m.items = nil
	for _, item := range items {
		item = m.op(item)
		item.worry = manageWorry(item.worry)
		if m.test(item) {
			throw(m.trueCase, item)
		} else {
			throw(m.falseCase, item)
		}
	}
}

func (m *monkey) catch(i item) {
	m.items = append(m.items, i)
}

type solution struct {
	monkeys []monkey
}

func (sol *solution) Parse(s []byte) error {
	monkeys, err := aoc.ParseDelimited(s, parseMonkey, []byte{'\n', '\n'})
	sol.monkeys = monkeys
	return err
}

func (sol *solution) throw(m int, i item) {
	sol.monkeys[m].catch(i)
}

func (sol solution) Part1() {
	manageWorry := func(worry int) int {
		return worry / 3
	}
	inspections := make([]int, len(sol.monkeys))
	for i := 0; i < 20; i++ {
		for i := range sol.monkeys {
			inspections[i] += len(sol.monkeys[i].items)
			sol.monkeys[i].turn(manageWorry, sol.throw)
		}
	}
	slices.SortFunc(inspections, func(a, b int) bool { return a > b })
	fmt.Println("Part 1", inspections[0]*inspections[1])
}

func (sol solution) Part2() {
	manageWorry := func(worry int) int {
		return worry % (2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23)
	}
	inspections := make([]int, len(sol.monkeys))
	for i := 0; i < 10000; i++ {
		for i := range sol.monkeys {
			inspections[i] += len(sol.monkeys[i].items)
			sol.monkeys[i].turn(manageWorry, sol.throw)
		}
	}
	slices.SortFunc(inspections, func(a, b int) bool { return a > b })
	fmt.Println("Part 2", inspections[0]*inspections[1])
}

func init() {
	aoc.RegisterSolution("11", &solution{})
}
