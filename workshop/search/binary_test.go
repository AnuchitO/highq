package binary

import (
	"math/rand"
	"sort"
	"testing"
)

var testData = []struct {
	name  string
	slice []int
	key   int
	found int // expected result
}{
	{"6 found at index 3",
		[]int{1, 3, 4, 6, 8, 9, 11},
		6, 3},
	{"9 found at index 5",
		[]int{1, 3, 4, 6, 8, 9, 11},
		9, 5},
	{"3 found at index 1",
		[]int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144},
		3, 1},
	{"55 found at index 7",
		[]int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144},
		55, 7},
	{"21 found at index 5",
		[]int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377},
		21, 5},
	{"34 found at index 6",
		[]int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377},
		34, 6},

	{"1 found at beginning of slice",
		[]int{1, 3, 6},
		1, 0},
	{"6 found at end of slice",
		[]int{1, 3, 6},
		6, 2},
	{"2 > 1 at index 0, < 3 at index 1",
		[]int{1, 3, 6},
		2, 1},
	{"2 < all values",
		[]int{11, 13, 16},
		2, 0},
	{"21 > all 3 values",
		[]int{11, 13, 16},
		21, 3},
	{"1 found at beginning of slice",
		[]int{1, 1, 1, 1, 1, 3, 6}, // duplicates
		1, 0},
	{"3 found at index 1",
		[]int{1, 3, 3, 3, 3, 3, 6},
		3, 1},
	{"6 found at index 4",
		[]int{1, 3, 3, 3, 6, 6, 6},
		6, 4},
	{"-2 > -3 at index 1, < -1 at index 2",
		[]int{-6, -3, -1}, // negatives
		-2, 2},
	{"0 > -7 at index 4, < 1 at index 5",
		[]int{-19, -17, -15, -12, -7, 1, 14, 35, 69, 124},
		0, 5},
	{"slice has no values",
		nil, 0, 0},
}

func TestSearchInts(t *testing.T) {
	for _, test := range testData {
		if !sort.IntsAreSorted(test.slice) {
			t.Skip("Invalid test data")
		}
		if x := SearchInts(test.slice, test.key); x != test.found {
			t.Fatalf("SearchInts(%v, %d) = %d, want %d",
				test.slice, test.key, x, test.found)
		}
	}
}

// Did you get it?  Did you cut and paste from the standard library?
// Whatever.  Now show you can work it.
//
// The test program will supply slices and keys and you will write a function
// that searches and returns one of the following messages:
//
//   k found at beginning of slice.
//   k found at end of slice.
//   k found at index fx.
//   k < all values.
//   k > all n values.
//   k > lv at lx, < gv at gx.
//   slice has no values.
//
// In your messages, substitute appropritate values for k, lv, and gv;
// substitute indexes for fx, lx, and gx; substitute a number for n.
//
// In the function Message you will demonstrate a number of different ways
// to test the result of SearchInts.  Note that you would probably never need
// all of these different tests in a real program.  The point of the exercise
// is just to show that it is possible to identify a number of different
// conditions from the return value.
func TestMessage(t *testing.T) {
	for _, test := range testData {
		if !sort.IntsAreSorted(test.slice) {
			t.Skip("Invalid test data")
		}
		if res := Message(test.slice, test.key); res != test.name {
			t.Fatalf("Message(%v, %d) =\n%q\nwant:\n%q",
				test.slice, test.key, res, test.name)
		}
	}
}

// Benchmarks also test searching larger random slices

func Benchmark1e2(b *testing.B) {
	s := make([]int, 1e2)
	for i := range s {
		s[i] = rand.Intn(len(s))
	}
	sort.Ints(s)
	k := rand.Intn(len(s))
	name := sort.SearchInts(s, k)
	res := SearchInts(s, k)
	if name != res {
		b.Fatalf(
			"Search of %d values gave different answer than sort.SearchInts",
			len(s))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInts(s, k)
	}
}

func Benchmark1e4(b *testing.B) {
	s := make([]int, 1e4)
	for i := range s {
		s[i] = rand.Intn(len(s))
	}
	sort.Ints(s)
	k := rand.Intn(len(s))
	name := sort.SearchInts(s, k)
	res := SearchInts(s, k)
	if name != res {
		b.Fatalf(
			"Search of %d values gave different answer than sort.SearchInts",
			len(s))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInts(s, k)
	}
}

func Benchmark1e6(b *testing.B) {
	s := make([]int, 1e6)
	for i := range s {
		s[i] = rand.Intn(len(s))
	}
	sort.Ints(s)
	k := rand.Intn(len(s))
	name := sort.SearchInts(s, k)
	res := SearchInts(s, k)
	if name != res {
		b.Fatalf(
			"Search of %d values gave different answer than sort.SearchInts",
			len(s))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInts(s, k)
	}
}
