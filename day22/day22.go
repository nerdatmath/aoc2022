// Package day22 solves the day22 puzzle.
package day22

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/nerdatmath/aoc2022/aoc"
	"github.com/nerdatmath/aoc2022/day22/cubes"
)

const cubeSize = 50

const (
	wall  byte = '#'
	blank byte = ' '
	open  byte = '.'
)

type position struct {
	row, col int
}

func (p position) String() string {
	r, row, col := p.region()
	return fmt.Sprintf("%s(%d,%d)", r, row, col)
}

type direction struct {
	row, col int
}

func (d direction) String() string {
	switch d {
	case direction{row: -1, col: 0}:
		return "^"
	case direction{row: 1, col: 0}:
		return "v"
	case direction{row: 0, col: -1}:
		return "<"
	case direction{row: 0, col: 1}:
		return ">"
	default:
		panic("unexpected direction")
	}
}

type self struct {
	pos position
	d   direction
}

type shape interface {
	start() position
	at(position) byte
	step(position, direction) (position, direction)
}

type field [][]byte

func (f field) start() position {
	p := position{}
	for ; f.at(p) == blank; p.col++ {
	}
	return p
}

func (f field) at(p position) byte {
	if p.row < 0 || p.row >= len(f) || p.col < 0 || p.col >= len(f[p.row]) {
		return blank
	}
	return f[p.row][p.col]
}

func (f field) step(p position, d direction) (position, direction) {
	p = p.add(d)
	for f.at(p) == blank {
		if p.row >= len(f) {
			p.row = 0
		} else if p.row < 0 {
			p.row = len(f) - 1
		} else if p.col >= len(f[p.row]) {
			p.col = 0
		} else if p.col < 0 {
			p.col = len(f[p.row]) - 1
		} else {
			p = p.add(d)
		}
	}
	return p, d
}

func (p position) add(d direction) position {
	return position{row: p.row + d.row, col: p.col + d.col}
}

func (s *self) forward(n int, f shape) {
	for ; n > 0; n-- {
		next, d := f.step(s.pos, s.d)
		if f.at(next) == wall {
			break
		}
		s.pos = next
		s.d = d
	}
}

func (s *self) turn(dir byte) {
	switch dir {
	case 'L':
		s.d.row, s.d.col = -s.d.col, s.d.row
	case 'R':
		s.d.row, s.d.col = s.d.col, -s.d.row
	}
}

func (s *self) location() (x, y int) {
	return s.pos.col + 1, s.pos.row + 1
}

func (s *self) facing() int {
	if s.d.col == 0 {
		return 2 - s.d.row // 3 or 1
	}
	return 1 - s.d.col // 2 or 0
}

func (s *self) password() int {
	x, y := s.location()
	facing := s.facing()
	return 1000*y + 4*x + facing
}

func parseField(s []byte) (field, error) {
	return bytes.Split(s, []byte{'\n'}), nil
}

type instruction func(*self, shape)

