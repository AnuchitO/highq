package ticket

import "testing"

func TestPriceFree(t *testing.T) {
	p := Price(Order{Age: 4})
	if p != 0 {
		t.Error("Expected", 0, "Got", p)
	}
}
