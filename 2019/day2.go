package a2019

func runOpCode1(program []int) int {
	pointer := 0
	for program[pointer] != 99 {
		if program[pointer] == 1 {
			program[program[pointer+3]] = program[program[pointer+1]] + program[program[pointer+2]]
		} else if program[pointer] == 2 {
			program[program[pointer+3]] = program[program[pointer+1]] * program[program[pointer+2]]
		}

		pointer += 4
	}
	return program[0]
}
