package a2020

import "testing"

type test182020 struct {
	input   string
	result  int
	result2 int
}

var tests182020 = []test182020{
	{input: "1 + 2 * 3 + 4 * 5 + 6", result: 71, result2: 231},
	{input: "1 + (2 * 3) + (4 * (5 + 6))", result: 51, result2: 51},
	{input: "2 * 3 + (4 * 5)", result: 26, result2: 46},
	{input: "5 + (8 * 3 + 9 + 3 * 4 * 3)", result: 437, result2: 1445},
	{input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", result: 12240, result2: 669060},
	{input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", result: 13632, result2: 23340},
}

func TestSum(t *testing.T) {
	for _, te := range tests182020 {
		res := runSingleSum(te.input, map[string]int{"*": 1, "+": 1, "(": 0, ")": 0})
		if res != te.result {
			t.Fatalf("Bad results1 %v -> %v, should have been %v", te.input, res, te.result)
		}
		res2 := runSingleSum(te.input, map[string]int{"*": 1, "+": 2, "(": 0, ")": 0})
		if res2 != te.result2 {
			t.Fatalf("Bad results2 %v -> %v, should have been %v", te.input, res2, te.result2)
		}
	}
}
