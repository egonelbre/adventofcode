package main

import (
	"fmt"
	"os"

	"github.com/egonelbre/adventofcode/2019/day15/g"
	"github.com/egonelbre/adventofcode/2019/day15/intcode"
)

const (
	Unknown = g.Color(0)
	Wall    = g.Color(1)
	Empty   = g.Color(2)
	Droid   = g.Color(3)
	Oxygen  = g.Color(4)
)

var colors = map[g.Color]rune{
	Unknown: ' ',
	Wall:    '█',
	Empty:   '.',
	Droid:   '🤖',
	Oxygen:  'O',
}

type Action = int64

const (
	GoNorth = Action(1)
	GoSouth = Action(2)
	GoWest  = Action(3)
	GoEast  = Action(4)
)

func Reverse(action Action) Action {
	switch action {
	case GoNorth:
		return GoSouth
	case GoSouth:
		return GoNorth
	case GoWest:
		return GoEast
	case GoEast:
		return GoWest
	}
	panic("invalid action")
}

func Movement(action Action) g.Vector {
	switch action {
	case GoNorth:
		return g.Vector{X: 0, Y: -1}
	case GoSouth:
		return g.Vector{X: 0, Y: 1}
	case GoWest:
		return g.Vector{X: -1, Y: 0}
	case GoEast:
		return g.Vector{X: 1, Y: 0}
	}
	panic("invalid action")
}

type Status = int64

const (
	HitWall     = Status(0)
	Moved       = Status(1)
	FoundOxygen = Status(2)

	Errored = Status(-1)
)

func main() {
	search := NewSearch()
	err := search.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	search.Map.Image().Print(colors)
}

type Search struct {
	Map   *g.SparseImage
	Code  map[g.Vector]intcode.Code
	Queue PendingActionQueue
}

type PendingActionQueue []PendingAction

func (q *PendingActionQueue) Enqueue(at g.Vector, action Action, traveled int64) {
	*q = append(*q, PendingAction{at, action, traveled})
}

func (q *PendingActionQueue) Dequeue() (PendingAction, bool) {
	if len(*q) == 0 {
		return PendingAction{}, false
	}

	first := (*q)[0]
	*q = (*q)[1:]
	return first, true
}

type PendingAction struct {
	At       g.Vector
	Action   Action
	Traveled int64
}

func NewSearch() *Search {
	search := &Search{
		Map:  g.NewSparseImage(Unknown),
		Code: make(map[g.Vector]intcode.Code),
	}

	zero := g.Vector{}

	search.Map.Set(zero, Empty)
	search.Code[zero] = DroidCode.Clone()
	search.AddExplore(zero, 0)

	return search
}

func (search *Search) AddExplore(at g.Vector, traveled int64) {
	for action := GoNorth; action <= GoEast; action++ {
		if search.Map.At(at.Add(Movement(action))) != Unknown {
			continue
		}

		search.Queue.Enqueue(at, action, traveled)
	}
}

func (search *Search) Run() error {
	for {
		act, ok := search.Queue.Dequeue()
		if !ok {
			return fmt.Errorf("did not find oxygen")
		}

		nextat := act.At.Add(Movement(act.Action))
		traveled := act.Traveled + 1

		if search.Map.At(nextat) != Unknown {
			continue
		}

		status, code, err := Explore(search.Code[act.At], act.At, act.Action)
		if err != nil {
			return err
		}

		switch status {
		case HitWall:
			search.Map.Set(nextat, Wall)
		case Moved:
			search.Map.Set(nextat, Empty)
			search.Code[nextat] = code
			search.AddExplore(nextat, traveled)
		case FoundOxygen:
			search.Map.Set(nextat, Oxygen)
			search.Code[nextat] = code
			search.AddExplore(nextat, traveled)

			fmt.Println("found oxygen", traveled)
			return nil
		}
	}
}

func Explore(code intcode.Code, at g.Vector, action Action) (Status, intcode.Code, error) {
	cpu := &intcode.Computer{
		Code: code.Clone(),
	}

	ok, err := intcode.WriteValue(cpu, action)
	if err != nil {
		return Errored, nil, err
	}
	if !ok {
		return Errored, nil, fmt.Errorf("writing value failed")
	}

	output, ok, err := intcode.ReadValue(cpu)
	if err != nil {
		return Errored, nil, err
	}
	if !ok {
		return Errored, nil, fmt.Errorf("reading value failed")
	}

	return output, cpu.Code, nil
}
