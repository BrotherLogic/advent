package main

import "testing"

func TestDay10P1(t *testing.T) {

	cases := []struct {
		in string
		want string
	} {
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, c := range cases {
		answer := LookAndSay(c.in)		
		if answer != c.want {
			t.Errorf("Spec(%q) == %q, want %q", c.in, answer, c.want)
		}
	}

	last := LookApply("1", 5)
	if last != "312211" {
		t.Errorf("Spec(%q) == %q, want %q", "1", last, "312211")
	}
}
