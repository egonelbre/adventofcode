// +build ignore

package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

type Chemistry struct {
	LastID     Atom
	ToAtom     map[string]Atom
	ToString   map[Atom]string
	Transforms map[Atom][]Molecule
}

func NewChemistry() *Chemistry {
	return &Chemistry{
		LastID:     0x61,
		ToAtom:     make(map[string]Atom, 256),
		ToString:   make(map[Atom]string, 256),
		Transforms: make(map[Atom][]Molecule, 256),
	}
}

func (chem *Chemistry) Atom(name string) Atom {
	atom, ok := chem.ToAtom[name]
	if !ok {
		if len(name) == 1 {
			atom = Atom(name[0])
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

func (chem *Chemistry) AddTransform(from Atom, into Molecule) {
	chem.Transforms[from] = append(chem.Transforms[from], into)
}

type Atom byte
type Molecule []Atom

func (m Molecule) String() string {
	r := make([]byte, len(m))
	for i, x := range m {
		r[i] = byte(x)
	}
	return string(r)
}

type Transforms map[Atom][][]Atom

type MoleculeSet map[string]struct{}

func (set MoleculeSet) Add(m Molecule) {
	set[m.String()] = struct{}{}
}

func Enum(mol Molecule, all MoleculeSet, chem *Chemistry) {
	for i, atom := range mol {
		for _, repl := range chem.Transforms[atom] {
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
	for _, repl := range chem.Transforms[atom] {
		EnumR(append(head, repl...), rest, all, chem)
	}
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
		chem.AddTransform(from, into)
	}

	s.Scan()
	target := chem.Molecule(s.Text())
	fmt.Printf("%+v", chem)
	fmt.Println(target)

	all := make(MoleculeSet)
	Enum(target, all, chem)
	fmt.Println(len(all))
}

func main() {
	solve(practice)
	solve(input)
}

var practice = `H => HO
H => OH
O => HH

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
