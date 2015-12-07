package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"


func WorkRules(rules map[string]string, key string) int{
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
