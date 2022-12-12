// Package day8 solves the day8 puzzle.
package day8

import (
	"bytes"
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type solution struct {
	h [][]byte
}

func (sol *solution) Parse(s []byte) error {
	sol.h = bytes.Split(s, []byte{'\n'})
	return nil
}

func (sol solution) Part1() {
	w, h := len(sol.h[0]), len(sol.h)
	vis := map[[2]int]struct{}{}
	mark := func(x, y int) { vis[[2]int{x, y}] = struct{}{} }
	mark(0, 0)
	mark(w-1, 0)
	mark(0, h-1)
	mark(w-1, h-1)
	for x := 1; x < w-1; x++ {
		maxPos := 0
		mark(x, maxPos)
		for y := 1; y < h-1; y++ {
			if sol.h[y][x] > sol.h[maxPos][x] {
				maxPos = y
				mark(x, maxPos)
			}
		}
		maxPos = h - 1
		mark(x, maxPos)
		for y := h - 2; y > 0; y-- {
			if sol.h[y][x] > sol.h[maxPos][x] {
				maxPos = y
				mark(x, maxPos)
			}
		}
	}
	for y := 1; y < h-1; y++ {
		maxPos := 0
		mark(maxPos, y)
		for x := 1; x < w-1; x++ {
			if sol.h[y][x] > sol.h[y][maxPos] {
				maxPos = x
				mark(maxPos, y)
			}
		}
		maxPos = w - 1
		mark(maxPos, y)
		for x := w - 2; x > 0; x-- {
			if sol.h[y][x] > sol.h[y][maxPos] {
				maxPos = x
				mark(maxPos, y)
			}
		}
	}
	// for y := range sol.h {
	// 	for x := range sol.h[0] {
	// 		if _, ok := vis[[2]int{x, y}]; ok {
	// 			fmt.Print(string(sol.h[y][x]))
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println("Part 1", len(vis))
}

func (sol solution) Part2() {
	xMax, yMax := len(sol.h[0]), len(sol.h)
	up := map[[2]int]int{}
	dn := map[[2]int]int{}
	lt := map[[2]int]int{}
	rt := map[[2]int]int{}
	for x := 1; x < xMax-1; x++ {
		upBlockedByHeight := [10]int{}
		for y := 0; y < yMax; y++ {
			h := sol.h[y][x] - '0'
			up[[2]int{x, y}] = y - upBlockedByHeight[h]
			for i := byte(0); i <= h; i++ {
				upBlockedByHeight[i] = y
			}
		}
		dnBlockedByHeight := [10]int{}
		for y := yMax - 1; y >= 0; y-- {
			h := sol.h[y][x] - '0'
			dn[[2]int{x, y}] = (yMax - 1) - y - dnBlockedByHeight[h]
			for i := byte(0); i <= h; i++ {
				dnBlockedByHeight[i] = (yMax - 1) - y
			}
		}
	}
	for y := 1; y < yMax-1; y++ {
		ltBlockedByHeight := [10]int{}
		for x := 0; x < xMax; x++ {
			h := sol.h[y][x] - '0'
			lt[[2]int{x, y}] = x - ltBlockedByHeight[h]
			for i := byte(0); i <= h; i++ {
				ltBlockedByHeight[i] = x
			}
		}
		rtBlockedByHeight := [10]int{}
		for x := xMax - 1; x >= 0; x-- {
			h := sol.h[y][x] - '0'
			rt[[2]int{x, y}] = (xMax - 1) - x - rtBlockedByHeight[h]
			for i := byte(0); i <= h; i++ {
				rtBlockedByHeight[i] = (xMax - 1) - x
			}
		}
	}
	max := 0
	for y := 1; y < yMax-1; y++ {
		for x := 1; x < xMax-1; x++ {
			pt := [2]int{x, y}
			score := up[pt] * dn[pt] * lt[pt] * rt[pt]
			if score > max {
				max = score
			}
		}
	}
	fmt.Println("Part 2", max)
}

func init() {
	aoc.RegisterSolution("8", &solution{})
}
