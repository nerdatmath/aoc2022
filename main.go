// Binary aoc2022 solves the Advent of Code 2022 puzzles.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nerdatmath/aoc2022/aoc"
	"github.com/nerdatmath/aoc2022/io"

	_ "github.com/nerdatmath/aoc2022/day1"
	_ "github.com/nerdatmath/aoc2022/day2"
	_ "github.com/nerdatmath/aoc2022/day2b"
	_ "github.com/nerdatmath/aoc2022/day3"
	_ "github.com/nerdatmath/aoc2022/day4"
	_ "github.com/nerdatmath/aoc2022/day5"
	_ "github.com/nerdatmath/aoc2022/day6"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "day number required")
		os.Exit(1)
	}
	day := os.Args[1]
	s, err := aoc.Lookup(day)
	if err != nil {
		log.Fatalln(err)
	}
	p, err := io.OpenAndReadAll(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Parse(p); err != nil {
		log.Fatalln(err)
	}
	s.Part1()
	s.Part2()
}
