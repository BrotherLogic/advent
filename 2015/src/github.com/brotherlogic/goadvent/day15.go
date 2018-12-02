package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"

func Mult(arr []int) int{
	count := 1
	for i := 0 ; i < len(arr) ; i++ {
		count *= arr[i]
	}
	return count
}

func ComputeArr(arr []int, mapper map[int][]int, cal bool) int {
	vals := make([]int,5)
	for i := 0 ; i < len(vals); i++ {
		for j := 0 ; j < len(arr) ; j++ {
			vals[i] += arr[j]*mapper[j][i]
		}
		vals[i] = Max(0,vals[i])
	}
	
	if !cal || vals[4] == 500 {
		return Mult(vals[:4])
	} else {
		return 0
	}
}

func ProcArrs(leng int, m map[int][]int, maxv int, cal bool) int{
	arr := make([]int, leng)
	best := 0
	for arr[0] != maxv {
		if Sum(arr) == maxv {
			val := ComputeArr(arr,m,cal)
			if val > best {
				best = val
			}
		}

		for i := len(arr) - 1 ; i >= 0 ; i-- {
			arr[i]++
			if arr[i] == maxv {
				arr[i] = 0
				if i == 0 {
					return best
				}
			} else {
				break
			}
		}
	}

	return best
}


func dayfifteen() {
	file,_ := os.Open("input-day15")
	scanner := bufio.NewScanner(file)

	count := 0
	maxv := 100

	var m map[int][]int
	m = make(map[int][]int)
	matcher := regexp.MustCompile(`(\w+): capacity (.*?), durability (.*?), flavor (.*?), texture (.*?), calories (.*)`)
	
	for scanner.Scan() {
		text := scanner.Text()
		match := matcher.FindAllStringSubmatch(text,-1)
		sarr := make([]int, len(match[0])-2)
		for i := 0 ; i < len(sarr) ; i++ {
			sarr[i],_ = strconv.Atoi(match[0][i+2])
		}
		m[count] = sarr
		count++;
	}

	fmt.Printf("Answer = %d\n", ProcArrs(count, m, maxv, false))
	fmt.Printf("Answer2 = %d\n", ProcArrs(count, m, maxv, true))
}
