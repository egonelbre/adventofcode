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

	lines := strings.Split(data, "\n")

	code := strings.TrimSpace(lines[0])
	container := strings.TrimSpace(lines[1])

	fmt.Println("code:", code)
	fmt.Println("inpu:", container)

	matches := 0

	for length := len(code); length < len(container); length++ {
		fmt.Println(length)
		count := [256]int{}
		for _, b := range []byte(container[:length]) {
			count[b]++
		}
		if match(&count, code) {
			fmt.Println("match", 0, container[:length])
		}
		for tail := length; tail < len(container); tail++ {
			count[container[tail-length]]--
			count[container[tail]]++

			if match(&count, code) {
				fmt.Println("match:", tail-length, container[tail-length+1:tail+1])
				matches++
				if matches > 10 {
					return
				}
			}
		}
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func match(data *[256]int, code string) bool {
	for _, b := range []byte(code) {
		if data[b] == 0 {
			return false
		}
	}
	return true
}
