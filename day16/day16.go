// Package day16 solves the day16 puzzle.
package day16

import (
	"bytes"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
	"golang.org/x/exp/slices"
)

type valve struct {
	name      string
	rate      int
	connected []string
}

func parseValve(s []byte) (valve, error) {
	var name string
	var rate int
	p := bytes.Index(s, []byte("; "))
	_, err := fmt.Sscanf(string(s[:p]), "Valve %s has flow rate=%d", &name, &rate)
	if err != nil {
		return valve{}, err
	}
	connected := string(s[p+2:])
	connected = strings.TrimPrefix(connected, "tunnels lead to valves ")
	connected = strings.TrimPrefix(connected, "tunnel leads to valve ")
	return valve{
		name:      name,
		rate:      rate,
		connected: strings.Split(connected, ", "),
	}, nil
}

type solution struct {
	valves []valve
}

func (sol *solution) Parse(s []byte) error {
	valves, err := aoc.ParseLines(s, parseValve)
	if err != nil {
		return err
	}
	sol.valves = valves
	return err
}

func shortestPaths(nodes []string, edges mapset.Set[[2]string]) map[[2]string]int {
	g := map[[2]string]int{}
	for e := range edges.Iter() {
		g[e] = 1
	}
	for _, n := range nodes {
		g[[2]string{n, n}] = 0
	}
	for _, k := range nodes {
		for _, i := range nodes {
			for _, j := range nodes {
				ik, ok := g[[2]string{i, k}]
				if !ok {
					continue
				}
				jk, ok := g[[2]string{j, k}]
				if !ok {
					continue
				}
				ij, ok := g[[2]string{i, j}]
				if !ok || ij > ik+jk {
					g[[2]string{i, j}] = ik + jk
				}
			}
		}
	}
	return g
}

type actor struct {
	pos  string
	time int
}

type pq []actor

func (p *pq) pop() actor {
	a := (*p)[0]
	*p = (*p)[1:]
	return a
}

func (p *pq) push(a actor) {
	*p = append(*p, a)
	slices.SortFunc(*p, func(a, b actor) bool { return a.time > b.time })
}

func (p *pq) clone() pq {
	return pq(slices.Clone(*p))
}

func best(actors pq, targets []valve, paths map[[2]string]int) int {
	a := actors.pop()
	max := 0
	for i, t := range targets {
		timeLeft := a.time - (paths[[2]string{a.pos, t.name}] + 1)
		if timeLeft <= 0 {
			continue
		}
		actors := actors.clone()
		actors.push(actor{pos: t.name, time: timeLeft})
		remaining := slices.Delete(slices.Clone(targets), i, i+1)
		score := timeLeft*t.rate + best(actors, remaining, paths)
		if score > max {
			max = score
		}
	}
	return max
}

func (sol solution) targets() []valve {
	targets := []valve{}
	for _, v1 := range sol.valves {
		if v1.rate > 0 {
			targets = append(targets, v1)
		}
	}
	return targets
}

func (sol solution) paths() map[[2]string]int {
	edges := mapset.NewSet[[2]string]()
	nodes := []string{}
	for _, v1 := range sol.valves {
		nodes = append(nodes, v1.name)
		for _, v2 := range v1.connected {
			edges.Add([2]string{v1.name, v2})
		}
	}
	return shortestPaths(nodes, edges)
}

func (sol solution) Part1() {
	actors := []actor{{pos: "AA", time: 30}}
	score := best(actors, sol.targets(), sol.paths())
	fmt.Println("Part 1", score)
}

func (sol solution) Part2() {
	actors := []actor{{pos: "AA", time: 26}, {pos: "AA", time: 26}}
	score := best(actors, sol.targets(), sol.paths())
	fmt.Println("Part 2", score)
}

func init() {
	aoc.RegisterSolution("16", &solution{})
}
