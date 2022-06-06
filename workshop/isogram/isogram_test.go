package isogram

import "testing"

var testCases = []struct {
	name  string
	input string
	want  bool
}{
	{
		name:  "empty string",
		input: "",
		want:  true,
	},
	{
		name:  "isogram with only lower case characters",
		input: "isogram",
		want:  true,
	},
	{
		name:  "word with one duplicated character",
		input: "eleven",
		want:  false,
	},
	{
		name:  "longest reported english isogram",
		input: "subdermatoglyphic",
		want:  true,
	},
	{
		name:  "word with duplicated character in mixed case",
		input: "Alphabet",
		want:  false,
	},
	{
		name:  "hypothetical isogrammic word with hyphen",
		input: "thumbscrew-japingly",
		want:  true,
	},
	{
		name:  "isogram with duplicated hyphen",
		input: "six-year-old",
		want:  true,
	},
	{
		name:  "made-up name that is an isogram",
		input: "Emily Jung Schwartzkopf",
		want:  true,
	},
	{
		name:  "duplicated character in the middle",
		input: "accentor",
		want:  false,
	},
}

func TestIsIsogram(t *testing.T) {
	for _, c := range testCases {
		if IsIsogram(c.input) != c.want {
			t.Fatalf("FAIL: %s\nWord %q, want %t, got %t", c.name, c.input, c.want, !c.want)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range testCases {
			IsIsogram(c.input)
		}

	}
}
