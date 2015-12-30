package main

import "testing"

func TestDay24P1(t *testing.T) {

	numbers := []int {1,2,3,4,5,7,8,9,10,11}
	val := ProcNumbers(numbers,3)

	if val != 99 {
		t.Errorf("Failed on %v\n",val)
	}
}
