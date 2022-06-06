package reverse

import (
	"testing"
	"testing/quick"
)

type reverseTestCase struct {
	description string
	input       string
	expected    string
}

var testCases = []reverseTestCase{
	{
		description: "an empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "a word",
		input:       "robot",
		expected:    "tobor",
	},
	{
		description: "a capitalized word",
		input:       "Ramen",
		expected:    "nemaR",
	},
	{
		description: "a sentence with punctuation",
		input:       "I'm hungry!",
		expected:    "!yrgnuh m'I",
	},
	{
		description: "a palindrome",
		input:       "racecar",
		expected:    "racecar",
	},
	{
		description: "a multi-byte test case",
		input:       "Hello, 世界",
		expected:    "界世 ,olleH",
	},
}

func TestReverse(t *testing.T) {
	for _, testCase := range testCases {
		if res := String(testCase.input); res != testCase.expected {
			t.Fatalf("FAIL: %s(%s)\nExpected: %q\nActual: %q",
				testCase.description, testCase.input, testCase.expected, res)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestReverseOfReverse(t *testing.T) {
	assertion := func(s string) bool {
		return s == String(String(s))
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			String(test.input)
		}
	}
}
