package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func cleanup(input string) []idRange {
	idRanges := []idRange{}
	splittedInput := strings.Split(input, ",")
	for _, line := range splittedInput {
		if line == "" {
			panic("invalid line: " + line)
		}

		splittedLine := strings.Split(line, "-")
		if len(splittedLine) != 2 {
			panic("invalid line: " + line)
		}
		start, err := strconv.Atoi(splittedLine[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(splittedLine[1])
		if err != nil {
			panic(err)
		}

		idRanges = append(idRanges, idRange{start, end})
	}
	return idRanges
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	idRanges := cleanup(string(input))
	fmt.Println(part1(idRanges))
	fmt.Println(part2(idRanges))
}

func part1(input []idRange) int {
	invalidIDSum := 0
	for _, idRange := range input {
		for i := idRange.start; i <= idRange.end; i++ {
			if !isValidIDPart1(strconv.Itoa(i)) {
				invalidIDSum += i
			}
		}
	}
	return invalidIDSum
}

func isValidIDPart1(id string) bool {
	n := len(id)
	if n%2 != 0 {
		return true
	}

	first := id[:n/2]
	second := id[n/2:]

	return first != second
}

func part2(input []idRange) int {
	invalidIDSum := 0
	for _, idRange := range input {
		for i := idRange.start; i <= idRange.end; i++ {
			if !isValidIDPart2(strconv.Itoa(i)) {
				invalidIDSum += i
			}
		}
	}
	return invalidIDSum
}

func isValidIDPart2(id string) bool {
	concat := id + id
	cutoff := concat[1 : len(concat)-1]
	if strings.Contains(cutoff, id) {
		return false
	}
	return true
}
