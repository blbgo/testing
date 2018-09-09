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
	a.IsError(errors.New("This is an error!"))
	a.ErrorContains(errors.New("This is an error!"), "error")
	a.IsTrue(true)
	a.IsFalse(false)
	a.IsNil(nil)
	a.NotNil(a)
	a.AreEqual(5, 5)
	a.AreNotEqual(2, 4)
	a.AreEqual(5, a.R1AndNoError(fiveNoError()))

	// can't think of any way to test fail cases without failing awkward!
}
