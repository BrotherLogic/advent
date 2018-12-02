package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func Product(arr []int) int {
	val := 1
	for i := 0 ; i < len(arr) ; i++ {
		val *= arr[i]
	}
	return val
}

func ComputeArrangement(mainarr [][]int) (int,int){
	return len(mainarr[0]), Product(mainarr[0])
}

func Search(sofar []int, numbers []int, pointer int, matchv int, totalc int) int {
	best1 := -1
	if len(sofar) == totalc {
		if Sum(sofar) == matchv {
			return Product(sofar)
		}
	} else {
		for i := pointer ; i < len(numbers) ; i++ {
			sofar = append(sofar,numbers[i])
			newv := Search(sofar, numbers, i+1, matchv, totalc)
			if newv > 0 {
				if best1 == -1 {
					best1 = newv
				} else {
					best1 = Min(best1,newv)
				}
			}
			sofar = sofar[0:len(sofar)-1]
		}
	}

	return best1
}

func Build(mainarr [][]int, numbers []int, pointer int, matchv int) (int,int) {
	best1 := 999999
	best2 := 999999
	if pointer >= len(numbers) {
		val1, val2 := ComputeArrangement(mainarr)
		if val1 < best1 {
			best1 = val1
			best2 = val2
		} else if val2 < best2 {
			best2 = val2
		}
	} else {
		for i := 0 ; i < len(mainarr) ; i++ {
			if (i == 0 || len(mainarr[i]) < len(mainarr[i-1])) && Sum(mainarr[i]) + numbers[pointer] <= matchv {
				mainarr[i] = append(mainarr[i],numbers[pointer])
				val1, val2 := Build(mainarr, numbers, pointer+1, matchv)
				if val1 < best1 {
					best1 = val1
					best2 = val2
				} else if val2 < best2 {
					best2 = val2
				}
				mainarr[i] = mainarr[i][0:len(mainarr[i])-1]
			}
		}
	}

	return best1,best2
}

func ProcNumbers(numbers []int, gap int) int{

	match := Sum(numbers)/gap
	arr := [][]int {}

	for i := 0 ; i < gap ; i++ {
		narr := make([]int, 0)
		arr = append(arr, narr)
	}
	
	_,val2 := Build(arr, numbers, 0, match)
	return val2
}

func SimpleProc(numbers []int, gap int) int {
	sofar := make([]int, 0)
	init := 1
	found := false
	match := Sum(numbers)/gap
	newv := -1
	for !found {
		tempv := Search(sofar, numbers, 0, match, init)
		if tempv > 0 && (newv == -1 || tempv < newv) {
			newv = tempv
			found = true
		}

		if init > 5 {
			found = true
		}
		
		init++
	}

	return newv
}
 
func daytwentyfour() {
	file,_ := os.Open("input-day24")
	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		intv,_ := strconv.Atoi(text)
		numbers = append(numbers,intv)
	}

	fmt.Printf("Result = %v\n",SimpleProc(numbers,3))
	fmt.Printf("Result = %v\n",SimpleProc(numbers,4))
}
