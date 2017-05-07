# Overview

[![Build Status](https://travis-ci.org/tinyhubs/assert.svg?branch=master)](https://travis-ci.org/tinyhubs/assert)
[![GoDoc](https://godoc.org/github.com/tinyhubs/assert?status.svg)](https://godoc.org/github.com/tinyhubs/assert)
[![Language](https://img.shields.io/badge/language-go-lightgrey.svg)](https://github.com/tinyhubs/assert)
[![License](https://img.shields.io/badge/license-New%20BSD-yellow.svg?style=flat)](LICENSE)
[![codecov](https://codecov.io/gh/tinyhubs/assert/branch/master/graph/badge.svg)](https://codecov.io/gh/tinyhubs/assert)
[![goreport](https://www.goreportcard.com/badge/github.com/tinyhubs/assert)](https://www.goreportcard.com/report/github.com/tinyhubs/assert)

Package assert is a rich assertor and extensiable assert tools.

# Sample usage

- Step 1: Import the assert package.

```go
import "github.com/tinyhubs/assert"
```

- Step 2: Create a new Assert object at the begin of the test function.

```go
expect := assert.New(t)
```

or

```go
expect := &assert.Assert{t}
```

- Step 3: Use the `Equal`,`NotEqual`,`True`,`False`,`Panic`,`NoPanic` function to check your result.

```go
expect.Equal("We expect equal", "123", "456")
```

# Function

- `Equal()` `NotEqual()` : check the two value is equal or not

- `True()` `False()` : check the value is true of false

- `Panic()` `NoPanic()` : check the function is panic or not panic

# Examples

See [assert_test.go](assert_test.go)

# Notice

Maybe you found that the parameter `message` of the assert functions looks odd.
That I planned for it. That because I'm from java. And I found that many people like to
ignore the `message` parameter in the unit test code. But that made the code is hard to maintained.
So, after I'm in go, I decide to force the people give a message for assert function.


```
	//assert(t, "Expect-the-values-is-equal").Equal("123", "456")
	//
	//assert.Equal(t, "Expect-the-values-is-equal", "123", "456")
	//
	//assert.Equal("123", "456", t, "Expect-the-values-is-equal")
	//
	//assert.Equal(t, "123", "456")
	//assert.Equalf(t, "123", "456", "Expect-the-values-is-equal")
	//
	//assert.Equal(t, "123", "456")
	//assert.Equalr(t, "Expect-the-values-is-equal", "123", "456")
	//assert.Equali(t, "Expect-the-values-is-equal", "123", "456")
	//assert.Emptyi(t, "Expect-the-values-is-equal", "123", "456")
	//assert.NoPanicr(t, "Expect-the-values-is-equal", "123", "456")

	//assert.Equal(t, "Expect-the-values-is-equal", "123", "456")
	//assert.Equalh(t, "Expect-the-values-is-equal", "123", "456")
	//assert.Equalp(t, "Expect-the-values-is-equal", "123", "456")
	//assert.Equalq(t, "Expect-the-values-is-equal", "123", "456")
	//assert.Equalq(t, "Expect-the-values-is-equal", "123", "456")
```