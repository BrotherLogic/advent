package main

import "testing"

var day3Part1Table = []struct {
	in  []string
	out int
}{
	{[]string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}, 4}}

var day3Part2Table = []struct {
	in  []string
	out int
}{
	{[]string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}, 3}}

func TestDay3Part1(t *testing.T) {
	for _, tt := range day3Part1Table {
		if computeOverlap(tt.in) != tt.out {
			t.Errorf("Error in overlap computation: %v", computeOverlap(tt.in))
		}
	}
}

func TestDay3Part2(t *testing.T) {
	for _, tt := range day3Part2Table {
		if computeUnique(tt.in) != tt.out {
			t.Errorf("Error in overlap computation: %v", computeUnique(tt.in))
		}
	}
}
