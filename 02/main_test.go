package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	p1 := part1(cleanup(input))
	if p1 != 1227775554 {
		t.Errorf("expected 1227775554, got %d", p1)
	}
	p2 := part2(cleanup(input))
	if p2 != 4174379265 {
		t.Errorf("expected 4174379265, got %d", p2)
	}
}
