package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`

	expectedP1 := 357
	expectedP2 := 3121910778619
	p1 := solution(cleanup(input), 2)
	if p1 != expectedP1 {
		t.Errorf("expected %d, got %d", expectedP1, p1)
	}
	p2 := solution(cleanup(input), 12)
	if p2 != expectedP2 {
		t.Errorf("expected %d, got %d", expectedP2, p2)
	}
}
