package a2020

import (
	"fmt"
	"strconv"
	"strings"
)

//Program a piece oof code
type Program struct {
	instructions []string
	seen         []bool
	pointer      int
	accumulator  int
}

func (p *Program) runToLoop() error {
	for p.pointer < len(p.instructions) {
		if p.seen[p.pointer] {
			return fmt.Errorf("Looped")
		}
		p.seen[p.pointer] = true
		fields := strings.Fields(p.instructions[p.pointer])

		num32, _ := strconv.ParseInt(string(fields[1][1:]), 10, 32)
		num := int(num32)
		if fields[1][0] == '-' {
			num = 0 - num
		}

		switch fields[0] {
		case "nop":
			p.pointer++
		case "jmp":
			p.pointer += num
		case "acc":
			p.pointer++
			p.accumulator += num
		}

	}

	return nil
}

//FixAndRunCode the code and runs it
func FixAndRunCode(code []string) int {
	for i := 0; i < len(code); i++ {
		if !strings.HasPrefix(code[i], "acc") {
			cp := make([]string, len(code))
			copy(cp, code)

			prefix := "jmp"
			if strings.HasPrefix(code[i], "jmp") {
				prefix = "nop"
			}
			cp[i] = prefix + string(cp[i][3:])

			pr := &Program{
				instructions: cp,
				seen:         make([]bool, len(code)),
			}

			if pr.runToLoop() == nil {
				return pr.accumulator
			}
		}
	}

	return -1
}

//RunCode runs a program
func RunCode(code []string) int {
	pr := &Program{
		instructions: code,
		seen:         make([]bool, len(code)),
	}
	pr.runToLoop()
	return pr.accumulator
}
