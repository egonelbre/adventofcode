package main

type Reactions struct {
	ByOutput map[string]Reaction
	Sorted   []string
}

func NewReactions() *Reactions {
	return &Reactions{
		ByOutput: map[string]Reaction{},
	}
}

func (reactions *Reactions) Add(reaction Reaction) {
	reactions.ByOutput[reaction.Output.Name] = reaction
}

func (reactions *Reactions) Reduce(state map[string]int64) {
	for _, chem := range reactions.Sorted {
		need := state[chem]
		delete(state, chem)

		reaction := reactions.ByOutput[chem]
		repeat := DivCeil(need, reaction.Output.Count)

		for _, input := range reaction.Input {
			state[input.Name] += input.Count * repeat
		}
	}
}

func DivCeil(a, b int64) int64 {
	return (a + b - 1) / b
}

func (reactions *Reactions) Sort() {
	type Dependency struct {
		NeededBy int
		Needs    []string
	}
	needs := map[string]*Dependency{}

	for _, reaction := range reactions.ByOutput {
		dep, ok := needs[reaction.Output.Name]
		if !ok {
			dep = &Dependency{}
			needs[reaction.Output.Name] = dep
		}
		dep.Needs = reaction.Needs()
	}

	for _, node := range needs {
		for _, need := range node.Needs {
			if dep, ok := needs[need]; ok {
				dep.NeededBy++
			}
		}
	}

	reactions.Sorted = nil
	for len(needs) > 0 {
		for name, node := range needs {
			if node.NeededBy > 0 {
				continue
			}

			reactions.Sorted = append(reactions.Sorted, name)
			delete(needs, name)
			for _, need := range node.Needs {
				if dep, ok := needs[need]; ok {
					dep.NeededBy--
				}
			}
		}
	}
}

type Chemical struct {
	Name  string
	Count int64
}

type Reaction struct {
	Input  []Chemical
	Output Chemical
}

func (r Reaction) Needs() []string {
	var needs []string
	for _, in := range r.Input {
		needs = append(needs, in.Name)
	}
	return needs
}
