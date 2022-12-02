// Binary day1 solves the day 1 puzzle.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nerdatmath/aoc2022/day1"
	"github.com/nerdatmath/aoc2022/day2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "day number required")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1":
		if err := day1.Part1(); err != nil {
			log.Fatalln(err)
		}
		if err := day1.Part2(); err != nil {
			log.Fatalln(err)
		}
	case "2":
		if err := day2.Part1(); err != nil {
			log.Fatalln(err)
		}
		if err := day2.Part2(); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalf("invalid day number %q", os.Args[1])
	}
}
