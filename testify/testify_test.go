package testify_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		want := 555
		got := 555

		assert.Equal(t, want, got, "they should be equal")
	})

	t.Run("not equal", func(t *testing.T) {
		want := 888
		got := 999

		assert.NotEqual(t, want, got, "they should be equal")
	})

	t.Run("nil", func(t *testing.T) {
		var p *Person
		assert.Nil(t, p)
	})

	t.Run("not nil", func(t *testing.T) {
		pp := &Person{FirstName: "AnuchitO"}
		if assert.NotNil(t, pp) {
			// now we know that object isn't nil, we are safe to make
			// further assertions without causing any errors
			assert.Equal(t, "AnuchitO", pp.FirstName)

		}
	})

}
