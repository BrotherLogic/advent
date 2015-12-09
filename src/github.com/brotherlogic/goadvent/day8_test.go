package main

import "testing"

func TestDay8P1(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{"\"\"",2-0},
		{"\"abc\"", 5-3},
		{"\"aaa\\\"aaa\"",10-7},
		{"\"\\x27\"",6-1},
		{"\"\\\\x\"",3},
		{"\"\\xat\"",6},
		{"\"\\xaf\"",5},
		{"\"rw\\\"\\\"wwn\\\\fgbjumq\\\"vgvoh\\xd0\\\"mm\"",10},
	}

	for _, c := range cases {
		answer := Convert(c.in)
		
		if answer != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, answer, c.want)
		}
	}
}

func TestDay8P2(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{"\"\"", 6-2},
		{"\"abc\"",9-5},
		{"\"aaa\\\"aaa\"",16-10},
		{"\"\\x27\"",11-6},
	}

	for _, c := range cases {
		answer := Convert2(c.in)
		
		if answer != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, answer, c.want)
		}
	}
}

