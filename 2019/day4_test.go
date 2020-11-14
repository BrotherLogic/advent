package a2019

import "testing"

func TestValid(t *testing.T) {
	if !Valid(111111) {
		t.Errorf("Bad valid on 111111")
	}

	if Valid(223450) {
		t.Errorf("Bad valid on 223450")
	}

	if Valid(123789) {
		t.Errorf("Bad valid on 123789")
	}
}

func TestTightValid(t *testing.T) {
	if !TightValid(112233) {
		t.Errorf("Bad valid on 112233")
	}

	if TightValid(123444) {
		t.Errorf("Bad valid on 123444")
	}

	if !TightValid(111122) {
		t.Errorf("Bad valid on 111122")
	}
}
