package main

import "fmt"

func main() {
	moons := []Moon{
		{Pos: Vector{13, 9, 5}},
		{Pos: Vector{8, 14, -2}},
		{Pos: Vector{-5, 4, 11}},
		{Pos: Vector{2, -6, 1}},
	}

	for step := int64(0); step <= 1000; step++ {
		if step%10 == 0 {
			PrintState(step, moons)
		}
		Step(moons)
	}
}

func Step(moons []Moon) {
	for i := range moons {
		pos := moons[i].Pos

		force := Vector{}
		for k := range moons {
			if i == k {
				continue
			}
			target := &moons[k]
			gravity := target.Pos.Sub(pos).Sign()
			force = force.Add(gravity)
		}

		moon := &moons[i]
		moon.Vel = moon.Vel.Add(force)
		moon.Pos = moon.Pos.Add(moon.Vel)
	}
}

func PrintState(step int64, moons []Moon) {
	fmt.Println("===", "After", step, "===")
	totalEnergy := int64(0)
	for i := range moons {
		moon := &moons[i]
		energy := moon.Energy()
		totalEnergy += energy
		fmt.Printf("%#v E=%v\n", moon, moon.Energy())
	}
	fmt.Printf("E=%v\n", totalEnergy)
	fmt.Println()
}

type Moon struct {
	Pos Vector
	Vel Vector
}

func (moon *Moon) PotentialEnergy() int64 {
	return moon.Pos.ManhattanLength()
}

func (moon *Moon) KineticEnergy() int64 {
	return moon.Vel.ManhattanLength()
}

func (moon *Moon) Energy() int64 {
	return moon.PotentialEnergy() * moon.KineticEnergy()
}

type Vector struct{ X, Y, Z int64 }

func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Vector) Sign() Vector {
	return Vector{
		X: Sign(a.X),
		Y: Sign(a.Y),
		Z: Sign(a.Z),
	}
}

func (a Vector) ManhattanLength() int64 {
	return Abs(a.X) + Abs(a.Y) + Abs(a.Z)
}

func Sign(v int64) int64 {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}

	return 0
}

func Abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
