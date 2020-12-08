package a2020

import "testing"

func TestCode(t *testing.T) {
	program := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	res := RunCode(program)
	if res != 5 {
		t.Errorf("Bad acc: %v", res)
	}

	res = FixAndRunCode(program)
	if res != 8 {
		t.Errorf("Bad acc: %v", res)
	}
}
