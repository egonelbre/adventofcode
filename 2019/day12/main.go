package main

import "fmt"

func main() {
	initial := []Moon{
		{Pos: Vector{13, 9, 5}},
		{Pos: Vector{8, 14, -2}},
		{Pos: Vector{-5, 4, 11}},
		{Pos: Vector{2, -6, 1}},
	}

	moons := append([]Moon{}, initial...)
	for step := int32(0); step <= 1000; step++ {
		if step%100 == 0 {
			PrintState(step, moons)
		}
		Step(moons)
	}

	moons = append([]Moon{}, initial...)

	converged := Convergence(moons)
	fmt.Println("converges after ", converged)
}

func Convergence(moons []Moon) int64 {
	var x [4]Dimension
	var y [4]Dimension
	var z [4]Dimension

	for i, moon := range moons {
		x[i].Pos, x[i].Vel = moon.Pos.X, moon.Vel.X
		y[i].Pos, y[i].Vel = moon.Pos.Y, moon.Vel.Y
		z[i].Pos, z[i].Vel = moon.Pos.Z, moon.Vel.Z
	}

	xconv := DimensionConvergence(x)
	yconv := DimensionConvergence(y)
	zconv := DimensionConvergence(z)

	return LCM(xconv, LCM(yconv, zconv))
}

func DimensionConvergence(moons [4]Dimension) int64 {
	seen := map[[4]Dimension]struct{}{}
	step := int64(0)

	for {
		if _, ok := seen[moons]; ok {
			return step
		}
		seen[moons] = struct{}{}

		StepDimension(moons[:])
		step++
	}
}

func Step(moons []Moon) {
	for i := range moons {
		pos := moons[i].Pos

		force := Vector{}
		for k := range moons {
			target := &moons[k]
			gravity := target.Pos.Sub(pos).Sign()
			force = force.Add(gravity)
		}

		moon := &moons[i]
		moon.Vel = moon.Vel.Add(force)
	}

	for i := range moons {
		moon := &moons[i]
		moon.Pos = moon.Pos.Add(moon.Vel)
	}
}

type Dimension struct {
	Pos int32
	Vel int32
}

func StepDimension(moons []Dimension) {
	for i := range moons {
		pos := moons[i].Pos

		force := int32(0)
		for k := range moons {
			gravity := Sign(moons[k].Pos - pos)
			force = force + gravity
		}

		moons[i].Vel += force
	}

	for i := range moons {
		moons[i].Pos += moons[i].Vel
	}
}

func PrintState(step int32, moons []Moon) {
	fmt.Println("===", "After", step, "===")
	totalEnergy := int32(0)
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

func (moon *Moon) PotentialEnergy() int32 {
	return moon.Pos.ManhattanLength()
}

func (moon *Moon) KineticEnergy() int32 {
	return moon.Vel.ManhattanLength()
}

func (moon *Moon) Energy() int32 {
	return moon.PotentialEnergy() * moon.KineticEnergy()
}

type Vector struct{ X, Y, Z int32 }

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

func (a Vector) ManhattanLength() int32 {
	return Abs(a.X) + Abs(a.Y) + Abs(a.Z)
}

func Sign(v int32) int32 {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}

	return 0
}

func Abs(v int32) int32 {
	if v < 0 {
		return -v
	}
	return v
}

func GCD(a, b int64) int64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}
