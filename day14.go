// +build ignore

package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Raindeer struct {
	Name   string
	Speed  int // km/s
	Sprint int // seconds
	Rest   int // seconds

	Distance   int
	Sprinting  bool
	Points     int
	ActionTime int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Distance(r *Raindeer, seconds int) int {
	dist := 0
	time := 0
	for time < seconds {
		sprinttime := min(r.Sprint, seconds-time)
		dist += r.Speed * sprinttime
		time += sprinttime

		resttime := min(r.Rest, seconds-time)
		time += resttime
	}

	return dist
}

func Simulate(rs []*Raindeer, seconds int) {
	for i := 0; i < seconds; i++ {
		max := 0
		for _, r := range rs {
			// start next action
			if r.ActionTime == 0 {
				r.Sprinting = !r.Sprinting
				if r.Sprinting {
					r.ActionTime = r.Sprint
				} else {
					r.ActionTime = r.Rest
				}
			}
			// do action
			if r.Sprinting {
				r.Distance += r.Speed
			}
			r.ActionTime--

			// find max
			if max < r.Distance {
				max = r.Distance
			}
		}

		for _, r := range rs {
			if r.Distance >= max {
				r.Points++
			}
		}
	}
}

func main() {
	raindeer := []*Raindeer{}

	s := bufio.NewScanner(strings.NewReader(lines))
	for s.Scan() {
		line := s.Text()

		var name string
		var speed, sprint, rest int

		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&name, &speed, &sprint, &rest)
		if err != nil {
			panic(err)
		}

		r := &Raindeer{}
		r.Name = name
		r.Speed = speed
		r.Sprint = sprint
		r.Rest = rest
		raindeer = append(raindeer, r)
	}

	for _, r := range raindeer {
		fmt.Println(r.Name, Distance(r, 2503))
	}

	Simulate(raindeer, 2503)
	for _, r := range raindeer {
		fmt.Println(r)
	}
}

var lines = `Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.
Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.
Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.
Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.`
