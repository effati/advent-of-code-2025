package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rollingWindowQueue struct {
	batteries []int
	maxSize   int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	splittedInput := cleanup(string(input))
	fmt.Println(solution(splittedInput, 2))  // part 1
	fmt.Println(solution(splittedInput, 12)) // part 2
}

func solution(input [][]int, maxSize int) int {
	sum := 0
	for _, line := range input {
		queue := newRollingWindowQueue(maxSize)
		for _, battery := range line {
			queue.replaceIfBigger(battery)
		}
		sum += queue.jointVoltage()
	}
	return sum
}

func cleanup(input string) [][]int {
	splittedInput := strings.Split(input, "\n")
	for _, line := range splittedInput {
		if line == "" || len(line) < 2 {
			panic("invalid line: " + line)
		}
	}

	batteries := [][]int{}
	for _, line := range splittedInput {
		battery := []int{}
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			battery = append(battery, digit)
		}
		batteries = append(batteries, battery)
	}
	return batteries
}

func newRollingWindowQueue(maxSize int) rollingWindowQueue {
	return rollingWindowQueue{
		batteries: make([]int, maxSize),
		maxSize:   maxSize,
	}
}

func (q *rollingWindowQueue) replaceIfBigger(battery int) {
	bestQueue := newRollingWindowQueue(q.maxSize)
	copy(bestQueue.batteries, q.batteries)
	bestJointVoltage := bestQueue.jointVoltage()
	for i := 0; i < q.maxSize; i++ {
		tmpQueue := newRollingWindowQueue(q.maxSize)
		copy(tmpQueue.batteries, q.batteries)
		tmpQueue.batteries = append(tmpQueue.batteries[:i], tmpQueue.batteries[i+1:]...)
		tmpQueue.batteries = append(tmpQueue.batteries, battery)
		tmpJointVoltage := tmpQueue.jointVoltage()
		if tmpJointVoltage > bestJointVoltage {
			bestJointVoltage = tmpJointVoltage
			copy(bestQueue.batteries, tmpQueue.batteries)
		}
	}
	*q = bestQueue
}

func (q *rollingWindowQueue) jointVoltage() int {
	jointVoltage := strings.Builder{}
	for _, battery := range q.batteries {
		jointVoltage.WriteString(fmt.Sprintf("%d", battery))
	}
	jointVoltageInt, err := strconv.Atoi(jointVoltage.String())
	if err != nil {
		panic(err)
	}
	return jointVoltageInt
}
