package a2019

import "testing"

func TestOpCode1(t *testing.T) {
	result := runOpCode1([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
	if result != 3500 {
		t.Errorf("Bad run: %v, should be 3500", result)
	}
}
