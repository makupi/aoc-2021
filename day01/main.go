package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := fetchInput()
	part01(input)
	part02(input)
}

func fetchInput() []int64 {
	raw, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}
	input := bytes.Split(raw, []byte("\n"))
	var parsed []int64
	for _, line := range input {
		l, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			panic(err)
		}
		parsed = append(parsed, l)
	}
	return parsed
}

func part01(input []int64) {
	increased := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increased++
		}
	}
	fmt.Printf("Part 01: Increased %d times\n", increased)
}

func part02(input []int64) {
	last := input[0] + input[1] + input[2]
	increased := 0
	for i := 3; i < len(input); i++ {
		cur := input[i] + input[i-1] + input[i-2]
		if cur > last {
			increased++
		}
		last = cur
	}
	fmt.Printf("Part 02: Increased %d times\n", increased)
}
