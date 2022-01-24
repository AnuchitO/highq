package fixtures

import (
	"os"
	"testing"
)

func TestFigtures(t *testing.T) {
	f, err := os.Open("testdata/big.json")
	t.Log(f.Name(), err)
}
