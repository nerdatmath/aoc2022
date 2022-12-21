// Package day19 solves the day19 puzzle.
package day19

import (
	"fmt"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

type material int

const (
	ore material = iota
	clay
	obsidian
	geode
)

type blueprint struct {
	id    int
	costs [4][4]int // robot -> material -> cost
}

func parseBlueprint(s []byte) (blueprint, error) {
	var (
		id                      int
		oreRobotCostsOre        int
		clayRobotCostsOre       int
		obsidianRobotCostsOre   int
		obsidianRobotCostsClay  int
		geodeRobotCostsOre      int
		geodeRobotCostsObsidian int
	)
	_, err := fmt.Sscanf(string(s), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&id,
		&oreRobotCostsOre,
		&clayRobotCostsOre,
		&obsidianRobotCostsOre,
		&obsidianRobotCostsClay,
		&geodeRobotCostsOre,
		&geodeRobotCostsObsidian,
	)
	return blueprint{
		id: id,
		costs: [4][4]int{
			ore:      {ore: oreRobotCostsOre},
			clay:     {ore: clayRobotCostsOre},
			obsidian: {ore: obsidianRobotCostsOre, clay: obsidianRobotCostsClay},
			geode:    {ore: geodeRobotCostsOre, obsidian: geodeRobotCostsObsidian},
		},
	}, err
}

type solution struct {
	bps []blueprint
}

func (sol *solution) Parse(s []byte) error {
	bps, err := aoc.ParseLines(s, parseBlueprint)
	sol.bps = bps
	return err
}

type state struct {
	materials [4]int
	robots    [4]int
}

func (s state) canAfford(cost [4]int) (state, bool) {
	for m, c := range cost {
		s.materials[m] -= c
		if s.materials[m] < 0 {
			return state{}, false
		}
	}
	return s, true
}

func (s state) run() state {
	for m, count := range s.robots {
		s.materials[m] += count
	}
	return s
}

func (b blueprint) edges(s state) []state {
	out := []state{}
	if s, ok := s.canAfford(b.costs[geode]); ok {
		s = s.run()
		s.robots[geode]++
		out = append(out, s)
	}
	if s, ok := s.canAfford(b.costs[obsidian]); ok && s.robots[obsidian] < b.costs[geode][obsidian] {
		s = s.run()
		s.robots[obsidian]++
		out = append(out, s)
	}
	if s, ok := s.canAfford(b.costs[clay]); ok && s.robots[clay] < b.costs[obsidian][clay] {
		s = s.run()
		s.robots[clay]++
		out = append(out, s)
	}
	if s, ok := s.canAfford(b.costs[ore]); ok && s.robots[ore] < 4 {
		s = s.run()
		s.robots[ore]++
		out = append(out, s)
	}
	if len(out) == 0 || (s.materials[ore] < 10 && s.robots[ore] < 4) {
		out = append(out, s.run())
	}
	return out
}

func less(a, b state) bool {
	if a.materials[geode]+2 < b.materials[geode] {
		return true
	}
	for m := ore; m <= geode; m++ {
		if a.materials[m] > b.materials[m] {
			return false
		}
		if a.robots[m] > b.robots[m] {
			return false
		}
	}
	return true
}

func addState(s state, states mapset.Set[state]) {
	for _, t := range states.ToSlice() {
		if less(t, s) {
			states.Remove(t)
		}
		if less(s, t) {
			return
		}
	}
	states.Add(s)
}

func maxGeodes(bp blueprint, minutes int) int {
	states := mapset.NewSet(state{
		robots: [4]int{
			ore: 1,
		},
	})
	for i := 0; i < minutes-1; i++ {
		newStates := mapset.NewSet[state]()
		for s := range states.Iter() {
			for _, s := range bp.edges(s) {
				addState(s, newStates)
			}
		}
		states = newStates
	}
	max := 0
	for s := range states.Iter() {
		s = s.run()
		if s.materials[geode] > max {
			max = s.materials[geode]
		}
	}
	return max
}

func (sol solution) Part1() {
	wg := sync.WaitGroup{}
	g := make([]int, len(sol.bps))
	for i, bp := range sol.bps {
		i := i
		bp := bp
		wg.Add(1)
		go func() {
			g[i] = maxGeodes(bp, 24)
			wg.Done()
		}()
	}
	wg.Wait()
	sum := 0
	for i, bp := range sol.bps {
		sum += bp.id * g[i]
	}
	fmt.Println("Part 1", sum)
}

func (sol solution) Part2() {
	bps := sol.bps[:3]
	wg := sync.WaitGroup{}
	g := make([]int, len(bps))
	for i, bp := range bps {
		i := i
		bp := bp
		wg.Add(1)
		go func() {
			g[i] = maxGeodes(bp, 32)
			wg.Done()
		}()
	}
	wg.Wait()
	product := 1
	for i := range bps {
		product *= g[i]
	}
	fmt.Println("Part 2", product)
}

func init() {
	aoc.RegisterSolution("19", &solution{})
}
