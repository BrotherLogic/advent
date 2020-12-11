package a2020

import "testing"

func TestBasic(t *testing.T) {
	lines := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	res := FindDiff(lines, 1, 3)
	if res != 7*5 {
		t.Errorf("Bad result: %v", res)
	}

	travel := Travel(lines)
	if travel != 8 {
		t.Errorf("Bad result: %v", travel)
	}
}

func TestTravel(t *testing.T) {
	nums := []int{28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	val := Travel(nums)
	if val != 19208 {
		t.Errorf("BAD: %v", val)
	}
}
