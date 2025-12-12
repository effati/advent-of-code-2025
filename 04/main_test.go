package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	expectedP1 := 13
	p1, _ := singleRemoval(cleanup(input), 3)
	if p1 != expectedP1 {
		t.Errorf("expected %d, got %d", expectedP1, p1)
	}

	expectedP2 := 43
	p2 := solution(cleanup(input), 3)
	if p2 != expectedP2 {
		t.Errorf("expected %d, got %d", expectedP2, p2)
	}
}
