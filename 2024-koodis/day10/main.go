package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Block struct {
	ID      uint16
	Variant uint16
	S0, S1  uint16
	Height  uint16

	Larger  []*Block
	Smaller int

	Total int
	Best  []*Block
}

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	blocks := []Block{}

	for i, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		sizes := strings.Split(line, "x")

		x := uint16(must(strconv.Atoi(sizes[0])))
		y := uint16(must(strconv.Atoi(sizes[1])))
		z := uint16(must(strconv.Atoi(sizes[2])))

		switch {
		case x == y && y == z:
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 0,
				S0:      x, S1: x,
				Height: x,
			})

		case x == y:
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 0,
				S0:      x, S1: x,
				Height: z,
			})

			s0, s1 := sort2(x, z)
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 1,
				S0:      s0, S1: s1,
				Height: x,
			})

		case y == z:
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 0,
				S0:      y, S1: y,
				Height: x,
			})

			s0, s1 := sort2(y, x)
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 1,
				S0:      s0, S1: s1,
				Height: y,
			})

		case x == z:
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 0,
				S0:      z, S1: z,
				Height: y,
			})

			s0, s1 := sort2(y, z)
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 1,
				S0:      s0, S1: s1,
				Height: z,
			})

		default:
			s0, s1 := sort2(x, y)
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 0,
				S0:      s0, S1: s1,
				Height: z,
			})

			s0, s1 = sort2(y, z)
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 1,
				S0:      s0, S1: s1,
				Height: x,
			})

			s0, s1 = sort2(z, x)
			blocks = append(blocks, Block{
				ID:      uint16(i),
				Variant: 2,
				S0:      s0, S1: s1,
				Height: y,
			})
		}
	}

	for i := range blocks {
		block := &blocks[i]

		block.Total = int(block.Height)
		block.Best = []*Block{block}

		for k := range blocks {
			other := &blocks[k]

			if block.ID == other.ID {
				continue
			}

			if block.S0 < other.S0 && block.S1 < other.S1 {
				block.Smaller++
				other.Larger = append(other.Larger, block)
			}
		}
	}

	roots := []*Block{}
	smallest := []*Block{}
	for i := range blocks {
		block := &blocks[i]

		if block.Smaller == 0 {
			smallest = append(smallest, block)
		}
		if len(block.Larger) == 0 {
			roots = append(roots, block)
		}
	}

	for len(smallest) > 0 {
		small := smallest[len(smallest)-1]
		smallest = smallest[:len(smallest)-1]
		if len(small.Larger) == 0 {
			continue
		}

		for _, larger := range small.Larger {
			if larger.Total < int(larger.Height)+small.Total {
				larger.Total = int(larger.Height) + small.Total
				larger.Best = append([]*Block{larger}, small.Best...)
			}
			larger.Smaller--
			if larger.Smaller == 0 {
				smallest = append(smallest, larger)
			}
		}
	}

	slices.SortFunc(roots, func(a, b *Block) int {
		return a.Total - b.Total
	})
	for _, root := range roots {
		fmt.Println(root.Total)
		for _, chain := range root.Best {
			fmt.Printf("\t%v\n", chain.Size())
		}
	}
}

func outputDot(w io.Writer, blocks []Block) {
	printf := func(format string, args ...any) {
		fmt.Fprintf(w, format, args...)
	}

	printf("digraph X {\n")
	defer printf("}\n")

	for i := range blocks {
		block := &blocks[i]
		printf("\t%v [label=\"#%v %vx%v %v\"]\n", block, block.ID, block.S0, block.S1, block.Height)
	}

	for i := range blocks {
		block := &blocks[i]
		for _, larger := range block.Larger {
			printf("\t%v -> %v\n", larger, block)
		}
	}
}

func (block *Block) Size() string {
	return fmt.Sprintf("%vx%v %v", block.S0, block.S1, block.Height)
}

func (block *Block) String() string {
	return fmt.Sprintf("p%vx%v", block.ID, block.Variant)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func sort2(a, b uint16) (uint16, uint16) {
	if a < b {
		return a, b
	}
	return b, a
}
