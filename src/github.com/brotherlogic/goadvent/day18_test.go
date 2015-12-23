package main

import "fmt"
import "testing"

func TestDay18P1(t *testing.T) {
	arr := make([]string,0)
	arr = append(arr, ".#.#.#")
	arr = append(arr, "...##.")
	arr = append(arr, "#....#")
	arr = append(arr, "..#...")
	arr = append(arr, "#.#..#")
	arr = append(arr, "####..")

	initArr := BuildArr(arr)

	if CountArr(initArr) != 15 {
		t.Errorf("Wrong init count %v\n",CountArr(initArr))
	}

	fmt.Printf("%v,%v\n", ComputeNeighbors(initArr,0,0),initArr)
	following1 := Iterate(initArr)
	fmt.Printf("11,%v, %v\n",CountArr(following1),following1)
	following2 := Iterate(following1)
	following3 := Iterate(following2)
	following4 := Iterate(following3)

	if CountArr(following4) != 4 {
		t.Errorf("Wrong four count %v\n", CountArr(following4))
	}
}
