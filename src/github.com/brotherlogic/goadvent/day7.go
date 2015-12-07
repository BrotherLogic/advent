package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strconv"
import "strings"


func WorkRules(rules map[string]string, key string) int{
	response := rules[key]

	if strings.Contains(response, "AND") {
		parts := strings.Split(response, " AND ")
		val1 := WorkRules(rules, parts[0])
		val2 := WorkRules(rules, parts[1])
		return val1 & val2
	}

	if strings.Contains(response, "OR") {
		parts := strings.Split(response, " OR ")
		val1 := WorkRules(rules, parts[0])
		val2 := WorkRules(rules, parts[1])
		return val1 | val2
	}
	
	if strings.Contains(response, "LSHIFT") {
		parts := strings.Split(response, " LSHIFT ")
		val1 := WorkRules(rules, parts[0])
		val2,_ := strconv.Atoi(parts[1])
		return val1 << uint8(val2)
	}

	if strings.Contains(response, "RSHIFT") {
		parts := strings.Split(response, " RSHIFT ")
		val1 := WorkRules(rules, parts[0])
		val2,_ := strconv.Atoi(parts[1])
		return val1 >> uint8(val2)
	}

	if strings.Contains(response, "OR") {
		parts := strings.Split(response, " OR ")
		val1 := WorkRules(rules, parts[0])
		val2 := WorkRules(rules, parts[1])
		return val1 | val2
	}
	
	if strings.Contains(response, "NOT") {
		parts := strings.Split(response, "NOT ")
		val1 := uint16(WorkRules(rules, parts[1]))
		return int(^val1)
	}

	m, err := regexp.MatchString(`\d+`, response)
	if  err == nil && m {
		conv, _ := strconv.Atoi(response)
		return conv
	}
	
	return 0
}


func dayseven() {
	re := regexp.MustCompile(`(.*) -> (.*)`)
	var rules map[string]string
	rules = make(map[string]string)
	
	file, err := os.Open("input-day7")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		result := re.FindAllStringSubmatch(text,-1)
		rules[result[0][2]] = result[0][1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("'a' = %d\n", WorkRules(rules, "a"))
}
