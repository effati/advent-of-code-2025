package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	expectedP1 := 3
	p1 := part1(cleanup(input))
	if p1 != expectedP1 {
		t.Errorf("expected %d, got %d", expectedP1, p1)
	}

	expectedP2 := 14
	p2 := part2(cleanup(input))
	if p2 != expectedP2 {
		t.Errorf("expected %d, got %d", expectedP2, p2)
	}

}
