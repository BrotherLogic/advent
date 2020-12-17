package a2020

import "testing"

type test15 struct {
	input  []int
	goal   int
	result int
}

var tests15 = []test15{
	{input: []int{0, 3, 6}, goal: 2020, result: 436},
	{input: []int{0, 3, 6}, goal: 30000000, result: 175594},
}

func TestGame1(t *testing.T) {
	for _, te := range tests15 {
		res := RunGame(te.input, te.goal)
		if res != te.result {
			t.Errorf("Bad game result: %v -> %v", res, te.result)
		}
	}
}
