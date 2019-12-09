package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	universe, err := ParseUniverse(Input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid input: %v\n", err)
		os.Exit(1)
	}

	// PART 1

	com, ok := universe.ObjectByName["COM"]
	if !ok {
		fmt.Fprintf(os.Stderr, "COM not found\n")
		os.Exit(1)
	}

	indirect := IndirectOrbits{}
	err = indirect.Add(com, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "indirect computation failed: %v\n", err)
	}

	fmt.Println("total objects", len(universe.Objects))
	fmt.Println("total direct and indirect orbits", indirect.Total())

	// PART 2
	you, ok := universe.ObjectByName["YOU"]
	if !ok {
		fmt.Fprintf(os.Stderr, "YOU not found\n")
		os.Exit(1)
	}

	san, ok := universe.ObjectByName["SAN"]
	if !ok {
		fmt.Fprintf(os.Stderr, "SAN not found\n")
		os.Exit(1)
	}

	common := CommonAncestor(you, san)
	if common == nil {
		fmt.Fprintf(os.Stderr, "did not find common ancestor\n")
		os.Exit(1)
	}

	youtx := indirect[you] - indirect[common] - 1
	santx := indirect[san] - indirect[common] - 1
	fmt.Println("Minimum transfers from YOU to SAN", youtx+santx)
}

type IndirectOrbits map[*Object]int64

func (info IndirectOrbits) Total() int64 {
	var total int64
	for _, v := range info {
		total += v
	}
	return total
}

func (info IndirectOrbits) Add(obj *Object, orbits int64) error {
	if _, added := info[obj]; added {
		return fmt.Errorf("cycle detected %v", obj)
	}

	info[obj] = orbits
	for _, satellite := range obj.Satellites {
		if err := info.Add(satellite, orbits+1); err != nil {
			return err
		}
	}

	return nil
}

type Universe struct {
	Objects      []*Object
	ObjectByName map[string]*Object
}

func NewUniverse() *Universe {
	return &Universe{
		Objects:      nil,
		ObjectByName: make(map[string]*Object),
	}
}

func ParseUniverse(input string) (*Universe, error) {
	universe := NewUniverse()
	for _, orbitdef := range strings.Split(input, "\n") {
		orbitdef = strings.TrimSpace(orbitdef)
		if orbitdef == "" {
			continue
		}

		tokens := strings.SplitN(orbitdef, ")", 2)
		if len(tokens) < 2 {
			return nil, fmt.Errorf("invalid orbit definition %q", orbitdef)
		}

		planet := universe.EnsureObject(tokens[0])
		satellite := universe.EnsureObject(tokens[1])

		planet.EnsureSatellite(satellite)
	}

	return universe, nil
}

func (universe *Universe) EnsureObject(name string) *Object {
	obj, ok := universe.ObjectByName[name]
	if ok {
		return obj
	}

	obj = &Object{Name: name}
	universe.Objects = append(universe.Objects, obj)
	universe.ObjectByName[name] = obj

	return obj
}

type Object struct {
	Name string

	Planet     *Object
	Satellites []*Object
}

func (planet *Object) HasSatellite(satellite *Object) bool {
	for _, x := range planet.Satellites {
		if x == satellite {
			return true
		}
	}
	return false
}

func (planet *Object) EnsureSatellite(satellite *Object) {
	if planet.HasSatellite(satellite) {
		return
	}

	planet.Satellites = append(planet.Satellites, satellite)
	satellite.Planet = planet
}

func CommonAncestor(a, b *Object) *Object {
	for ax := a.Planet; ax != nil; ax = ax.Planet {
		for bx := b.Planet; bx != nil; bx = bx.Planet {
			if bx == ax {
				return ax
			}
		}
	}
	return nil
}
