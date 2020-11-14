package a2019

import "testing"

type test struct {
	input  []int
	output int
}

var tests = []test{
	{input: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, output: 3500},
	{input: []int{1, 0, 0, 0, 99}, output: 2},
	{input: []int{2, 3, 0, 3, 99}, output: 2},
	{input: []int{2, 4, 4, 5, 99, 0}, output: 2},
	{input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, output: 30},
}

func TestOpCode1(t *testing.T) {
	for _, test := range tests {
		ti := make([]int, len(test.input))
		copy(ti, test.input)
		result := RunOpCode1(test.input)
		if result != test.output {
			t.Errorf("Bad run: %v led to %v, should be %v", ti, result, test.output)
		}
	}
}
