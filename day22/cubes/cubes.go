// Package cube handles cube geometry.
//
// A cube has 8 vertices, connected in a cubical shape as follows:
//
//	0-----1
//	|\    |\
//	| 4-----5
//	2-|---3 |
//	 \|    \|
//	  6-----7
package cubes

// A Vertex is a corner of a cube.
type Vertex uint

// Opposite returns the opposite Vertex across the given Dimension.
func Opposite(v Vertex, dim Dimension) Vertex {
	return v ^ Vertex(dim)
}

// A Dimension is one of the three dimensions X, Y, or Z.
type Dimension uint

const (
	X Dimension = 1 << iota
	Y
	Z
)

func (d Dimension) String() string {
	return map[Dimension]string{
		X: "X",
		Y: "Y",
		Z: "Z",
	}[d]
}

// Ortho returns the orthogonal dimension from the other two.
func Ortho(d1, d2 Dimension) Dimension {
	return 0b111 ^ d1 ^ d2
}
