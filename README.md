# Overview

[![Build Status](https://travis-ci.org/tinyhubs/assert.svg?branch=master)](https://travis-ci.org/tinyhubs/assert)
[![GoDoc](https://godoc.org/github.com/tinyhubs/assert?status.svg)](https://godoc.org/github.com/tinyhubs/assert)
[![Language](https://img.shields.io/badge/language-go-lightgrey.svg)](https://github.com/tinyhubs/assert)
[![License](https://img.shields.io/badge/license-New%20BSD-yellow.svg?style=flat)](LICENSE)
[![codecov](https://codecov.io/gh/tinyhubs/assert/branch/master/graph/badge.svg)](https://codecov.io/gh/tinyhubs/assert)
[![goreport](https://www.goreportcard.com/badge/github.com/tinyhubs/assert)](https://www.goreportcard.com/report/github.com/tinyhubs/assert)

`et` is a rich assertor and extensiable assert tools.

# Installation

et do not required the others package.

```bash
$ go get -u github.com/tinyhubs/et
```

# Examples

```go
package examples

import (
	"testing"
	"github.com/tinyhubs/et/assert"
)

func TestAssert_Equal(t *testing.T) {
	assert.Equal(t, "123", "456")
}

func TestAssert_Equali(t *testing.T) {
	assert.Equali(t, "Expect-the-values-is-equal", "123", "456")
}
```

The output looks like below. 每个失败的用例的第二行堆栈信息就是断言出错的的assert或者是expect的代码的位置.

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example1_test.go:9
		Expect:123, Actual:456
	et-core.go:18:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example1_test.go:13
		Expect-the-values-is-equal
		Expect:123, Actual:456
```

# Function

为了节省打字,et提供了大量的断言函数.但是,et提供的只能满足大家通常的诉求,
你总会遇到这些断言函数无法满足您的诉求的场景.因此,et还提供了一个扩展机制方便您扩展出自己的断言函数.

## 基本断言

- 相等检测: `assert.Equal`

```go
assert.Equal(t, "123", "456")
```

- 布尔检测: `assert.True`

```go
assert.True(t, "123" == "456")
```

- 检测是否会抛异常: `assert.Panic`

```go
assert.Panic(t, func() { /* Do nothing. */ })
```

- 检测是否匹配正则表达式: `assert.Match`

```go
assert.Match(t, `^[a-zA-Z0-9-_]+@timo\.com$`, "libbylg@126.com")
```

- 检测是否为nil: `assert.Nil`

```go
assert.Nil(t, bytes.NewBufferString(""))
```

## assert vs expect

et提供了assert和expect两套接口,所有在assert中存在的函数,except中也会存在.
比如,assert里面有`assert.Equal`,那么也会存在`except.Equal`.
assert系接口和expect系的接口也可以组合使用,以便实现更丰富的功能.

assert和expect的区别就是assert会立即中断当前用例的执行,而expect会一直执行下去直到用例执行完毕.
下面有个例子,注意检查行号:

```go
package examples

import (
	"testing"
	"github.com/tinyhubs/et/expect"
	"github.com/tinyhubs/et/assert"
)

func Test_2_Expect_Equal(t *testing.T) {
	expect.Equal(t, "123", "456")
	expect.Equal(t, 333, 7788)
	expect.Equal(t, "sina", "sina")
} //  stoped here

func Test_2_Assert_Equal(t *testing.T) {
	assert.Equal(t, "123", "456") //  stoped here
	assert.Equal(t, 333, 7788)
	assert.Equal(t, "sina", "sina")
}
```

这个例子的输出如下:

```
	et-core.go:28:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:10
		Expect:123, Actual:456
	et-core.go:28:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:11
		Expect:333, Actual:7788
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:16
		Expect:123, Actual:456
```

考虑到代码的可读性,将断言的意图作为assert调用的一个参数是个很好的实践,et也支持这种用法.
et的每个基本assert接口都支持一个`i`后缀的函数,使用这些`i`后缀的函数你可以在assert调用的时候,将你得断言意图作为调用的一个参数.

比如,通常我们会这样写测试代码:

```go
//  Check the email is xxx@timo.com
assert.Match(t, `^[a-zA-Z0-9-_]+@timo\.com$`, email)
```

但是更好的实践是这样写:

```go
assert.Matchi(t, `Check the email is xxx@timo.com`, `^[a-zA-Z0-9-_]+@timo\.com$`, email)
```


## 反逻辑的断言

为提高断言代码的可用性,et为每个正向断言和反向断言都提供了函数. 比如:
如果你用`assert.Equal`来检测两个数据是否相等,那么你也应该知道其实您也可以使用`assert.NotEqual`来检测两个数不相等.
类似的还有:

`assert.Equal` vs `assert.NotEqual`

`assert.True` vs `assert.False`

`assert.Panic` vs `assert.NoPanic`

`i`系列函数也是支持的:

`assert.Equali` vs `assert.NotEquali`

`assert.Truei` vs `assert.Falsei`

`assert.Panici` vs `assert.NoPanici`


## Need more assert or expect functions

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