package a2019

import (
	"fmt"
)

//FindLayer finds the key layer
func FindLayer(nums []int, width, height int) int {
	layerMap := make(map[int]map[int]int)

	for i, num := range nums {
		layer := i / (width * height)
		//log.Printf("%v -> %v", layer, num)

		if _, ok := layerMap[layer]; !ok {
			layerMap[layer] = make(map[int]int)
		}

		layerMap[layer][num]++
	}

	lowest := layerMap[0][0]
	val := 0
	for _, mapper := range layerMap {
		if mapper[0] < lowest {
			lowest = mapper[0]
			val = mapper[1] * mapper[2]
		}
	}

	return val
}

//PrintGraph prints out the result
func PrintGraph(nums []int, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			found := false
			for i := x + y*width; i < len(nums); i += width * height {
				//log.Printf("%v -> %v", i, nums[i])
				if nums[i] != 2 {
					if nums[i] == 1 {
						fmt.Printf("X")
					} else {
						fmt.Printf(" ")
					}
					//log.Printf("FOUND")
					found = true
					break
				}
			}
			if !found {
				//log.Printf("NOT FOUND")
			}
			//return
		}
		//return
		fmt.Println()
	}
}
