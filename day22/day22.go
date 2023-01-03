// Package day22 solves the day22 puzzle.
package day22

import (
	"bytes"
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/nerdatmath/aoc2022/aoc"
)

const cubeSize = 4

type direction int

const (
	up direction = iota
	down
	left
	right
	in
	out
)

type side int

const (
	top side = iota
	bottom
	leftMost
	rightMost
	front
	back
)

type point struct {
	side    side
	x, y, z int
}

func awayFrom(b side) direction {
	switch b {
	case front:
		return out
	case back:
		return in
	case rightMost:
		return left
	case leftMost:
		return right
	case bottom:
		return up
	case top:
		return down
	default:
		panic("bad box")
	}
}

func next(p point, d direction) (point, direction) {
	changeSide := func(s side) (point, direction) {
		d = awayFrom(p.side)
		p.side = s
		return p, d
	}
	switch d {
	case left:
		if p.x == 0 {
			return changeSide(leftMost)
		}
		p.x--
	case right:
		if p.x == cubeSize-1 {
			return changeSide(rightMost)
		}
		p.x++
	case up:
		if p.y == 0 {
			return changeSide(top)
		}
		p.y--
	case down:
		if p.y == cubeSize-1 {
			return changeSide(bottom)
		}
		p.y++
	case in:
		if p.z == 0 {
			return changeSide(front)
		}
		p.z--
	case out:
		if p.z == cubeSize-1 {
			return changeSide(back)
		}
		p.z++
	default:
		panic("bad direction")
	}
	return p, d
}

func reverse(d direction) direction {
	switch d {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	case in:
		return out
	case out:
		return in
	}
	panic("invalid direction")
}

func turnRight(s side, d direction) direction {
	switch s {
	case front:
		switch d {
		case left:
			return up
		case up:
			return right
		case right:
			return down
		case down:
			return left
		}
	case back:
		return reverse(turnRight(front, d))
	case top:
		switch d {
		case left:
			return out
		case out:
			return right
		case right:
			return in
		case in:
			return left
		}
	case bottom:
		return reverse(turnRight(top, d))
	case rightMost:
		switch d {
		case up:
			return out
		case out:
			return down
		case down:
			return in
		case in:
			return up
		}
	case leftMost:
		return reverse(turnRight(rightMost, d))
	}
	panic("invalid side / direction combination")
}

func turnLeft(s side, d direction) direction {
	return reverse(turnRight(s, d))
}

type self struct {
	pos   point
	d     direction
	walls mapset.Set[point]
}

func newSelf(walls mapset.Set[point]) *self {
	return &self{
		d:     right,
		walls: walls,
	}
}

func (s *self) forward(n int) {
	for ; n > 0; n-- {
		pos, d := next(s.pos, s.d)
		fmt.Println("attempting to move to", pos, d)
		if s.walls.Contains(pos) {
			fmt.Println("wall encountered")
			break
		}
		fmt.Println("moving to", pos)
		s.pos = pos
	}
}

func (s *self) turn(dir byte) {
	switch dir {
	case 'L':
		fmt.Println("turn left")
		s.d = turnLeft(s.pos.side, s.d)
	case 'R':
		fmt.Println("turn right")
		s.d = turnRight(s.pos.side, s.d)
	}
}

func (s *self) mapLocationAndFacing() (x, y, facing int) {
	panic("notimplemented")
}

func (s *self) password() int {
	x, y, facing := s.mapLocationAndFacing()
	return 1000*y + 4*x + facing
}

func parseMap(s []byte) ([]side, mapset.Set[point], error) {
	boxes := []side{}
	walls := mapset.NewSet[point]()
	b := side{}
	for y, line := range bytes.Split(s, []byte{'\n'}) {
		left := bytes.IndexAny(line, ".#")
		right := len(line) - 1
		if y == 0 {
			b.left = left
			b.right = right
		}
		if left != b.left || right != b.right {
			boxes = append(boxes, b)
			b = side{left: left, right: right, top: y}
		}
		b.bottom = y
		for i, b := range line[left:] {
			if b == '#' {
				walls.Add(point{left + i, y})
			}
		}
	}
	boxes = append(boxes, b)
	return boxes, walls, nil
}

type instruction func(*self)

func parseInstructions(s []byte) ([]instruction, error) {
	inst := []instruction{}
	for len(s) != 0 {
		switch s[0] {
		case 'L', 'R':
			dir := s[0]
			inst = append(inst, func(me *self) { me.turn(dir) })
			s = s[1:]
			continue
		}
		i := bytes.IndexAny(s, "LR")
		if i == -1 {
			i = len(s)
		}
		n, err := strconv.Atoi(string(s[:i]))
		if err != nil {
			return nil, err
		}
		inst = append(inst, func(s *self) { s.forward(n) })
		s = s[i:]
	}
	return inst, nil
}

type solution struct {
	boxes        []side
	walls        mapset.Set[point]
	instructions []instruction
}

func (sol *solution) Parse(s []byte) error {
	m, i, _ := bytes.Cut(s, []byte("\n\n"))
	boxes, walls, err := parseMap(m)
	if err != nil {
		return err
	}
	instructions, err := parseInstructions(i)
	if err != nil {
		return err
	}
	sol.boxes = boxes
	sol.walls = walls
	sol.instructions = instructions
	return nil
}

func (sol solution) Part1() {
	fmt.Println("starting part 1")
	fmt.Println("walls", sol.walls)
	s := newSelf(sol.boxes, sol.walls)
	for _, f := range sol.instructions {
		fmt.Println(s.pos, s.o)
		f(s)
	}
	fmt.Println("Part 1", s.password())
}

func (sol solution) Part2() {
}

func init() {
	aoc.RegisterSolution("22", &solution{})
}
