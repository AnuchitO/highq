package sum

import "testing"

func TestSum(t *testing.T) {
	r := Sum(2, 3)
	if r != 5 {
		t.Error("Expected", 5, "Got", r)
	}
}

func TestSumAll(t *testing.T) {
	r := Sum([]int{2, 3, 3}...)
	if r != 8 {
		t.Error("Expected", 8, "Got", r)
	}
}
