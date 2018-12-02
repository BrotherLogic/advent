package main

import "fmt"
import "testing"

func TestDay17oP1(t *testing.T) {
	arr := []int{20,15,10,5,5}

	if ComputeAllPerms(arr,25) != 4 {
		fmt.Printf("Done\n")
		t.Errorf("Wrong Compute %v\n", ComputeAllPerms(arr,25))
	}

	if FindTheMin(arr,25) != 3 {
		t.Errorf("Wrong Min %v\n", FindTheMin(arr,25))
	}
}