func parseInstructions(s []byte) ([]instruction, error) {
	inst := []instruction{}
	for len(s) != 0 {
		switch s[0] {
		case 'L', 'R':
			dir := s[0]
			inst = append(inst, func(me *self, _ shape) { me.turn(dir) })
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
		inst = append(inst, func(s *self, f shape) { s.forward(n, f) })
		s = s[i:]
	}
	return inst, nil
}

type solution struct {
	f            field
	instructions []instruction
}

func (sol *solution) Parse(s []byte) error {
	m, i, _ := bytes.Cut(s, []byte("\n\n"))
	f, err := parseField(m)
	if err != nil {
		return err
	}
	instructions, err := parseInstructions(i)
	if err != nil {
		return err
	}
	sol.f = f
	sol.instructions = instructions
	return nil
}

func (sol solution) Part1() {
	s := self{pos: sol.f.start(), d: direction{col: 1}}
	for _, f := range sol.instructions {
		f(&s, sol.f)
	}
	fmt.Println("Part 1", s.password())
}

type region struct {
	anchor position
}

func (r region) String() string {
	row, col := r.anchor.row/cubeSize, r.anchor.col/cubeSize
	return fmt.Sprintf("%c%d", rune('A'+row), col+1)
}

func (r region) corners() []position {
	p := r.anchor
	return []position{
		p,
		{row: p.row, col: p.col + cubeSize - 1},
		{row: p.row + cubeSize - 1, col: p.col},
		{row: p.row + cubeSize - 1, col: p.col + cubeSize - 1},
	}
}

func (r region) neighbors() []region {
	p := r.anchor
	n := []region{}
	for _, row := range []int{p.row - cubeSize, p.row + cubeSize} {
		n = append(n, region{position{row, p.col}})
	}
	for _, col := range []int{p.col - cubeSize, p.col + cubeSize} {
		n = append(n, region{position{p.row, col}})
	}
	return n
}

func (r region) present(f field) bool {
	return f.at(r.anchor) != blank
}

type cube struct {
	f       field
	edges   map[edge][2]cubes.Vertex
	corners map[corner]position
	dims    map[edge]cubes.Dimension
}

func (c cube) start() position {
	return c.f.start()
}

func (c cube) at(p position) byte {
	return c.f.at(p)
}

func (p position) region() (region, int, int) {
	a := position{
		row: p.row - p.row%cubeSize,
		col: p.col - p.col%cubeSize,
	}
	return region{anchor: a}, p.row % cubeSize, p.col % cubeSize
}

func (c cube) step(p position, d direction) (position, direction) {
	next := position{row: p.row + d.row, col: p.col + d.col}
	if c.at(next) != blank {
		return next, d
	}
	r, row, col := p.region()
	e := edge{r: r, d: d}
	vs := c.edges[e]
	dim := c.dims[e]
	p0 := c.corners[corner{v: vs[0], dim: dim}]
	p1 := c.corners[corner{v: vs[1], dim: dim}]
	offset := col
	if d.row == 0 {
		// previously, horizontal movement
		offset = row
	}
	if p0.row == p1.row {
		// next movement is vertical
		next.row = p0.row
		next.col = p0.col + (p1.col-p0.col)/(cubeSize-1)*offset
		d.col = 0
		d.row = 1
		if _, row, _ := next.region(); row != 0 {
			d.row = -1
		}
		return next, d
	}
	// next movement is horizontal
	next.col = p0.col
	next.row = p0.row + (p1.row-p0.row)/(cubeSize-1)*offset
	d.row = 0
	d.col = 1
	if _, _, col = next.region(); col != 0 {
		d.col = -1
	}
	return next, d
}

type corner struct {
	v   cubes.Vertex
	dim cubes.Dimension
}

type edge struct {
	r region
	d direction
}

func buildCube(f field) cube {
	visited := map[region]struct{}{}
	verts := map[position]cubes.Vertex{}
	edges := map[edge][2]cubes.Vertex{}
	dims := map[edge]cubes.Dimension{}
	corners := map[corner]position{}
	rowDims := map[region]cubes.Dimension{}
	colDims := map[region]cubes.Dimension{}
	neighbors := func(r region) []region {
		n := []region{}
		for _, r := range r.neighbors() {
			if r.present(f) {
				n = append(n, r)
			}
		}
		return n
	}
	setVert := func(v cubes.Vertex, dim cubes.Dimension, p position) {
		verts[p] = v
		corners[corner{v: v, dim: dim}] = p
	}
	setDimsAndVerts := func(row, col cubes.Dimension, r region, v cubes.Vertex) {
		fmt.Printf("setDimsAndVerts(row: %s, col: %s, region: %v, vertex: %d)\n", row, col, r, v)
		cs := r.corners()
		rowDims[r] = row
		colDims[r] = col
		dim := cubes.Ortho(row, col)
		v1 := cubes.Opposite(v, row)
		setVert(v, dim, cs[0])
		setVert(v1, dim, cs[1])
		setVert(cubes.Opposite(v, col), dim, cs[2])
		setVert(cubes.Opposite(v1, col), dim, cs[3])
		for d, ps := range map[direction][2]position{
			{row: -1, col: 0}: {cs[0], cs[1]},
			{row: 0, col: -1}: {cs[0], cs[2]},
			{row: 0, col: 1}:  {cs[1], cs[3]},
			{row: 1, col: 0}:  {cs[2], cs[3]},
		} {
			e := edge{r: r, d: d}
			edges[e] = [2]cubes.Vertex{verts[ps[0]], verts[ps[1]]}
			if d.col != 0 {
				dims[e] = row
			} else {
				dims[e] = col
			}
		}
	}
	visit := func(r region, ns []region) {
		defer func() { visited[r] = struct{}{} }()
		if len(verts) == 0 {
			// The first one is easy.
			setDimsAndVerts(cubes.X, cubes.Y, r, cubes.Vertex(0))
			return
		}
		for _, n := range ns {
			if _, ok := visited[n]; !ok {
				continue
			}
			if n.anchor.row == r.anchor.row {
				// Horizontal connection
				fmt.Printf("horizontal connection %v -> %v\n", n, r)
				colDim := colDims[n]
				rowDim := cubes.Ortho(colDim, rowDims[n])
				v := verts[n.anchor]
				if r.anchor.col > n.anchor.col {
					v = cubes.Opposite(v, rowDims[n])
				} else {
					v = cubes.Opposite(v, rowDim)
				}
				setDimsAndVerts(rowDim, colDim, r, v)
			} else {
				// Vertical connection
				fmt.Printf("vertical connection %v -> %v\n", n, r)
				rowDim := rowDims[n]
				colDim := cubes.Ortho(rowDim, colDims[n])
				v := verts[n.anchor]
				if r.anchor.row > n.anchor.row {
					v = cubes.Opposite(v, colDims[n])
				} else {
					v = cubes.Opposite(v, colDim)
				}
				setDimsAndVerts(rowDim, colDim, r, v)
			}
			break
		}
	}
	aoc.BFS([]region{{anchor: f.start()}}, func(r region) []region {
		ns := neighbors(r)
		visit(r, ns)
		return ns
	})
	return cube{f: f, edges: edges, corners: corners, dims: dims}
}

func (sol solution) Part2() {
	c := buildCube(sol.f)
	s := self{pos: c.start(), d: direction{col: 1}}
	for _, f := range sol.instructions {
		f(&s, c)
	}
	fmt.Println("Part 2", s.password())
}

func init() {
	aoc.RegisterSolution("22", &solution{})
}
