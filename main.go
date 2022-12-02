// Binary day1 solves the day 1 puzzle.
package main

import (
	"log"

	"github.com/nerdatmath/aoc2022/day1"
	"github.com/nerdatmath/aoc2022/day2"
)

func main() {
	if err := day1.Part1(); err != nil {
		log.Fatalln(err)
	}
	if err := day1.Part2(); err != nil {
		log.Fatalln(err)
	}
	if err := day2.Part1(); err != nil {
		log.Fatalln(err)
	}
	if err := day2.Part2(); err != nil {
		log.Fatalln(err)
	}
}
