package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dialStart = 50

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	splittedInput := strings.Split(string(input), "\n")
	fmt.Println(part1(splittedInput))
}

func part1(input []string) int {
	dial := dialStart
	numberOfZeroes := 0

	for _, line := range input {
		if line == "" || len(line) < 2 {
			panic("invalid line: " + line)
		}
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if line[0] == 'R' {
			dial = (dial + steps) % 100
		} else {
			dial = (dial - steps) % 100
		}

		if dial == 0 {
			numberOfZeroes++
		}
	}

	return numberOfZeroes
}
