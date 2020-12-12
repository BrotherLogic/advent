package a2020

import "testing"

func TestMoves(t *testing.T) {
	moves := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	dist := MakeMoves(moves)
	if dist != 25 {
		t.Errorf("Bad moves: %v", dist)
	}

	ndist := MoveWaypoint(moves)
	if ndist != 286 {
		t.Errorf("Bad waypoint: %v", ndist)
	}
}
