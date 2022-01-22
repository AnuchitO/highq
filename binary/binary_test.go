package binary

import (
	"testing"
)

var testCases = []struct {
	binary string
	want   int
}{
	{"1", 1},
	{"10", 2},
	{"11", 3},
	{"100", 4},
	{"1001", 9},
	{"11010", 26},
	{"10001101000", 1128},
	{"0", 0},
}

func TestParseBinary(t *testing.T) {
	for _, tt := range testCases {
		got, _ := ParseBinary(tt.binary)

		if got != tt.want {
			t.Fatalf("ParseBinary(%v): got %d, want %v",
				tt.binary, got, tt.want)
		}
	}
}

func TestParseInvalidBinary(t *testing.T) {
	want := `"a" is not a vaid digit`

	_, err := ParseBinary("1a")

	if err.Error() != want {
		t.Error("should return error:", want, "but got:", err.Error())
	}
}

// Benchmark combined time for all tests
func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			ParseBinary(tt.binary)
		}
	}
}
