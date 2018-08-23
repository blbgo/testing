// Package assert is a testing tool for go.
package assert

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// Assert is the top level type and holds the testing instance
type Assert struct {
	*testing.T
}

// New Creates an assert object that will use the given testing instance
func New(t *testing.T) *Assert {
	return &Assert{
		T: t,
	}
}

// Log writes to the testing log.  Will be printed only if the test fails
// or the -test.v flag is set
func (a *Assert) Log(args ...interface{}) {
	a.Helper()
	a.T.Log(args...)
}

// Logf writes to the testing log.  Will be printed only if the test fails or
// the -test.v flag is set
func (a *Assert) Logf(format string, args ...interface{}) {
	a.Helper()
	a.T.Logf(format, args...)
}

func (a *Assert) log(args ...interface{}) {
	if len(args) > 0 {
		a.T.Log(args...)
	}
}

// Fail forces the test to fail writing a log message if any arguments
// provided
func (a *Assert) Fail(args ...interface{}) {
	a.Helper()
	a.log(args...)
	a.Fatal(errors.New("Forced fail"))
}

// NoError asserts that the provided err is nil
func (a *Assert) NoError(err error, args ...interface{}) {
	if err != nil {
		a.Helper()
		a.log(args...)
		a.Fatal(err)
	}
}

// IsError asserts that the provided err is not nil
func (a *Assert) IsError(err error, args ...interface{}) {
	if err == nil {
		a.Helper()
		a.log(args...)
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
