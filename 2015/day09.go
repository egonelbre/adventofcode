// +build ignore

package main

import (
	"bufio"
	"fmt"
	"strings"
)

const MaxID = 10

type Map struct {
	LastID int
	ID     map[string]int
	Dist   [MaxID][MaxID]int
}

func (m *Map) City(name string) int {
	id, ok := m.ID[name]
	if !ok {
		id = m.LastID
		m.LastID++
		if m.LastID >= MaxID {
			panic("too many")
		}
		m.ID[name] = id
	}
	return id
}

func (m *Map) min(at int, tovisit int) int {
	if tovisit == 0 {
		return 0
	}

	min := 1 << 20
	for i := 0; i < m.LastID; i++ {
		bit := 1 << uint(i)
		if tovisit&bit == 0 {
			continue
		}
		dist := m.Dist[at][i] + m.min(i, tovisit&^bit)
		if dist < min {
			min = dist
		}
	}
	return min
}

func (m *Map) Min() int {
	min := 1 << 20
	tovisit := (1 << uint(m.LastID)) - 1
	for i := 0; i < m.LastID; i++ {
		bit := 1 << uint(i)
		dist := m.min(i, tovisit&^bit)
		if dist < min {
			min = dist
		}
	}
	return min
}

func (m *Map) max(at int, tovisit int) int {
	if tovisit == 0 {
		return 0
	}

	max := 0
	for i := 0; i < m.LastID; i++ {
		bit := 1 << uint(i)
		if tovisit&bit == 0 {
			continue
		}
		dist := m.Dist[at][i] + m.max(i, tovisit&^bit)
		if dist > max {
			max = dist
		}
	}
	return max
}

func (m *Map) Max() int {
	max := 0
	tovisit := (1 << uint(m.LastID)) - 1
	for i := 0; i < m.LastID; i++ {
		bit := 1 << uint(i)
		dist := m.max(i, tovisit&^bit)
		if dist > max {
			max = dist
		}
	}
	return max
}

type Dist [3][3]int

func main() {
	m := Map{}
	m.ID = make(map[string]int)

	s := bufio.NewScanner(strings.NewReader(lines))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		var from, to string
		var dist int
		_, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &dist)
		if err != nil {
			panic(err)
		}

		m.Dist[m.City(from)][m.City(to)] = dist
		m.Dist[m.City(to)][m.City(from)] = dist
	}

	fmt.Println(m.Min())
	fmt.Println(m.Max())
}

var lines = `Tristram to AlphaCentauri = 34
Tristram to Snowdin = 100
Tristram to Tambi = 63
Tristram to Faerun = 108
Tristram to Norrath = 111
Tristram to Straylight = 89
Tristram to Arbre = 132
AlphaCentauri to Snowdin = 4
AlphaCentauri to Tambi = 79
AlphaCentauri to Faerun = 44
AlphaCentauri to Norrath = 147
AlphaCentauri to Straylight = 133
AlphaCentauri to Arbre = 74
Snowdin to Tambi = 105
Snowdin to Faerun = 95
Snowdin to Norrath = 48
Snowdin to Straylight = 88
Snowdin to Arbre = 7
Tambi to Faerun = 68
Tambi to Norrath = 134
Tambi to Straylight = 107
Tambi to Arbre = 40
Faerun to Norrath = 11
Faerun to Straylight = 66
Faerun to Arbre = 144
Norrath to Straylight = 115
Norrath to Arbre = 135
Straylight to Arbre = 127`
