package clock

import (
	"reflect"
	"testing"
)

// Clock type API:
//
// Time(hour, minute int) Clock    // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock

const testVersion = 2

func TestCreateClock(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, n := range timeTests {
		if got := Time(n.h, n.m); got.String() != n.want {
			t.Fatalf("Time(%d, %d) = %q, want %q", n.h, n.m, got, n.want)
		}
	}
	t.Log(len(timeTests), "test cases")
}

func TestAddMinutes(t *testing.T) {
	for _, a := range addTests {
		if got := Time(a.h, a.m).Add(a.a); got.String() != a.want {
			t.Fatalf("Time(%d, %d).Add(%d) = %q, want %q",
				a.h, a.m, a.a, got, a.want)
		}
	}
	t.Log(len(addTests), "test cases")
}

func TestCompareClocks(t *testing.T) {
	for _, e := range eqTests {
		clock1 := Time(e.c1.h, e.c1.m)
		clock2 := Time(e.c2.h, e.c2.m)
		got := clock1 == clock2
		if got != e.want {
			t.Log("Clock1:", clock1)
			t.Log("Clock2:", clock2)
			t.Logf("Clock1 == Clock2 is %t, want %t", got, e.want)
			if reflect.DeepEqual(clock1, clock2) {
				t.Log("(Hint: see comments in clock_test.go.)")
			}
			t.FailNow()
		}
	}
	t.Log(len(eqTests), "test cases")
}
