package main

type Wire struct {
}

type Path []Segment

// Crossing logic

type AbsoluteSegment struct {
	At Vector
	Segment
}

func (a AbsoluteSegment) End() Vector {
	return a.At.Add(a.Vector())
}

func (a AbsoluteSegment) Canon() AbsoluteSegment {
	x := a
	switch x.Direction {
	case Left:
		x.At = a.End()
		x.Direction = Right
	case Up:
		x.At = a.End()
		x.Direction = Down
	}
	return x
}

func (a AbsoluteSegment) Contains(p Vector) bool {
	if a.At.X == p.X {
		y0, y1 := MinMax(a.At.Y, a.End().Y)
		return y0 <= p.Y && p.Y <= y1
	} else if a.At.Y == p.Y {
		x0, x1 := MinMax(a.At.X, a.End().X)
		return x0 <= p.X && p.X <= x1
	} else {
		return false
	}
}

func Crossing(a, b AbsoluteSegment) (Vector, bool) {
	// parallel lines cannot cross
	if a.IsHorizontal() == b.IsHorizontal() {
		return Vector{}, false
	}

	// one of them is horizontal and the other vertical
	var horizontal, vertical AbsoluteSegment
	if a.IsHorizontal() {
		horizontal, vertical = a, b
	} else {
		horizontal, vertical = b, a
	}

	expected := Vector{
		X: vertical.At.X,
		Y: horizontal.At.Y,
	}
	online := horizontal.Contains(expected) && vertical.Contains(expected)

	return expected, online
}

// Wire definition

type Segment struct {
	Direction
	Length int64
}

func (seg Segment) Vector() Vector {
	switch seg.Direction {
	case Up:
		return Vector{0, -seg.Length}
	case Down:
		return Vector{0, seg.Length}
	case Left:
		return Vector{-seg.Length, 0}
	case Right:
		return Vector{seg.Length, 0}
	default:
		return Vector{0, 0}
	}
}

type Direction byte

const (
	Left  Direction = 'L'
	Right           = 'R'
	Down            = 'D'
	Up              = 'U'
)

func (dir Direction) IsHorizontal() bool { return dir == Left || dir == Right }
func (dir Direction) IsVertical() bool   { return dir == Up || dir == Down }

// Vector definition

type Vector struct{ X, Y int64 }

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}
func (a Vector) ManhattanDist(b Vector) int64 {
	return a.Sub(b).ManhattanLength()
}
func (a Vector) ManhattanLength() int64 {
	return Abs(a.X) + Abs(a.Y)
}

func Abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

func MinMax(a, b int64) (min, max int64) {
	if a < b {
		return a, b
	}
	return b, a
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
