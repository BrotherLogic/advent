package a2020

import "testing"

func TestReport(t *testing.T) {
	nums := []int{1721,
		979,
		366,
		299,
		675,
		1456}

	if ComputeExpenseReport(nums) != 514579 {
		t.Errorf("Bad value: %v", ComputeExpenseReport(nums))
	}

	if ComputeExpenseReportThree(nums) != 241861950 {
		t.Errorf("Bad value: %v", ComputeExpenseReportThree(nums))
	}
}
