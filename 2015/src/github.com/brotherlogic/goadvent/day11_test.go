package main

import "testing"

func TestDay11P1(t *testing.T) {

	if string(Next([]byte("a"))) != "b" {
		t.Errorf("Next a is %q\n", string(Next([]byte("a"))))
	}
	if string(Next([]byte("aa"))) != "ab" {
		t.Errorf("Next aa is %q\n", Next([]byte("aa")))
	}
	if string(Next([]byte("az"))) != "ba" {
		t.Errorf("Next az is %q\n", Next([]byte("az")))
	}
	if !FirstReq([]byte("hijklmmn")) {
		t.Errorf("hijklmn fails first\n")
	}
	if SecondReq([]byte("hijklmmn")) {
		t.Errorf("hijklmn fails second\n")
	}
	if !ThirdReq([]byte("abbceffg")) {
		t.Errorf("abbceffg fails third\n")
	}
	if FirstReq([]byte("abbceffg")) {
		t.Errorf("abbceffg fails first\n")
	}
	if ThirdReq([]byte("abbcegjk")) {
		t.Errorf("abbcegjk fails third\n")
	}
	if NewPassword("abcdefgh") != "abcdffaa" {
		t.Errorf("%q is not abcdffaa\n", NewPassword("abcdefgh"))
	}
	if NewPassword("ghijklmn") != "ghjaabcc" {
		t.Errorf("%q is not abcdffaa\n", NewPassword("abcdefg"))
	}

}
