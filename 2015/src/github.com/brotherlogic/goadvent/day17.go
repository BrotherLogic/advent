package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func SumEnabled(arr []int, enabled []int) int {
	sum := 0
	for i := 0 ; i < len(arr) ; i++ {
		sum += arr[i]*enabled[i]
	}
	return sum
}

func AddThemUp(arr []int, enabled []int, curr int, val int) map[int]int {
	var mapper map[int]int
	mapper = make(map[int]int)
	
	if SumEnabled(arr, enabled) == val {
		mapper[Sum(enabled)] += 1
	}

	for i := curr ; i < len(enabled) ; i++ {
		enabled[i] = 1
		nmap := AddThemUp(arr, enabled, i+1, val)
		for key,value := range nmap {
			mapper[key] += value
		}
		enabled[i] = 0
	}

	return mapper
}

func ComputeAllPerms(arr []int, val int) int {
	curr := make([]int, len(arr))
	mapper := AddThemUp(arr, curr, 0, val)
	sum := 0
	for _,value := range mapper {
		sum += value
	}
	return sum
}

func FindTheMin(arr []int, val int) int {
	curr := make([]int, len(arr))
	mapper := AddThemUp(arr, curr, 0, val)
	min_key := 900
	min_val := 0
	for key,value := range mapper {
		if key < min_key && value > 0 {
			min_key = key
			min_val = value
		}
	}
	return min_val
}


func dayseventeen() {
	file,_ := os.Open("input-day17")
	scanner := bufio.NewScanner(file)
	arr := make([]int,0)
	
	for scanner.Scan() {
		text := scanner.Text()
		val,_ := strconv.Atoi(text)
		arr = append(arr, val)
	}
	fmt.Printf("Answer = %v\n", ComputeAllPerms(arr,150))
	fmt.Printf("Answe2 = %v\n", FindTheMin(arr,150))
}
