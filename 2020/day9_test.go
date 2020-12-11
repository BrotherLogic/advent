package a2020

import "testing"

func TestCalc(t *testing.T) {
	nums := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	res := FindFirst(nums, 5)
	if res != 127 {
		t.Errorf("Bad nums: %v", res)
	}

	cmin, cmax := FindContiguous(nums, res)
	if cmin+cmax != 62 {
		t.Errorf("Bad range: %v, %v", cmin, cmax)
	}

}

func TestValids(t *testing.T) {
	nums := []int64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}

	valids := []int64{26, 49}
	for _, v := range valids {
		res := FindFirst(append(nums, v), 25)
		if res != 0 {
			t.Errorf("Bad nums: %v", res)
		}
	}

	invalids := []int64{100, 50}
	for _, iv := range invalids {
		res := FindFirst(append(nums, iv), 25)
		if res != iv {
			t.Errorf("Bad nums: %v", res)
		}
	}
}

func TestValids2(t *testing.T) {
	nums := []int64{
		20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 45,
	}

	valids := []int64{26, 64, 66}
	for _, v := range valids {
		res := FindFirst(append(nums, v), 25)
		if res != 0 {
			t.Errorf("Bad nums: %v", v)
		}
	}

	invalids := []int64{65}
	for _, iv := range invalids {
		res := FindFirst(append(nums, iv), 25)
		if res != iv {
			t.Errorf("Bad nums: %v", res)
		}
	}
}
