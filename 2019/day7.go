package a2019

type processor struct {
	program []int
	ppoint  int
	input   []int
	ipoint  int
	rb      int
}

func (p *processor) run() (int, bool) {
	output, pp, pi, nrb := ProcessOpCode2(p.program, p.ppoint, p.input, p.ipoint, p.rb)
	p.ppoint = pp
	p.ipoint = pi
	p.rb = nrb
	return output, pp < 0 || p.program[pp] == 99
}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func rcopy(p []int) []int {
	ti := make([]int, len(p))
	copy(ti, p)
	return ti
}

func calcMax(program []int, phase []int) int {
	amp := 0
	amps := []processor{
		processor{program: rcopy(program), ppoint: 0, input: []int{phase[0], 0}, ipoint: 0},
		processor{program: rcopy(program), ppoint: 0, input: []int{phase[1]}, ipoint: 0},
		processor{program: rcopy(program), ppoint: 0, input: []int{phase[2]}, ipoint: 0},
		processor{program: rcopy(program), ppoint: 0, input: []int{phase[3]}, ipoint: 0},
		processor{program: rcopy(program), ppoint: 0, input: []int{phase[4]}, ipoint: 0},
	}

	output := 0
	halt := false
	for !halt || amp != 0 {
		output, halt = amps[amp].run()
		namp := (amp + 1) % len(amps)
		amps[namp].input = append(amps[namp].input, output)
		amp = namp
	}

	return output
}

//FindMax finds the max
func FindMax(program []int, phase1 []int) int {
	maxv := 0
	Perm(phase1, func(arr []int) {
		val := calcMax(program, arr)
		if val > maxv {
			maxv = val
		}
	})
	return maxv
}
