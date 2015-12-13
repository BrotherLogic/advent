package main

import "bufio"
import "container/list"
import "fmt"
import "math"
import "os"
import "regexp"
import "strconv"

func Compute(tocheck []int, valarr [][]int) int {
	count := 0
	for i := 1 ; i < len(tocheck) ; i++ {
		count += valarr[tocheck[i]][tocheck[i-1]] + valarr[tocheck[i-1]][tocheck[i]]
	}

	count += valarr[tocheck[0]][tocheck[len(tocheck)-1]] + valarr[tocheck[len(tocheck)-1]][tocheck[0]]
	fmt.Printf("Checking %v = %v\n", tocheck, count)
	return count
}

func MinPermutation(tobuild []int, toadd []int, valarr [][]int) int {
	if len(toadd) == 0 {
		return Compute(tobuild, valarr)
	}
	best := 0
	for i:= 0; i < len(toadd) ; i++ {
		tobuild[len(tobuild)-len(toadd)] = toadd[i]
		newarr := make([]int, 0)
		newarr = append(newarr, toadd[:i]...)
		newarr = append(newarr, toadd[i+1:]...)
		nval := MinPermutation(tobuild, newarr, valarr)
		
		if nval > best {
			best = nval
		}
	}
	return best
}

func ProcValues(arr [][]int) int {
	tobuild := make([]int, len(arr))
	toadd := make([]int, len(arr))
	for i := 0 ; i < len(toadd) ; i++ {
		toadd[i] = i
	}
	return MinPermutation(tobuild, toadd, arr)
}

func ProcSeatingArray(arr []string) int {
	num := int((1 + math.Sqrt(float64(1+4*len(arr))))/2)
	narr := [][]int {}
	for i := 0 ; i < num ; i++ {
		row := make([]int, num)
		narr = append(narr, row)
	}

	count := 0
	var m map[string]int
	m = make(map[string]int)

	matcher := regexp.MustCompile(`(\w*) would (\w*) (\d*) happiness units by sitting next to (\w*).`)
	
	for i := 0; i < len(arr); i++ {
		match := matcher.FindAllStringSubmatch(arr[i],-1)

		if _, ok := m[match[0][1]]; !ok {
			m[match[0][1]] = count
			count++
		}
		if _, ok := m[match[0][4]]; !ok {
			m[match[0][4]] = count
			count++
		}

		multiplier := 1
		if match[0][2] == "lose" {
			multiplier = -1
		}
		val,_ :=  strconv.Atoi(match[0][3])
		narr[m[match[0][1]]][m[match[0][4]]] = val*multiplier
		
	}
	
	return ProcValues(narr)
}

func daythirteen() {
	file,_ := os.Open("input-day13")
	list := list.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		list.PushFront(text)
	}
	arr := make([]string, list.Len())
	count := 0
	for e := list.Front(); e != nil; e = e.Next() {
		arr[count] = e.Value.(string)
		count++
	}
	fmt.Printf("Best = %d\n", ProcSeatingArray(arr))
}
