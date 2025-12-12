package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := cleanup(string(input))
	fmt.Println(singleRemoval(grid, 3)) // part 1
	fmt.Println(solution(grid, 3))      // part 2
}

func solution(input [][]string, maxAdjacentRolls int) int {
	rollCount := 0
	grid := input
	for {
		singleRollCount, newGrid := singleRemoval(grid, maxAdjacentRolls)
		if singleRollCount == 0 {
			break
		}
		rollCount += singleRollCount
		grid = newGrid
	}
	return rollCount
}

func singleRemoval(input [][]string, maxAdjacentRolls int) (int, [][]string) {
	newGrid := make([][]string, len(input))
	for i := range input {
		newGrid[i] = make([]string, len(input[i]))
		copy(newGrid[i], input[i])
	}
	neighborDeltas := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	rollCount := 0

	for y, row := range input {
		for x, char := range row {
			if !isRoll(char) {
				continue
			}
			neighborRollCount := 0
			for _, delta := range neighborDeltas {
				neighborY := y + delta[0]
				neighborX := x + delta[1]
				if neighborY < 0 || neighborY >= len(input) || neighborX < 0 || neighborX >= len(row) {
					continue
				}
				neighborChar := input[neighborY][neighborX]
				if isRoll(neighborChar) {
					neighborRollCount++
				}
			}
			if neighborRollCount <= maxAdjacentRolls {
				newGrid[y][x] = "."
				rollCount++
			}
		}
	}
	return rollCount, newGrid
}

func isRoll(char string) bool {
	return char == "@"
}

func cleanup(input string) [][]string {
	splittedInput := strings.Split(input, "\n")
	for _, line := range splittedInput {
		if line == "" {
			panic("invalid line: " + line)
		}
	}

	grid := [][]string{}
	for _, line := range splittedInput {
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	return grid
}
