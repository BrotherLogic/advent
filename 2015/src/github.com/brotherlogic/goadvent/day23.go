package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func ProcInstruction(text string, registers map[string]int) (map[string]int,int) {

	if strings.HasPrefix(text,"hlf") {
		register := text[len(text)-1:len(text)]
		registers[register] /= 2
	} else if strings.HasPrefix(text,"tpl") {
		register := text[len(text)-1:len(text)]
		registers[register] *= 3
	} else if strings.HasPrefix(text,"inc") {
		register := text[len(text)-1:len(text)]
		registers[register] += 1
	} else if strings.HasPrefix(text,"jmp") {
		pointer := Max(strings.Index(text,"-"),strings.Index(text,"+"))
		increment,_ := strconv.Atoi(text[pointer+1:len(text)])
		if text[pointer:pointer+1] == "-" {
			increment *= -1
		}
		return registers,increment
		
	} else if strings.HasPrefix(text,"jie") {
		pointer := Max(strings.Index(text,"-"),strings.Index(text,"+"))
		increment,_ := strconv.Atoi(text[pointer+1:len(text)])
		if text[pointer:pointer+1] == "-" {
			increment *= -1
		}
		register := text[4:5]
		if registers[register] % 2 == 0 {
			return registers,increment
		}
	} else if strings.HasPrefix(text,"jio") {
		pointer := Max(strings.Index(text,"-"),strings.Index(text,"+"))
		increment,_ := strconv.Atoi(text[pointer+1:len(text)])
		if text[pointer:pointer+1] == "-" {
			increment *= -1
		}
		register := text[4:5]
		if registers[register] == 1 {
			return registers,increment
		}
	} else{
		fmt.Printf("PROC %v\n", text)
	}
	return registers,1
}

 
func daytwentythree() {
	file,_ := os.Open("input-day23")
	scanner := bufio.NewScanner(file)

	var registers map[string]int
	registers = make(map[string]int)
	registers["a"] = 0
	registers["b"] = 0
	instructions := make([]string,0)
	for scanner.Scan() {
		text := scanner.Text()
		instructions = append(instructions,text)
	}

	curr_instruction := 0
	adjust := 0
	for curr_instruction < len(instructions) {
		registers,adjust = ProcInstruction(instructions[curr_instruction], registers)
		curr_instruction += adjust
	}

	fmt.Printf("Part 1 = %v\n", registers["b"])

	registers["a"] = 1
	registers["b"] = 0
	curr_instruction = 0
	adjust = 0
	for curr_instruction < len(instructions) {
		registers,adjust = ProcInstruction(instructions[curr_instruction], registers)
		curr_instruction += adjust
	}
	fmt.Printf("Part 2 = %v\n", registers["b"])
}
