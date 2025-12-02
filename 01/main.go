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
	for _, line := range splittedInput {
		if line == "" || len(line) < 2 {
			panic("invalid line: " + line)
		}
	}
	fmt.Println(part1(splittedInput))
	fmt.Println(part2(splittedInput))
}

func part1(input []string) int {
	dial := dialStart
	numberOfZeroes := 0

	for _, line := range input {
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

// I started out with negative numbers and mod and floors etc, but I gave up and just brute forced it
func part2(input []string) int {
	dial := dialStart
	zeroes := 0
	for _, line := range input {
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		for range steps {
			if line[0] == 'R' {
				dial++
			} else {
				dial--
			}

			dial = dial % 100

			if dial == 0 {
				zeroes++
			}
		}
	}

	return zeroes
}
