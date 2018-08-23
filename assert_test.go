package assert

import (
	"errors"
	"testing"
)

func TestAllMethods(t *testing.T) {
	a := New(t)

	a.Log("Somthing to log, only printed if test fails")
	a.Logf("Formated log, only %v if test fails", "printed")
	a.NoError(nil)
	a.IsError(errors.New("This is an error!"))
	a.ErrorContains(errors.New("This is an error!"), "error")
	a.IsTrue(true)
	a.IsFalse(false)
	a.IsNil(nil)
	a.NotNil(a)
	a.AreEqual(5, 5)
	a.AreNotEqual(2, 4)

	// can't think of any way to test fail cases without failing awkward!
}
