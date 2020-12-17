package a2020

import (
	"strconv"
	"strings"
)

func convert(d string) int {
	val := 0
	mult := 1
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] == '1' {
			val += mult
		}
		mult *= 2
	}

	return val
}

func convert2(d string, p int, m map[int]int64, val int64) {
	if p < len(d) {
		if d[p] == 'X' {
			convert2(d[:p]+"0"+d[p+1:], p+1, m, val)
			convert2(d[:p]+"1"+d[p+1:], p+1, m, val)
		} else {
			convert2(d, p+1, m, val)
		}
	}

	m[convert(d)] = val
}

func deconvert(v int) string {
	str := strconv.FormatInt(int64(v), 2)
	for i := 36 - len(str); i > 0; i-- {
		str = "0" + str
	}
	return str
}

func mask(m, val string) string {
	nstr := ""
	for i, r := range m {
		if r == '0' && val[i] == '1' {
			nstr = nstr + "0"
		} else if r == '1' && val[i] == '0' {
			nstr = nstr + "1"
		} else {
			nstr = nstr + string(val[i])
		}
	}

	return nstr
}

func maskv2(m, val string) string {
	nstr := ""
	for i, r := range m {
		if r == '1' {
			nstr = nstr + "1"
		} else if r == 'X' {
			nstr = nstr + "X"
		} else {
			nstr = nstr + string(val[i])
		}
	}
	return nstr
}

//RunMask runs the mask program
func RunMask(lines []string) int {
	m := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	memmap := make(map[int]int)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			m = line[7:]
		} else {
			elems := strings.Fields(line)
			loc, _ := strconv.ParseInt(elems[0][4:len(elems[0])-1], 10, 32)
			val, _ := strconv.ParseInt(elems[2], 10, 32)
			memmap[int(loc)] = convert(mask(m, deconvert(int(val))))
		}
	}

	sv := 0
	for _, val := range memmap {
		sv += val
	}
	return sv
}

//RunMask2 runs the mask program
func RunMask2(lines []string) int64 {
	m := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	memmap := make(map[int]int64)

	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			m = line[7:]
		} else {
			elems := strings.Fields(line)
			loc, _ := strconv.ParseInt(elems[0][4:len(elems[0])-1], 10, 32)
			val, _ := strconv.ParseInt(elems[2], 10, 32)
			//nv := convert(mask(m, deconvert(int(val))))
			convert2(maskv2(m, deconvert(int(loc))), 0, memmap, int64(val))
		}
	}

	sv := int64(0)
	for _, val := range memmap {
		sv += val
	}
	return sv
}
