package a2019

import "testing"

type d3test struct {
	input1 []string
	input2 []string
	output int
	len    int
}

var d3tests = []d3test{
	{
		input1: []string{"R8", "U5", "L5", "D3"},
		input2: []string{"U7", "R6", "D4", "L4"},
		output: 6,
		len:    30},
	{
		input1: []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		input2: []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
		output: 159,
		len:    610},
	{
		input1: []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		input2: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
		output: 135,
		len:    410},
}

func TestDistance(t *testing.T) {
	for _, test := range d3tests {
		ti1 := make([]string, len(test.input1))
		ti2 := make([]string, len(test.input2))
		copy(ti1, test.input1)
		copy(ti2, test.input2)
		result := ComputeDistance(test.input1, test.input2)
		len := ComputeLen(test.input1, test.input2)
		if result != test.output {
			t.Errorf("Bad run: %v,%v led to %v, should be %v", ti1, ti2, result, test.output)
		}
		if len != test.len {
			t.Errorf("Bad len: %v,%v led to %v, should be %v", ti1, ti2, len, test.len)
		}
	}
}
