package main

type Reactions struct {
	List []Reaction
}

func NewReactions() *Reactions {
	return &Reactions{
		List: []Reaction{},
	}
}

func (reactions *Reactions) Add(reaction Reaction) {
	reactions.List = append(reactions.List, reaction)
}

func (reactions *Reactions) Sort() {
	// TODO: topological sort over dependencies
}

type Chemical struct {
	Name  string
	Count int
}

type Reaction struct {
	Input  []Chemical
	Output Chemical
}
