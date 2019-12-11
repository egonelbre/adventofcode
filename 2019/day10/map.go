package main

import (
	"fmt"
	"math"
	"strings"
)

type Vector struct {
	X, Y int64
}

func (a Vector) IsZero() bool {
	return a == Vector{}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}
func (a Vector) SquareLength() int64 {
	return a.X*a.X + a.Y*a.Y
}
func (a Vector) Angle() float64 {
	tx := float64(a.X)
	ty := float64(a.Y)
	return math.Atan2(-tx, ty)
}

type Map struct {
	Data []Tile
	Size Vector
}

type Tile = byte

const (
	Outside  = Tile(0)
	Empty    = Tile('.')
	Asteroid = Tile('#')
)

func NewMap(size Vector) *Map {
	return &Map{
		Data: make([]byte, size.X*size.Y),
		Size: size,
	}
}

func ParseMap(lines string) (*Map, error) {
	m := &Map{}

	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if m.Size.X == 0 {
			m.Size.X = int64(len(line))
		}

		if m.Size.X != int64(len(line)) {
			return nil, fmt.Errorf("invalid line length %v expected %v", len(line), m.Size.X)
		}
		m.Data = append(m.Data, []byte(line)...)
		m.Size.Y++
	}

	return m, nil
}

func (m *Map) Fill(c Tile) {
	for i := range m.Data {
		m.Data[i] = c
	}
}

func (m *Map) Print() {
	for y := int64(0); y < m.Size.Y; y++ {
		for x := int64(0); x < m.Size.X; x++ {
			var c = m.Data[m.Index(x, y)]
			fmt.Print(rune(c))
		}
		fmt.Println()
	}
}

func (m *Map) Index(x, y int64) int64 {
	return m.Size.X*y + x
}

func (m *Map) Contains(at Vector) bool {
	return at.X >= 0 && at.Y >= 0 && at.X < m.Size.X && at.Y < m.Size.Y
}

func (m *Map) At(at Vector) Tile {
	if !m.Contains(at) {
		return Outside
	}
	return m.Data[m.Size.X*at.Y+at.X]
}

func (m *Map) Count(b Tile) int64 {
	var count int64
	for _, v := range m.Data {
		if v == b {
			count++
		}
	}
	return count
}
