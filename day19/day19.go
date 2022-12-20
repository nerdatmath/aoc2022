// Package day19 solves the day19 puzzle.
package day19

import (
	"fmt"

	"github.com/nerdatmath/aoc2022/aoc"
)

type blueprint struct {
	id   int
	or   int
	cr   int
	obro int
	obrc int
	gro  int
	grob int
}

func parseBlueprint(s []byte) (blueprint, error) {
	bp := blueprint{}
	_, err := fmt.Sscanf(string(s), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&bp.id, &bp.or, &bp.cr, &bp.obro, &bp.obrc, &bp.gro, &bp.grob)
	return bp, err
}

type solution struct {
	bps []blueprint
}

func (sol *solution) Parse(s []byte) error {
	bps, err := aoc.ParseLines(s, parseBlueprint)
	sol.bps = bps
	return err
}

func (sol solution) Part1() {
	for _, bp := range sol.bps {
		fmt.Println(bp)
	}
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("19", &solution{})
}
