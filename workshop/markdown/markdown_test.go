package markdown

import "testing"

func TestMarkdown(t *testing.T) {
	for _, test := range testCases {
		if html := Render(test.input); html != test.want {
			t.Fatalf("FAIL: Render(%q) = %q, want %q.", test.input, html, test.want)
		}
		t.Logf("PASS: %s\n", test.name)
	}
}

func BenchmarkMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Render(test.input)
		}
	}
}

var testCases = []struct {
	name  string
	input string
	want  string
}{
	{
		name:  "parses normal text as a paragraph",
		input: "This will be a paragraph",
		want:  "<p>This will be a paragraph</p>",
	},
	{
		name:  "parsing italics",
		input: "_This will be italic_",
		want:  "<p><em>This will be italic</em></p>",
	},
	{
		name:  "parsing bold text",
		input: "__This will be bold__",
		want:  "<p><strong>This will be bold</strong></p>",
	},
	{
		name:  "mixed normal, italics and bold text",
		input: "This will _be_ __mixed__",
		want:  "<p>This will <em>be</em> <strong>mixed</strong></p>",
	},
	{
		name:  "with h1 header level",
		input: "# This will be an h1",
		want:  "<h1>This will be an h1</h1>",
	},
	{
		name:  "with h2 header level",
		input: "## This will be an h2",
		want:  "<h2>This will be an h2</h2>",
	},
	{
		name:  "with h6 header level",
		input: "###### This will be an h6",
		want:  "<h6>This will be an h6</h6>",
	},
	{
		name:  "unordered lists",
		input: "* Item 1\n* Item 2",
		want:  "<ul><li>Item 1</li><li>Item 2</li></ul>",
	},
	{
		name:  "With a little bit of everything",
		input: "# Header!\n* __Bold Item__\n* _Italic Item_",
		want:  "<h1>Header!</h1><ul><li><strong>Bold Item</strong></li><li><em>Italic Item</em></li></ul>",
	},
}
