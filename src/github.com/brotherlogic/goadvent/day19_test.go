package main

import "testing"

func TestDay19P1(t *testing.T) {
	var tester map[string][]string
	tester = make(map[string][]string)
	tester["H"] = make([]string,0)
	tester["H"] = append(tester["H"],"HO")
	tester["H"] = append(tester["H"],"OH")
	tester["O"] = make([]string,0)
	tester["O"] = append(tester["O"],"HH")

	mapping := BuildMapping(tester, 0, "HOH")
	if len(mapping) != 4 {
		t.Errorf("Poor mapping %v\n", mapping)
	}
	
	mapping2 := BuildMapping(tester, 0, "HOHOHO")
	if len(mapping2) != 7 {
		t.Errorf("Poor second mapping %v,%v\n", len(mapping2), mapping2)
	}
}
