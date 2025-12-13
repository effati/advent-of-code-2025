package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type (
	intRange struct {
		start int
		end   int
	}
	inventory struct {
		freshIngredients []intRange
		ingredientIDs    []int
	}
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(cleanup(string(input))))
	fmt.Println(part2(cleanup(string(input))))
}

func part1(inventory inventory) int {
	freshCount := 0
	for _, id := range inventory.ingredientIDs {
		if slices.ContainsFunc(inventory.freshIngredients, func(ir intRange) bool {
			return id >= ir.start && id <= ir.end
		}) {
			freshCount++
		}
	}
	return freshCount
}

func part2(inventory inventory) int {
	mergedRanges := mergeOverlappingRanges(inventory.freshIngredients)
	count := 0

	for _, ir := range mergedRanges {
		count += (ir.end - ir.start) + 1 // to include the exlusive end
	}

	return count
}

func mergeOverlappingRanges(ranges []intRange) []intRange {
	sortedRanges := make([]intRange, len(ranges))
	copy(sortedRanges, ranges)

	slices.SortFunc(sortedRanges, func(a intRange, b intRange) int {
		return cmp.Compare(a.start, b.start)
	})

	outRanges := []intRange{sortedRanges[0]}
	for i := 1; i < len(sortedRanges); i++ {
		last := outRanges[len(outRanges)-1]
		curr := sortedRanges[i]

		if curr.start <= last.end { // if the current range overlaps with the last, merge them
			last.end = max(curr.end, last.end)
			outRanges[len(outRanges)-1] = last
		} else {
			outRanges = append(outRanges, curr)
		}
	}

	return outRanges
}

func cleanup(input string) inventory {
	inventory := inventory{
		freshIngredients: []intRange{},
		ingredientIDs:    []int{},
	}
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rangeLines := strings.Split(sections[0], "\n")
	numberLines := strings.Split(sections[1], "\n")
	for _, line := range rangeLines {
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
		inventory.freshIngredients = append(inventory.freshIngredients, intRange{start, end})
	}

	for _, line := range numberLines {
		id, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		inventory.ingredientIDs = append(inventory.ingredientIDs, id)
	}
	return inventory
}
