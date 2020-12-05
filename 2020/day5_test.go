package a2020

import (
	"log"
	"testing"
)

type test struct {
	input  string
	row    int
	column int
	seat   int
}

var tests = []test{
	{input: "FBFBBFFRLR", row: 44, column: 5, seat: 357},
	{input: "BFFFBBFRRR", row: 70, column: 7, seat: 567},
	{input: "FFFBBBFRRR", row: 14, column: 7, seat: 119},
	{input: "BBFFBBFRLL", row: 102, column: 4, seat: 820},
}

func TestTrans(t *testing.T) {
	for _, test := range tests {
		row, col, seat := findSeat(test.input)
		if row != test.row || col != test.column || seat != test.seat {
			t.Errorf("Bad location: %v -> %v, %v, %v (%v, %v, %v)", test.input, row, col, seat, test.row, test.column, test.seat)
		}
	}
}

func TestRange(t *testing.T) {
	ranges := []string{
		"LLL",
		"LLR",
		"LRL",
		"LRR",
		"RLL",
		"RLR",
		"RRL",
		"RRR",
	}

	for i, rang := range ranges {
		if locate(rang, 0, 0, 7) != i || convertBin(rang) != i {
			log.Printf("BAD %v -> %v", rang, locate(rang, 0, 0, 7))
		}
	}
}
