package acronym

import (
	"testing"
)

type acronymTest struct {
	input    string
	want string
}

var stringTestCases = []acronymTest{
	{
		input:    "Portable Network Graphics",
		want: "PNG",
	},
	{
		input:    "Ruby on Rails",
		want: "ROR",
	},
	{
		input:    "First In, First Out",
		want: "FIFO",
	},
	{
		input:    "GNU Image Manipulation Program",
		want: "GIMP",
	},
	{
		input:    "Complementary metal-oxide semiconductor",
		want: "CMOS",
	},
}

func TestAcronym(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Abbreviate(test.input)
		if actual != test.want {
			t.Errorf("Acronym test [%s], expected [%s], actual [%s]", test.input, test.want, actual)
		}
	}
}

func BenchmarkAcronym(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Abbreviate(test.input)
		}
	}
}
