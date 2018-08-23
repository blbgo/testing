## assert
A testing assert tool for go.

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg)](http://godoc.org/github.com/blbergwall/assert)
[![Go Report Card](https://goreportcard.com/badge/github.com/blbergwall/assert)](https://goreportcard.com/report/github.com/blbergwall/assert)
[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://github.com/blbergwall/assert/blob/master/LICENSE.txt)

### Install
```
go get github.com/blbergwall/assert
```

### Overview
Provide a collection of assert and testing helper methods

When an assert fails t.Fatal is called ending the current test with failure

### Usage

```
func TestATest(t *testing.T) {
	a := assert.New(t)

	result, err := methodUnderTest()
	a.NoError(err)
	a.AreEqual(5, result)
}
```

### Why?
Mainly because this package is used by another package I plan to provide.

I am aware of [stretchr/testify](https://github.com/stretchr/testify) and
maybe I should switch to using that.

Also I am fairly new to go and publishing open source code and want to learn
more about both.

### License
[MIT](https://github.com/blbergwall/depend/blob/master/LICENSE.txt)
