package main

import "encoding/json"
import "strings"
import "testing"

func TestDay12P1(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
		{"[1,2,3]", 6},
		{"{\"a\":2,\"b\":4}", 6},
		{"[[[3]]]", 3},
		{"{\"a\":{\"b\":4},\"c\":-1}", 3},
		{"{\"a\":[-1,1]}", 0},
		{"[-1,{\"a\":1}]", 0},
		{"[]", 0},
		{"{}", 0},
	}

	for _, c := range cases {
		answer := ProcJsonString(json.NewDecoder(strings.NewReader(c.in)))
		if answer != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, answer, c.want)
		}
	}
}

func TestDay12P2(t *testing.T) {

	cases := []struct {
		in string
		want int
	} {
//		{"[1,2,3]", 6},
		{"[1,{\"c\":\"red\",\"b\":2},3]", 4},
//		{"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 0},
//		{"[1,\"red\",5]", 6},
	}

	for _, c := range cases {
		answer := ProcJsonStringNoRed(json.NewDecoder(strings.NewReader(c.in)))
		if answer != c.want {
			t.Errorf("Spec(%q) == %d, want %d", c.in, answer, c.want)
		}
	}
}
