package main

import "testing"

var day9Table = []struct {
	in  string
	out int
}{
	{"17 players; last marble is worth 1104 points", 2764},
	{"9 players; last marble is worth 25 points", 32},
	{"10 players; last marble is worth 1618 points", 8317},
	{"13 players; last marble is worth 7999 points", 146373},
	{"21 players; last marble is worth 6111 points", 54718},
	{"30 players; last marble is worth 5807 points", 37305},
}

func TestDay9(t *testing.T) {
	for _, tt := range day9Table {
		solve, _ := runGame(tt.in, -1)
		if solve != tt.out {
			t.Fatalf("Game was wrong: %v (%v)", solve, tt.out)
		}
	}
}
