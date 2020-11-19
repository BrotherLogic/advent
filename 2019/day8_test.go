package a2019

import "testing"

func TestCount(t *testing.T) {
	val := FindLayer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 3, 2)
	if val != 0 {
		t.Errorf("Bad count: %v", val)
	}
}

func TestPrint(t *testing.T) {
	PrintGraph([]int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}, 2, 2)
	//t.Errorf("Oops")
}
