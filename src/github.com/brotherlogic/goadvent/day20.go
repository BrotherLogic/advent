package main

import "fmt"

func ComputePresents(val int) int {
	count := 0
	for j := val ; j > 0 ; j-- {
		if val % j == 0 {
			count += j*10
		}
	}
	return count
}

func ComputeElves(val int) int {
	i := 2097152
	for i > 0 {
		fmt.Printf("Val = %v = %v\n", i, ComputePresents(i))
		i--
	}
	return i
}

func FastCompute(val int) int {
	arr := make([]int, val)
	for i:= 1 ; i < len(arr) ; i++ {
		for j:= i ; j < len(arr) ; j+=i {
			arr[j] += i*10
		}
		if arr[i] >= val {
			return i
		}
	}

	return -1
}

func FastCompute2(val int) int {
	arr := make([]int, val)
	for i:= 1 ; i < len(arr) ; i++ {
		count := 0
		for j:= i ; j < len(arr) ; j+=i {
			arr[j] += i*11
			count++
			if count >= 50 {
				break
			}
		}
		if arr[i] >= val {
			return i
		}
	}

	return -1
}
 
func daytwenty() {
	input_value := 33100000

	fmt.Printf("Answer = %v\n", FastCompute(input_value))
	fmt.Printf("Answer = %v\n", FastCompute2(input_value))
}
