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
		got, err := ParseBinary(tt.binary)
		if err != nil {
			t.Fatalf("ParseBinary(%v) returned error %q.  Error not expected.",
				tt.binary, err)
		}

		if got != tt.want {
			t.Fatalf("ParseBinary(%v): got %d, want %v",
				tt.binary, got, tt.want)
		}
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
