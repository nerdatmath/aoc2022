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

// This seems like a traveling salesman problem.
// We want to visit each node (valve with rate > 0) via the shortest route
// and visit them in the order that maximizes the total metric.

// let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)
// for each edge (u, v) do
//     dist[u][v] ← w(u, v)  // The weight of the edge (u, v)
// for each vertex v do
//     dist[v][v] ← 0
// for k from 1 to |V|
//     for i from 1 to |V|
//         for j from 1 to |V|
//             if dist[i][j] > dist[i][k] + dist[k][j]
//                 dist[i][j] ← dist[i][k] + dist[k][j]
//             end if

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

func best(actors []actor, targets []valve, paths map[[2]string]int) int {
	max := 0
	for i, t := range targets {
		timeLeft := actors[0].time - (paths[[2]string{actors[0].pos, t.name}] + 1)
		if timeLeft <= 0 {
			continue
		}
		actors := append(slices.Clone(actors[1:]), actor{pos: t.name, time: timeLeft})
		slices.SortFunc(actors, func(a, b actor) bool { return a.time > b.time })
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
