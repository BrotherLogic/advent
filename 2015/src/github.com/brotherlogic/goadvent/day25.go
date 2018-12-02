package main

import "fmt"

func TriangleIndex(val int) int {
	sum := 0
	for i := val ; i > 0 ; i-- {
		sum += i
	}
	return sum
}

func GetIndex(row int, col int) int {

	offset := row-1
	triangle := TriangleIndex(col+offset)
	return triangle-offset
}

func FindCode(startv int, row int, col int) int{
	//Convert row, col into number
	index := GetIndex(row,col)

	currv := startv
	for i:= 0 ; i < index-1 ; i++ {
		currv *= 252533
		currv = currv % 33554393
	}
	return currv
}

func daytwentyfive() {
	fmt.Printf("Result = %v\n",FindCode(20151125, 2981, 3075))
}
