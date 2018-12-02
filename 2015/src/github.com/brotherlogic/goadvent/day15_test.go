package main

import "testing"

func TestDay15oP1(t *testing.T) {
	arr := []int{44,56}

	var m map[int][]int
	m = make(map[int][]int)
	m[0] = []int{-1,-2,6,3,8}
	m[1] = []int{2,3,-2,-1,3}

	if ComputeArr(arr,m, false) != 62842880 {
		t.Errorf("Wrong Compute %v\n", ComputeArr(arr,m, false))
	}

	if ProcArrs(2, m, 100, false) != 62842880 {
		t.Errorf("Wrong Proc %v\n", ProcArrs(2, m, 100, false))
	}
	if ProcArrs(2, m, 100, true) != 57600000 {
		t.Errorf("Wrong Proc %v\n", ProcArrs(2, m, 100, true))
	}

}
