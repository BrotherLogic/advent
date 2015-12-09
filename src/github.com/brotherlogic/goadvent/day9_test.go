package main

import "testing"

func TestDay9P1(t *testing.T) {

	arr := [][]int {}
	for i := 0 ; i < 3 ; i++ {
		row := make([]int,3)
		arr = append(arr,row)
	}

	arr[0][1] = 464
	arr[0][2] = 518
	arr[1][0] = 464
	arr[1][2] = 141
	arr[2][0] = 518
	arr[2][1] = 141
	
	answer := SearchMap(arr)
	answer2 := SearchMapMax(arr)
	
	if answer != 605 {
		t.Errorf("%d, want %d", answer, 605)
	}

	if answer2 != 982 {
		t.Errorf("%d, want %d", answer2, 982)
	}
}
