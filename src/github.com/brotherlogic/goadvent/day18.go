package main

import "bufio"
import "fmt"
import "os"

func ComputeNeighbors(arr [][]bool, x int, y int) int {
	count := 0
	for i := Max(0,x-1); i < Min(len(arr),x+2) ; i++ {
		for j := Max(0,y-1); j < Min(len(arr),y+2) ; j++ {
			if i != x || j != y {
				if arr[i][j] {
					count++
				}
			}
		}
	}

	return count
}

func Iterate(arr [][]bool) [][]bool {
	arr2 := MakeBoard(len(arr))
	for i := 0 ; i < len(arr) ; i++ {
		for j := 0; j < len(arr[i]); j++ {
			neighs := ComputeNeighbors(arr,i,j)
			if arr[i][j] {
				if neighs == 2 || neighs == 3 {
					arr2[i][j] = true
				}
			} else {
				if neighs == 3 {
					arr2[i][j] = true
				}
			}
		}
	}
	return arr2
}

func CountArr(arr [][]bool) int {
	count := 0
	for i := 0 ; i < len(arr) ; i++ {
		for j := 0 ; j < len(arr[i]) ; j++ {
			if arr[i][j] {
				count++
			}
		}
	}
	return count
}

func BuildArr(strs []string) [][]bool {
	arr := MakeBoard(len(strs[0]))

	for i := 0 ; i < len(strs) ; i++ {
		for j := 0 ; j < len(strs[i]) ; j++ {
			if strs[i][j] == '#'{
				arr[i][j] = true
			}
		}
	}

	return arr
}

func dayeighteen() {
	file,_ := os.Open("input-day18")
	scanner := bufio.NewScanner(file)
	arr := make([]string,0)
	
	for scanner.Scan() {
		text := scanner.Text()
		arr = append(arr, text)
	}
	initArr := BuildArr(arr)
	initArr2 := BuildArr(arr)
	initArr2[0][0] = true
	initArr2[0][len(initArr2)-1] = true
	initArr2[len(initArr2)-1][0] = true
	initArr2[len(initArr2)-1][len(initArr2)-1] = true
	
	for i:= 0 ; i < 100 ; i++ {
		initArr = Iterate(initArr)
		initArr2 = Iterate(initArr2)
		initArr2[0][0] = true
		initArr2[0][len(initArr2)-1] = true
		initArr2[len(initArr2)-1][0] = true
		initArr2[len(initArr2)-1][len(initArr2)-1] = true

	}
	fmt.Printf("Lights (100) = %v\n",CountArr(initArr))
	fmt.Printf("Light2 (100) = %v\n",CountArr(initArr2))
}
