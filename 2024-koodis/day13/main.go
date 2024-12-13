package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	need3 := 0
	need2 := 0
	need1 := 0

	messsages := strings.Split(data, "----------------------------------\n")
	for _, message := range messsages {
		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		lines := strings.Split(message, "\n")
		height := len(lines)
		width := 0
		for _, line := range lines {
			line = strings.TrimSpace(line)
			n := len([]rune(line))

			fmt.Printf("%3d %q\n", n, line)
			width = max(width, n)
		}

		n3, n2, n1 := 0, 0, 0

		switch height {
		case 1:
			n1 += width
		case 2:
			n2 += width / 2
			if width%2 == 1 {
				n1 += 2
			}
		case 3:
			n3 += width / 3
			leftover := width % 3
			switch leftover {
			case 1:
				n1 += 3
			case 2:
				n2 += 1
				n1 += 2
			}
		}

		fmt.Println(n3, n2, n1)

		need3 += n3
		need2 += n2
		need1 += n1
	}

	fmt.Printf("%v, %v, %v\n", need3, need2, need1)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
