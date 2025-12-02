package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	p1 := part1(input)
	if p1 != 3 {
		t.Errorf("expected 3, got %d", p1)
	}
}
