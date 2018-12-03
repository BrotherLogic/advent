package main

import "testing"

var day2Part2Table = []struct {
	in  []string
	out string
}{
	{[]string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}, "fgij"}}

var day2Part1Table = []struct {
	in  []string
	out int
}{
	{[]string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}, 12}}

func TestDay2Part1(t *testing.T) {
	for _, tt := range day2Part1Table {
		if computeChecksum(tt.in) != tt.out {
			t.Errorf("Error in checksum: %v", computeChecksum(tt.in))
		}
	}
}

func TestDay2Part2(t *testing.T) {
	for _, tt := range day2Part2Table {
		if computeDiff(tt.in) != tt.out {
			t.Errorf("Error in diffCalc: %v", computeDiff(tt.in))
		}
	}
}
