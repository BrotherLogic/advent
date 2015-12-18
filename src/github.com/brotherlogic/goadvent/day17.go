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

func AddThemUp(arr []int, enabled []int, curr int, val int) int {
	sum := 0
	
	if SumEnabled(arr, enabled) == val {
		sum++
	}

	for i := curr ; i < len(enabled) ; i++ {
		enabled[i] = 1
		sum += AddThemUp(arr, enabled, i+1, val)
		enabled[i] = 0
	}

	return sum
}

func ComputeAllPerms(arr []int, val int) int {
	curr := make([]int, len(arr))
	return AddThemUp(arr, curr, 0, val)
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
}
