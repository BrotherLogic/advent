package main

import "testing"

func TestDay5P1(t *testing.T) {

	cases := []struct {
		in string
		want bool
	} {
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", true},
		{"haegwjzuvuyypxyu", true},
		{"dvszwmarrgswjxmb", true},
	}

	for _, c := range cases {
		got := IsNice(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
