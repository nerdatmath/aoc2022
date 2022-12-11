// Package day7 solves the day7 puzzle.
package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nerdatmath/aoc2022/aoc"
)

type tree struct {
	parent *tree
	dirs   map[string]*tree
	files  map[string]int
}

func newTree(parent *tree) *tree {
	return &tree{
		parent: parent,
		dirs:   make(map[string]*tree),
		files:  make(map[string]int),
	}
}

type solution struct {
	root *tree
	cwd  *tree
}

func (sol *solution) Parse(s []byte) error {
	sol.root = newTree(nil)
	sol.cwd = sol.root
	if string(s[:2]) != "$ " {
		panic("expected prompt")
	}
	for _, c := range strings.Split(string(s[2:]), "\n$ ") {
		lines := strings.Split(c, "\n")
		words := strings.Split(lines[0], " ")
		cmd, args := words[0], words[1:]
		sol.cmd(cmd, args, lines[1:])
	}
	return nil
}

func (sol *solution) cmd(c string, args []string, input []string) {
	switch c {
	case "cd":
		sol.cd(args, input)
	case "ls":
		sol.ls(args, input)
	default:
		panic("unrecognized command")
	}
}

func (sol *solution) cd(args []string, input []string) {
	if len(args) != 1 || len(input) != 0 {
		panic("invalid cd command")
	}
	arg := args[0]
	switch arg {
	case "..":
		sol.cwd = sol.cwd.parent
	case "/":
		sol.cwd = sol.root
	default:
		sol.cwd = sol.cwd.dirs[arg]
	}
}

func (sol *solution) ls(args []string, input []string) {
	if len(args) != 0 {
		panic("invalid ls command")
	}
	for _, line := range input {
		words := strings.Split(line, " ")
		if len(words) != 2 {
			panic("invalid ls input")
		}
		name := words[1]
		switch words[0] {
		case "dir":
			sol.cwd.dirs[name] = newTree(sol.cwd)
		default:
			l, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}
			sol.cwd.files[name] = l
		}
	}
}

func size(d *tree, dirs map[*tree]int) int {
	s := 0
	for _, sd := range d.dirs {
		s += size(sd, dirs)
	}
	for _, fs := range d.files {
		s += fs
	}
	dirs[d] = s
	return s
}

func (sol solution) Part1() {
	dirs := map[*tree]int{}
	_ = size(sol.root, dirs)
	tot := 0
	for _, s := range dirs {
		if s <= 100000 {
			tot += s
		}
	}
	fmt.Println("Part 1", tot)
}

func (sol solution) Part2() {
	dirs := map[*tree]int{}
	tot := size(sol.root, dirs)
	needed := tot - 40000000
	min := tot
	for _, s := range dirs {
		if s < needed {
			continue
		}
		if s < min {
			min = s
		}
	}
	fmt.Println("Part 2", min)
}

func init() {
	aoc.RegisterSolution("7", &solution{})
}
