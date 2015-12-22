// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Chemistry struct {
	LastID         byte
	ToAtom         map[string]byte
	ToString       map[byte]string
	ReactionByAtom map[byte][]Molecule
	Reactions      []Reaction
}

type Reaction struct {
	From byte
	Into Molecule
}

func NewChemistry() *Chemistry {
	return &Chemistry{
		LastID:         'f',
		ToAtom:         make(map[string]byte, 256),
		ToString:       make(map[byte]string, 256),
		ReactionByAtom: make(map[byte][]Molecule, 256),
	}
}

func (chem *Chemistry) Atom(name string) byte {
	atom, ok := chem.ToAtom[name]
	if !ok {
		if len(name) == 1 {
			atom = byte(name[0])
		} else {
			atom = chem.LastID
			chem.LastID++
		}
		chem.ToAtom[name] = atom
		chem.ToString[atom] = name
	}
	return atom
}

func (chem *Chemistry) Molecule(molecule string) Molecule {
	mol := Molecule{}
	atomname := ""
	for _, r := range molecule {
		if unicode.IsUpper(r) {
			if atomname != "" {
				mol = append(mol, chem.Atom(atomname))
				atomname = ""
			}
		}
		atomname += string(r)
	}
	if atomname != "" {
		mol = append(mol, chem.Atom(atomname))
		atomname = ""
	}
	return mol
}

func (chem *Chemistry) AddReaction(from byte, into Molecule) {
	chem.ReactionByAtom[from] = append(chem.ReactionByAtom[from], into)
	chem.Reactions = append(chem.Reactions, Reaction{from, into})
}

type Molecule []byte

func (m Molecule) String() string { return string(m) }

type MoleculeSet map[string]struct{}

func (set MoleculeSet) Add(m Molecule) { set[m.String()] = struct{}{} }

func Enum(mol Molecule, all MoleculeSet, chem *Chemistry) {
	for i, atom := range mol {
		for _, repl := range chem.ReactionByAtom[atom] {
			x := append(append(mol[:i:i], repl...), mol[i+1:]...)
			all.Add(x)
		}
	}
}

func EnumR(head, tail Molecule, all MoleculeSet, chem *Chemistry) {
	all.Add(append(head, tail...))
	if len(tail) == 0 {
		return
	}

	atom, rest := tail[0], tail[1:]
	EnumR(append(head, atom), rest, all, chem)
	for _, repl := range chem.ReactionByAtom[atom] {
		EnumR(append(head, repl...), rest, all, chem)
	}
}

func IndexAll(mol Molecule, sub Molecule) (rs []int) {
	i := 0
	for i < len(mol) {
		x := bytes.Index(mol[i:], sub)
		if x < 0 {
			return rs
		}
		rs = append(rs, i+x)
		i = i + x + len(sub)
	}
	return rs
}

type MoleculeDist map[string]int

var smallest int

func Reduce(w int, mol Molecule, chem *Chemistry, cache MoleculeDist) int {
	if w > smallest {
		return 1 << 31
	}
	if v, cached := cache[mol.String()]; cached {
		return v
	}
	if len(mol) == 1 && mol[0] == 'e' {
		return 0
	}
	min := 1 << 31
	for _, react := range chem.Reactions {
		sub := react.Into
		matches := IndexAll(mol, sub)
		for _, m := range matches {
			next := append(append(mol[:m:m], react.From), mol[m+len(sub):]...)
			n := 1 + Reduce(w+1, next, chem, cache)
			if n < min {
				min = n
			}
		}
	}
	if min+w < smallest {
		smallest = min + w
		fmt.Println(smallest)
	}
	cache[mol.String()] = min
	return min
}

func solve(input string) {
	fmt.Println()
	fmt.Println("---")
	fmt.Println()
	chem := NewChemistry()
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		var fromname, intoname string
		_, err := fmt.Sscanf(line, "%s => %s", &fromname, &intoname)
		if err != nil {
			panic(err)
		}

		from := chem.Atom(fromname)
		into := chem.Molecule(intoname)
		chem.AddReaction(from, into)
	}

	s.Scan()
	target := chem.Molecule(s.Text())
	fmt.Printf("%+v\n", chem)
	fmt.Println(len(target), target)

	all := make(MoleculeSet)
	Enum(target, all, chem)
	fmt.Println("ALL:", len(all))

	cache := make(MoleculeDist)
	smallest = 1 << 31
	fmt.Println("RED:", Reduce(0, target, chem, cache))
}

func main() {
	solve(practice)
	solve(input)
}

var practice = `H => HO
H => OH
O => HH
e => H
e => O

HOH`

var input = `Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg

CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF`
