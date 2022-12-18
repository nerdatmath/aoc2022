package aoc

import (
	"bytes"
	"math"

	mapset "github.com/deckarep/golang-set/v2"
)

// ParseDelimited runs a parser on each section.
func ParseDelimited[T any](s []byte, parse func([]byte) (T, error), delim []byte) ([]T, error) {
	lines := bytes.Split(s, delim)
	ts := []T(nil)
	for _, line := range lines {
		if t, err := parse(line); err != nil {
			return nil, err
		} else {
			ts = append(ts, t)
		}
	}
	return ts, nil
}

// ParseLines runs a parser on each line.
func ParseLines[T any](s []byte, parse func([]byte) (T, error)) ([]T, error) {
	return ParseDelimited(s, parse, []byte{'\n'})
}

// Map.
func Map[S, T any](f func(S) T, ss []S) []T {
	ts := []T(nil)
	for _, s := range ss {
		ts = append(ts, f(s))
	}
	return ts
}

// Sum sums up the values.
func Sum(xs []int) int {
	return Fold(func(a, b int) int { return a + b }, 0, xs)
}

// Max returns the maximum of the values.
func Max(xs []int) int {
	return Fold(func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}, math.MinInt, xs)
}

// Fold.
func Fold[A, B any](f func(A, B) B, b B, as []A) B {
	for _, a := range as {
		b = f(a, b)
	}
	return b
}

// Filter.
func Filter[A any](f func(A) bool, as []A) []A {
	out := []A(nil)
	for _, a := range as {
		if f(a) {
			out = append(out, a)
		}
	}
	return out
}

// Breadth First Search.
func BFS[T comparable](start []T, edges func(T) []T) {
	active := mapset.NewSet(start...)
	visited := mapset.NewSet[T]()
	for active.Cardinality() != 0 {
		visited = visited.Union(active)
		next := []T{}
		for n := range active.Iter() {
			next = append(next, edges(n)...)
		}
		active = mapset.NewSet(next...).Difference(visited)
	}
}
