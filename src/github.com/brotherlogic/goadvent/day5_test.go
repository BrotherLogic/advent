package main

import "testing"

func TestDay5P2(t *testing.T) {

	cases := []struct {
		in string
		want bool
	} {
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", true},
	}

	for _, c := range cases {
		got := IsNiceAlso(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}


func TestDay5P1(t *testing.T) {

	cases := []struct {
		in string
		want bool
	} {
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, c := range cases {
		got := IsNice(c.in)
		if got != c.want {
			t.Errorf("Spec(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
