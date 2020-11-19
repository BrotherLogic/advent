package a2019

import (
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

func process(p, val int, program []int, rb int) int {
	if p == 0 {
		return program[program[val]]
	} else if p == 2 {
		return program[rb+program[val]]
	}
	return program[val]
}

func runCode(pointer int, iPoint int, program []int, input []int, rb int) (int, int, int, int) {
	s1, s2, s3, opcode := extract(program[pointer])

	switch opcode {
	case 1:
		//Basic add
		if s1 == 2 {
			program[rb+program[pointer+3]] = process(s3, pointer+1, program, rb) + process(s2, pointer+2, program, rb)
		} else {
			program[program[pointer+3]] = process(s3, pointer+1, program, rb) + process(s2, pointer+2, program, rb)
		}
		return pointer + 4, 0, iPoint, rb
	case 2:
		//Basic multiply
		if s1 == 2 {
			program[rb+program[pointer+3]] = process(s3, pointer+1, program, rb) * process(s2, pointer+2, program, rb)
		} else {
			program[program[pointer+3]] = process(s3, pointer+1, program, rb) * process(s2, pointer+2, program, rb)
		}
		return pointer + 4, 0, iPoint, rb
	case 3:
		//Halt if we have no input
		if iPoint >= len(input) {
			return -1, 0, iPoint, rb
		}

		// Input command
		if s3 == 2 {
			program[rb+program[pointer+1]] = input[iPoint]
		} else {
			program[program[pointer+1]] = input[iPoint]
		}
		return pointer + 2, 0, iPoint + 1, rb
	case 4:
		output := process(s3, pointer+1, program, rb)
		return pointer + 2, output, iPoint, rb
	case 5:
		test := process(s3, pointer+1, program, rb)
		if test > 0 {
			return process(s2, pointer+2, program, rb), 0, iPoint, rb
		}
		return pointer + 3, 0, iPoint, rb
	case 6:
		test := process(s3, pointer+1, program, rb)
		if test == 0 {
			return process(s2, pointer+2, program, rb), 0, iPoint, rb
		}
		return pointer + 3, 0, iPoint, rb
	case 7:
		t1 := process(s3, pointer+1, program, rb)
		t2 := process(s2, pointer+2, program, rb)
		if t1 < t2 {
			if s1 == 2 {
				program[rb+program[pointer+3]] = 1
			} else {
				program[program[pointer+3]] = 1
			}
		} else {
			if s1 == 2 {
				program[rb+program[pointer+3]] = 0
			} else {
				program[program[pointer+3]] = 0
			}
		}
		return pointer + 4, 0, iPoint, rb
	case 8:
		t1 := process(s3, pointer+1, program, rb)
		t2 := process(s2, pointer+2, program, rb)
		if t1 == t2 {
			if s1 == 2 {
				program[rb+program[pointer+3]] = 1
			} else {
				program[program[pointer+3]] = 1
			}
		} else {
			if s1 == 2 {
				program[rb+program[pointer+3]] = 0
			} else {
				program[program[pointer+3]] = 0
			}
		}
		return pointer + 4, 0, iPoint, rb
	case 9:
		if program[pointer] == 99 {
			return -1, 0, iPoint, rb
		}

		return pointer + 2, 0, iPoint, rb + process(s3, pointer+1, program, rb)
	}

	return 0, 0, 0, iPoint
}

// ProcessOpCode2 runs the new op code thingy
func ProcessOpCode2(program []int, pointer int, input []int, ipoint int, rb int) (int, int, int, int) {
	output := 0
	nrb := rb
	for pointer >= 0 {
		npointer, noutput, ninput, nrb := runCode(pointer, ipoint, program, input, rb)
		ipoint = ninput
		rb = nrb
		if noutput != 0 {
			output = noutput
		}

		//We've halted here - might be a halt or a request for input
		if npointer < 0 {
			break
		}
		pointer = npointer
	}

	return output, pointer, ipoint, nrb
}
