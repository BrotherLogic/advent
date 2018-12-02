package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"

func MatchSue2(sue map[string]int, maybe map[string]int) bool {
	for key, val := range maybe {
		if (key == "cats" || key == "trees") && val <= sue[key] {
			return false
		} else if (key == "pomeranians" || key == "goldfish") && val >= sue[key] {
			return false
		} else  if (key != "pomeranians" && key != "goldfish" && key != "cats" && key != "trees") &&  sue[key] != val {
			return false
		}

	}
	fmt.Printf("ACCEPT %v\n", maybe)
	return true;
	
}

func MatchSue1(sue map[string]int, maybe map[string]int) bool {
	for key, val := range maybe {
		if sue[key] != val {
			return false
		}
	}
	return true;
}

func daysixteen() {
	file,_ := os.Open("input-day16")
	scanner := bufio.NewScanner(file)

	var sue map[string]int
	sue = make(map[string]int)
	sue["children"] = 3
	sue["cats:"] =  7
	sue["samoyeds"] =  2
	sue["pomeranians"] = 3
	sue["akitas"] =  0
	sue["vizslas"] = 0
	sue["goldfish"] = 5
	sue["trees"] = 3
	sue["cars"] = 2
	sue["perfumes"] = 1
	
	matcher := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

	found := -1
	found2 := -1
	for scanner.Scan() {
		text := scanner.Text()
		match := matcher.FindAllStringSubmatch(text,-1)

		var pot map[string]int
		pot = make(map[string]int)
		val1,_ := strconv.Atoi(match[0][3])
		val2,_ := strconv.Atoi(match[0][5])
		val3,_ := strconv.Atoi(match[0][7]) 
		pot[match[0][2]] = val1
		pot[match[0][4]] = val2
		pot[match[0][6]] = val3
		suev,_ := strconv.Atoi(match[0][1])

		if MatchSue1(sue,pot) {
			found = suev
		}

		if MatchSue2(sue,pot) {
			found2 = suev
		}
		
	}

	fmt.Printf("Answer = %d\n", found)
	fmt.Printf("Answer = %d\n", found2)
}
