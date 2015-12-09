package main

import "bufio"
import "fmt"
import "log"
import "os"

type MapEntry struct {
	remaining[] int
	children[] MapEntry
	distance int
	location int
}

func Sum(arr[] int) int{
	sum := 0
	for i := 0 ; i < len(arr) ; i++ {
		sum += arr[i]
	}
	return sum
}

func SearchTree(root MapEntry) int {
	fmt.Printf("NUM = %v\n", root)
	if len(root.children) == 0 {
		return root.distance
	} else {
		best := SearchTree(root.children[0])
		for i := 1 ; i < len(root.children) ; i++ {
			best = Min(best, SearchTree(root.children[i]))
		}
		return best
	}
}

func BuildTree(root MapEntry, arr [][]int) MapEntry {
	fmt.Printf("BUILDING = %v with %v\n", root, arr)
	if root.location == -1 {
		root.children = make([]MapEntry, len(arr))
		for i := 0 ; i < len(arr) ; i++ {
			var child MapEntry
			child.remaining = make([]int, len(arr))
			child.remaining[i] = 1
			child.distance = 0
			child.location = i
			child = BuildTree(child, arr)
			root.children[i] = child
		}
	} else {
		left := len(arr)-Sum(root.remaining)
		if left > 0 {
			root.children = make([]MapEntry, left)
			child_pointer := 0
			for i := 0 ; i < len(root.remaining) ; i++ {
				if root.remaining[i] == 0 {
					var child MapEntry
					child.remaining = make([]int, len(root.remaining))
					for j := 0 ; j < len(child.remaining) ; j++ {
						child.remaining[j] = root.remaining[j]
					}
					child.remaining[i] = 1
					child.location = i
					child.distance = root.distance + arr[root.location][child.location]
					child = BuildTree(child, arr)
					root.children[child_pointer] = child
					child_pointer++
				}
			}
		}
	}

	return root
}

func SearchMap(arr [][]int) int {
	// Build out a tree structure
	var p MapEntry
	p.remaining = make([]int, len(arr))
	p.distance = 0
	p.location = -1

	p = BuildTree(p,arr)	
	return SearchTree(p)
}

func daynine() {
	
	file, err := os.Open("input-day9")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("%q\n",text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
