package main

import "math"

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

func (a Vector) Min(b Vector) Vector {
	return Vector{
		X: Min(a.X, b.X),
		Y: Min(a.Y, b.Y),
	}
}

func (a Vector) Max(b Vector) Vector {
	return Vector{
		X: Max(a.X, b.X),
		Y: Max(a.Y, b.Y),
	}
}

func RotateLeft(v Vector) Vector {
	return Vector{v.Y, -v.X}
}

func RotateRight(v Vector) Vector {
	return Vector{-v.Y, v.X}
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
