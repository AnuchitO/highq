package wordscore

import "testing"

func TestScore(t *testing.T) {
	for _, test := range tests {
		if actual := Score(test.input); actual != test.want {
			t.Errorf("Score(%q) want %d, Actual %d", test.input, test.want, actual)
		}
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Score(test.input)
		}
	}
}

var tests = []struct {
	input string
	want  int
}{
	{"", 0},
	{" \t\n", 0},
	{"a", 1},
	{"f", 4},
	{"street", 6},
	{"quirky", 22},
	{"oxyphenbutazone", 41},
	{"alacrity", 13},
}
