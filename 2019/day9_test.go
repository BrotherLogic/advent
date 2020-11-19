package a2019

import (
	"testing"
)

func TestOutput(t *testing.T) {
	p := processor{
		program: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
	}

	for i := 0; i < 100; i++ {
		p.program = append(p.program, 0)
	}

	output, _ := p.run()

	if output != 1219070632396864 {
		t.Errorf("PAH")
	}
}
