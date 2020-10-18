package assert

import (
	"errors"
	"testing"
)

func fiveNoError() (int, error) {
	return 5, nil
}

func TestAllMethods(t *testing.T) {
	a := New(t)

	a.Log("Can call testing.T methods.")
	a.NoError(nil)
	a.Error(errors.New("This is an error!"))
	a.ErrorContains(errors.New("This is an error!"), "error")
	a.True(true)
	a.False(false)
	a.Nil(nil)
	a.NotNil(a)
	a.Equal(5, 5)
	a.NotEqual(2, 4)
	a.Equal(5, a.R1AndNoError(fiveNoError()))

	// can't think of any way to test fail cases without failing awkward!
}
