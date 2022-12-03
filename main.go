// Binary aoc2022 solves the Advent of Code 2022 puzzles.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nerdatmath/aoc2022/day1"
	"github.com/nerdatmath/aoc2022/day2"
	"github.com/nerdatmath/aoc2022/day3"
	"github.com/nerdatmath/aoc2022/day4"
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
	case "3":
		if err := day3.Part1(); err != nil {
			log.Fatalln(err)
		}
		if err := day3.Part2(); err != nil {
			log.Fatalln(err)
		}
	case "4":
		if err := day4.Part1(); err != nil {
			log.Fatalln(err)
		}
		if err := day4.Part2(); err != nil {
			log.Fatalln(err)
		}
	// case "X":
	// 	if err := dayX.Part1(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	if err := dayX.Part2(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	default:
		log.Fatalf("invalid day number %q", os.Args[1])
	}
}
