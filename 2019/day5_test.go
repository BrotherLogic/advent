package a2019

import "testing"

func TestTheBasic(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	result := ProcessOpCode2(program, 12)
	if result != 1 {
		t.Errorf("Bad run: %v", result)
	}
}

func TestBasicImmediate(t *testing.T) {
	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}

	result := ProcessOpCode2(program, 12)
	if result != 1 {
		t.Errorf("Bad run: %v", result)
	}
}

func TestBasicZero(t *testing.T) {
	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	result := ProcessOpCode2(program, 0)
	if result != 0 {
		t.Errorf("Bad run: %v", result)
	}
}
