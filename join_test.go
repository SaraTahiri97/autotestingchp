package autotestingchp

import (
	"testing"
)

func TestOneElements(t *testing.T) {
	list := []string{"apple"}
	if JoinWithCommas(list) != "apple" {
		t.Error("From TestOneElements : Did  not match expected value . the result : ", JoinWithCommas(list))

	}
}
func TestTwoElements(t *testing.T) {
	list := []string{"apple", " orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("From TestTwoElements : Did  not match expected value . the result : ", JoinWithCommas(list))

	}
}
func TestMultipleElements(t *testing.T) {
	list := []string{"apple", " orange", "pear"}
	if JoinWithCommas(list) != "apple, orange, and pear" {
		t.Error("From TestMultipleElements : Did  not match expected value . the result :", JoinWithCommas(list))

	}
}
