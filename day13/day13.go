// Package day13 solves the day13 puzzle.
package day13

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nerdatmath/aoc2022/aoc"
)

type listOrInt struct {
	l []listOrInt
	i int
}

func parseListOrInt(s []byte, li *listOrInt) ([]byte, error) {
	if len(s) == 0 {
		return nil, errors.New("unexpected end")
	}
	if s[0] == '[' {
		s = s[1:]
		li.l = []listOrInt{}
		if len(s) == 0 {
			return nil, errors.New("unexpected end")
		}
		if s[0] == ']' {
			s = s[1:]
			return s, nil
		}
		var item listOrInt
		s, err := parseListOrInt(s, &item)
		if err != nil {
			return nil, err
		}
		li.l = append(li.l, item)
		if len(s) == 0 {
			return nil, errors.New("unexpected end")
		}
		for s[0] == ',' {
			var item listOrInt
			s = s[1:]
			s, err = parseListOrInt(s, &item)
			if err != nil {
				return nil, err
			}
			li.l = append(li.l, item)
		}
		if len(s) == 0 {
			return nil, errors.New("unexpected end")
		}
		if s[0] != ']' {
			return nil, fmt.Errorf("unexpected character %c", s[0])
		}
		s = s[1:]
		return s, nil
	}
	e := bytes.IndexAny(s, ",]")
	if e == -1 {
		e = len(s)
	}
	i, err := strconv.Atoi(string(s[:e]))
	if err != nil {
		return nil, err
	}
	li.i = i
	return s[e:], nil
}

func (li listOrInt) String() string {
	if li.l != nil {
		out := []string{}
		for _, x := range li.l {
			out = append(out, x.String())
		}
		return fmt.Sprintf("[%s]", strings.Join(out, ","))
	}
	return strconv.Itoa(li.i)
}

func cmpInt(l, r int) int {
	if l < r {
		return -1
	} else if l > r {
		return 1
	}
	return 0
}

func cmpList(l, r []listOrInt) int {
	for i := 0; ; i++ {
		a, b := i == len(l), i == len(r)
		if a && !b {
			return -1
		}
		if !a && b {
			return 1
		}
		if a && b {
			return 0
		}
		c := cmp(l[i], r[i])
		if c != 0 {
			return c
		}
	}
}

func cmp(l, r listOrInt) int {
	if l.l == nil && r.l == nil {
		return cmpInt(l.i, r.i)
	}
	ll, rl := l.l, r.l
	if ll == nil {
		ll = []listOrInt{l}
	}
	if rl == nil {
		rl = []listOrInt{r}
	}
	return cmpList(ll, rl)
}

type pair struct {
	l, r listOrInt
}

func parsePair(s []byte) (pair, error) {
	lines := bytes.Split(s, []byte{'\n'})
	if len(lines) != 2 {
		return pair{}, errors.New("pair is not 2 lines")
	}
	p := pair{}
	s, err := parseListOrInt(lines[0], &p.l)
	if err != nil {
		return pair{}, err
	}
	if len(s) != 0 {
		return pair{}, fmt.Errorf("extra data %q", string(s))
	}
	s, err = parseListOrInt(lines[1], &p.r)
	if err != nil {
		return pair{}, err
	}
	if len(s) != 0 {
		return pair{}, fmt.Errorf("extra data %q", string(s))
	}
	return p, nil
}

func (p pair) rightOrder() bool {
	return cmp(p.l, p.r) < 0
}

type solution struct {
	pairs []pair
}

func (sol *solution) Parse(s []byte) error {
	pairs, err := aoc.ParseDelimited(s, parsePair, []byte{'\n', '\n'})
	if err != nil {
		return err
	}
	sol.pairs = pairs
	return nil
}

func (sol solution) Part1() {
	sum := 0
	for i, p := range sol.pairs {
		if p.rightOrder() {
			sum += i + 1
		}
	}
	fmt.Println("Part 1", sum)
}

func (sol solution) Part2() {
	a := listOrInt{l: []listOrInt{{l: []listOrInt{{i: 2}}}}}
	b := listOrInt{l: []listOrInt{{l: []listOrInt{{i: 6}}}}}
	lta, ltb := 1, 2
	for _, p := range sol.pairs {
		for _, li := range []listOrInt{p.l, p.r} {
			if cmp(li, b) < 0 {
				ltb++
				if cmp(li, a) < 0 {
					lta++
				}
			}
		}
	}
	fmt.Println("Part 2", lta*ltb)
}

func init() {
	aoc.RegisterSolution("13", &solution{})
}
