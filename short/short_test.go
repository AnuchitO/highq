package short

import "testing"

func TestShort(t *testing.T) {
	t.Log(testing.Short())
}

func TestShortExampleForTestThatIsLong(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// too long to run
}
