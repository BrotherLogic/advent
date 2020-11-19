package a2019

//RunFull runs a full program
func RunFull(programIn []int) int {
	p := processor{
		program: programIn,
		input:   []int{1},
	}

	for i := 0; i < 10000; i++ {
		p.program = append(p.program, 0)
	}
	output, _ := p.run()
	return output
}

//RunFull2 runs a full program
func RunFull2(programIn []int) int {
	p := processor{
		program: programIn,
		input:   []int{2},
	}

	for i := 0; i < 10000; i++ {
		p.program = append(p.program, 0)
	}
	output, _ := p.run()
	return output
}
