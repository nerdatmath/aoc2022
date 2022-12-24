// Package day21 solves the day21 puzzle.
package day21

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/nerdatmath/aoc2022/aoc"
	"golang.org/x/exp/maps"
)

type monkeyFunc func(string) *poly

// A poly represents an expression (a*x + b), where x is a free variable.
type poly struct {
	a, b big.Rat
}

func (p *poly) isConst() bool {
	return p.a.Sign() == 0
}

func (p *poly) String() string {
	if p.isConst() {
		return p.b.RatString()
	}
	return fmt.Sprintf("%s*x + %s", &p.a, &p.b)
}

func (p *poly) eval(monkeyFunc) *poly {
	return p
}

func (p *poly) Add(a, b *poly) *poly {
	p.a.Add(&a.a, &b.a)
	p.b.Add(&a.b, &b.b)
	return p
}

func (p *poly) Neg(a *poly) *poly {
	p.a.Neg(&a.a)
	p.b.Set(&a.b)
	return p
}

func (p *poly) Sub(a, b *poly) *poly {
	p.a.Sub(&a.a, &b.a)
	p.b.Sub(&a.b, &b.b)
	return p
}

func (p *poly) Mul(a, b *poly) *poly {
	if !(a.isConst() || b.isConst()) {
		panic("quadratic!")
	}
	p.a.Add(
		(&big.Rat{}).Mul(&a.a, &b.b),
		(&big.Rat{}).Mul(&a.b, &b.a))
	p.b.Mul(&a.b, &b.b)
	return p
}

func (p *poly) Quo(a, b *poly) *poly {
	if !b.isConst() {
		panic("non constant divisor")
	}
	divisor := &b.b
	p.a.Quo(&a.a, divisor)
	p.b.Quo(&a.b, divisor)
	return p
}

func (p *poly) solve(r *big.Rat) *big.Rat {
	// x = -b / a
	r.Inv(&p.a)
	r.Mul(&p.b, r)
	r.Neg(r)
	return r
}

type monkey interface {
	fmt.Stringer
	eval(monkeyFunc) *poly
}

type num string

func (n num) eval(monkeyFunc) *poly {
	p := poly{}
	if _, err := fmt.Sscan(string(n), &p.b); err != nil {
		panic(err)
	}
	return &p
}

func (n num) String() string {
	return string(n)
}

type exp struct {
	a, b string
	op   string
}

func (e exp) String() string {
	return fmt.Sprint(e.a, e.op, e.b)
}

func (e exp) eval(m monkeyFunc) *poly {
	a := m(e.a)
	b := m(e.b)
	p := poly{}
	switch e.op {
	case "+":
		return p.Add(a, b)
	case "-":
		return p.Sub(a, b)
	case "*":
		return p.Mul(a, b)
	case "/":
		return p.Quo(a, b)
	}
	panic("operation not recognized")
}

func parseMonkey(s []byte) (monkey, error) {
	words := strings.Split(string(s), " ")
	switch len(words) {
	case 1:
		return num(words[0]), nil
	case 3:
		return exp{a: words[0], b: words[2], op: words[1]}, nil
	}
	return nil, errors.New("incorrect word count")
}

func parseLine(s []byte) (string, monkey, error) {
	i := bytes.Index(s, []byte(": "))
	m, err := parseMonkey(s[i+2:])
	return string(s[:i]), m, err
}

type solution struct {
	monkeys map[string]monkey
}

func (sol *solution) Parse(s []byte) error {
	sol.monkeys = map[string]monkey{}
	for _, l := range bytes.Split(s, []byte{'\n'}) {
		name, monkey, err := parseLine(l)
		if err != nil {
			return err
		}
		sol.monkeys[name] = monkey
	}
	return nil
}

func memoize[T any](f func(string) T) func(string) T {
	cache := map[string]T{}
	return func(s string) T {
		if v, ok := cache[s]; ok {
			return v
		}
		v := f(s)
		cache[s] = v
		return v
	}
}

func (sol solution) Part1() {
	monkeys := maps.Clone(sol.monkeys)
	var eval monkeyFunc
	eval = memoize(func(m string) *poly {
		monkey := monkeys[m]
		// fmt.Printf("Evaluating %s: %s...\n", m, monkey)
		v := monkey.eval(eval)
		// fmt.Printf("%s = %s = %s\n", m, monkey, v)
		return v
	})
	fmt.Println("Part 1", eval("root"))
}

func (sol solution) Part2() {
	monkeys := maps.Clone(sol.monkeys)
	human := poly{}
	human.a.SetInt64(1)
	monkeys["humn"] = &human
	root := monkeys["root"].(exp)
	root.op = "-"
	monkeys["root"] = root
	var eval monkeyFunc
	eval = memoize(func(m string) *poly {
		return monkeys[m].eval(eval)
	})
	fmt.Println("Part 2", eval("root").solve(&big.Rat{}).RatString())
}

func init() {
	aoc.RegisterSolution("21", &solution{})
}
