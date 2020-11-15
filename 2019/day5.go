package a2019

import (
	"log"
	"strconv"
)

func digitize(code, lent int) []int {
	str := strconv.Itoa(code)
	for i := len(str); i < lent; i++ {
		str = "0" + str
	}

	digits := []int{}
	for _, d := range str {
		val, _ := strconv.ParseInt(string(d), 10, 32)
		digits = append(digits, int(val))
	}

	return digits
}

func extract(code int) (int, int, int, int) {
	digits := digitize(code, 5)

	opcode := digits[4]
	s3 := digits[2]
	s2 := digits[1]
	s1 := digits[0]

	return s1, s2, s3, opcode
}

func process(p, val int, program []int) int {
	if p == 0 {
		return program[program[val]]
	}
	return program[val]
}

func runCode(pointer int, program []int, input int) (int, int) {
	s1, s2, s3, opcode := extract(program[pointer])

	switch opcode {
	case 1:
		//Basic add
		program[program[pointer+3]] = process(s3, pointer+1, program) + process(s2, pointer+2, program)
		return pointer + 4, 0
	case 2:
		//Basic multiply
		program[program[pointer+3]] = process(s3, pointer+1, program) * process(s2, pointer+2, program)
		return pointer + 4, 0
	case 3:
		// Input command
		program[program[pointer+1]] = input
		return pointer + 2, 0
	case 4:
		output := process(s1, pointer+1, program)
		return pointer + 2, output
	case 5:
		test := process(s3, pointer+1, program)
		if test > 0 {
			return process(s2, pointer+2, program), 0
		}
		return pointer + 3, 0
	case 6:
		test := process(s3, pointer+1, program)
		if test == 0 {
			return process(s2, pointer+2, program), 0
		}
		return pointer + 3, 0
	case 7:
		t1 := process(s3, pointer+1, program)
		t2 := process(s2, pointer+2, program)
		if t1 < t2 {
			program[program[pointer+3]] = 1
		} else {
			program[program[pointer+3]] = 0
		}
		return pointer + 4, 0
	case 8:
		t1 := process(s3, pointer+1, program)
		t2 := process(s2, pointer+2, program)
		if t1 == t2 {
			program[program[pointer+3]] = 1
		} else {
			program[program[pointer+3]] = 0
		}
		return pointer + 4, 0
	case 9:
		return -1, -1
	}

	log.Fatalf("Unknown opcode: %v (from %v)", opcode, program[pointer])
	return 0, pointer
}

// ProcessOpCode2 runs the new op code thingy
func ProcessOpCode2(program []int, input int) int {
	pointer := 0
	output := 0
	for pointer >= 0 {
		npointer, noutput := runCode(pointer, program, input)
		pointer = npointer
		if noutput > 0 {
			output = noutput
		}

	}

	return output
}
