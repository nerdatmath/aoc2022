// Package day1 solves the day1 puzzle.
package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func input() ([][]int, error) {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	elves := [][]int(nil)
	calories := []int(nil)
	for s.Scan() {
		if s.Text() == "" {
			elves = append(elves, calories)
			calories = nil
			continue
		}
		cal, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		calories = append(calories, cal)
	}
	if len(calories) > 0 {
		elves = append(elves, calories)
		calories = nil
	}
	return elves, nil
}

func Part1() error {
	elves, err := input()
	if err != nil {
		return err
	}
	max := 0
	for _, elf := range elves {
		cals := 0
		for _, cal := range elf {
			cals += cal
		}
		if cals > max {
			max = cals
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
	sums := []int{}
	for _, elf := range elves {
		cals := 0
		for _, cal := range elf {
			cals += cal
		}
		sums = append(sums, cals)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Println("Part 2", sums[0]+sums[1]+sums[2])
	return nil
}
