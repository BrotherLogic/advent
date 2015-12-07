package main

import "testing"

func TestDay7P1(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{"d",0},
		{"e",0},
		{"f",0},
		{"g",0},
		{"h",0},
		{"i",0},
		{"x",0},
		{"y",0},
	}

	var rules map[string]string
	rules = make(map[string]string)

	rules["x"] = "123"
	rules["y"] = "456"
	rules["d"] = "x AND y"
	rules["e"] = "x OR y"
	rules["f"] = "x LSHIFT 2"
	rules["g"] = "y RSHIFT 2"
	rules["h"] = "NOT x"
	rules["i"] = "NOT y"

	for _, c := range cases {
		answer := WorkRules(rules, c.in)
		
		if answer != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, answer, c.want)
		}
	}
}
