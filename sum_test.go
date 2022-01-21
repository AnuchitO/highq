package main

import "testing"

func TestSum(t *testing.T) {
	r := Sum(2, 3)
	if r != 5 {
		t.Error("Expected", 5, "Got", r)
	}
}