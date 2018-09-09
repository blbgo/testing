// Package assert is a testing tool for go.
package assert

import (
	"fmt"
	"strings"
	"testing"
)

// Assert is a testing.TB with assert functionality
type Assert struct {
	testing.TB
}

// New Creates an assert object that will use the given testing instance
func New(tb testing.TB) *Assert {
	return &Assert{TB: tb}
}

func (a *Assert) logIfSomthing(args ...interface{}) {
	if len(args) > 0 {
		a.Log(args...)
	}
}

// NoError asserts that the provided err is nil
func (a *Assert) NoError(err error, args ...interface{}) {
	if err != nil {
		a.Helper()
		a.logIfSomthing(args...)
		a.Fatal(err)
	}
}

// IsError asserts that the provided err is not nil
func (a *Assert) IsError(err error, args ...interface{}) {
	if err == nil {
		a.Helper()
		a.logIfSomthing(args...)
		a.Fatal("Expected error but NO error!")
	}
}

// ErrorContains asserts that the provided err is not nil and that it contains
// the provided string
func (a *Assert) ErrorContains(err error, substr string) {
	if err == nil {
		a.Helper()
		a.Fatalf("Expected error containing \"%s\" but no error!", substr)
	} else if !strings.Contains(err.Error(), substr) {
		a.Helper()
		a.Fatalf("Error \"%s\" should contain \"%s\" but does not!", err, substr)
	}
}

// IsTrue asserts that the provided bool is true
func (a *Assert) IsTrue(b bool, args ...interface{}) {
	if !b {
		a.Helper()
		if len(args) > 0 {
			a.Fatal(args...)
		}
		a.Fatal("Expected true but is false")
	}
}

// IsFalse asserts that the provided bool is false
func (a *Assert) IsFalse(b bool, args ...interface{}) {
	if b {
		a.Helper()
		if len(args) > 0 {
			a.Fatal(args...)
		}
		a.Fatal("Expected false but is true")
	}
}

// IsNil asserts that the provided interface{} is nil
func (a *Assert) IsNil(i interface{}, args ...interface{}) {
	if i != nil {
		a.Helper()
		if len(args) > 0 {
			a.Fatal(args...)
		}
		a.Fatal("Expected nil but is not nil")
	}
}

// NotNil asserts that the provided interface{} is not nil
func (a *Assert) NotNil(i interface{}, args ...interface{}) {
	if i == nil {
		a.Helper()
		if len(args) > 0 {
			a.Fatal(args...)
		}
		a.Fatal("Expected not nil but is nil")
	}
}

// AreEqual asserts that the expect == actual
func (a *Assert) AreEqual(expect interface{}, actual interface{}, args ...interface{}) {
	if expect != actual {
		a.Helper()
		if len(args) > 0 {
			a.Fatal(args...)
		}
		a.Fatal(fmt.Sprintf("Expected %v but is %v", expect, actual))
	}
}

// AreNotEqual asserts that the expect != actual
func (a *Assert) AreNotEqual(expect interface{}, actual interface{}, args ...interface{}) {
	if expect == actual {
		a.Helper()
		if len(args) > 0 {
			a.Fatal(args...)
		}
		a.Fatal(fmt.Sprintf("Expected not %v but is %v", expect, actual))
	}
}

// R1AndNoError is ment to be called with the results of a function that
// returns 1 result and a posible error.  it fails if the error is not nil and
// otherwise returns just the result.  So it can be used in a construct like:
//   a.AreEqual(5, a.R1AndNoError(funcReturningIntAndError()))
// this asserts first that funcReturningIntAndError did not return an error
// and second that its first result is 5.
func (a *Assert) R1AndNoError(r1 interface{}, err error) interface{} {
	if err != nil {
		a.Helper()
		a.Log("Expected no error but got error")
		a.Fatal(err)
	}
	return r1
}
