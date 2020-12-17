package a2020

import "testing"

func TestProgram(t *testing.T) {
	lines := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	val := RunMask(lines)
	if val != 165 {
		t.Errorf("Bad mask: %v", val)
	}

	lines2 := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
	val2 := RunMask2(lines2)
	if val2 != 208 {
		t.Errorf("Bad run2: %v", val2)
	}
}

func TestMask(t *testing.T) {

	v := "000000000000000000000000000000001011"
	m := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	e := "000000000000000000000000000001001001"

	nv := mask(m, v)
	if nv != e {
		t.Errorf("Bad mask: %v", nv)
	}

	dv := convert(v)
	if dv != 11 {
		t.Errorf("Bad convert: %v", dv)
	}

	if deconvert(11) != v {
		t.Errorf("Bad deconvert: %v -> %v", deconvert(11), len(deconvert(11)))
		t.Errorf("Bad deconvert: %v -> %v", v, len(v))
	}
}
