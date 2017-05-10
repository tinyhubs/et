# 概要

[![Build Status](https://travis-ci.org/tinyhubs/assert.svg?branch=master)](https://travis-ci.org/tinyhubs/assert)
[![GoDoc](https://godoc.org/github.com/tinyhubs/assert?status.svg)](https://godoc.org/github.com/tinyhubs/assert)
[![Language](https://img.shields.io/badge/language-go-lightgrey.svg)](https://github.com/tinyhubs/assert)
[![License](https://img.shields.io/badge/license-New%20BSD-yellow.svg?style=flat)](LICENSE)
[![codecov](https://codecov.io/gh/tinyhubs/assert/branch/master/graph/badge.svg)](https://codecov.io/gh/tinyhubs/assert)
[![goreport](https://www.goreportcard.com/badge/github.com/tinyhubs/assert)](https://www.goreportcard.com/report/github.com/tinyhubs/assert)

`et` 全名是`EasyTest`,它是一个有着丰富断言函数的且可扩展的测试辅助库.

# 安装

`et`无额外的包依赖,您可以直接下载代码到本地,也可以通过下面的命令获取:

```bash
$ go get -u github.com/tinyhubs/et
```

# 示例

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

上面的例子的输出如下,每个失败的用例的第二行堆栈信息就是断言出错的的assert或者是expect的代码的位置.

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example1_test.go:9
		Expect:123, Actual:456
	et-core.go:18:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example1_test.go:13
		Expect-the-values-is-equal
		Expect:123, Actual:456
```

# 功能
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

assert和expect的区别就是assert会立即中断当前用例的执行,而expect会一直执行下去直到用例执行完毕:

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

## 断言意图作为代码

测试代码也需要良好的维护性,但是大家写测试代码的时候往往只能在断言代码附近写上注释来表达断言的意图.
但是,注释是很容易失去维护,所以如果能将断言的意图作为assert调用的一个参数那么大家修改测试代码时就不会再忽略掉了.
这是一个很好的实践,et也支持这种用法.

et的每个基本assert接口都支持一个`i`后缀的函数,使用这些`i`后缀的函数你可以在assert调用的时候,将你得断言意图作为调用的一个参数.

比如,通常我们会这样写测试代码,注释放在assert函数调用的注释里面:

```go
//  Check the email is xxx@timo.com
assert.Match(t, `^[a-zA-Z0-9-_]+@timo\.com$`, email)
```

但是更好的实践是直接将上面的注释写到断言函数参数里面:

```go
assert.Matchi(t, "Check the email is xxx@timo.com", `^[a-zA-Z0-9-_]+@timo\.com$`, email)
```


## 反逻辑断言(NOT-Assertor)

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


## 扩展断言机制

假设一个场景,我们现在需要一个判断某个整数是否在某个范围之内的断言怎么办?这个是et默认没有提供的断言.
但是,聪明如你,你一定会想到其实可以用`assert.True`来实现这个断言.所以你可能会这样写代码:

```go
assert.True(t, (value >= min) && (value <= max))
```

这个代码可以工作,但是我们总有那么点遗憾:为啥et不提供一个`assert.Inrange`的断言函数呢?这样咱们就可以这样写代码了:

```go
assert.Inrange(t, min, max, value)
```

其实,et提供了一种机制可以帮你实现自己的断言.

#### 方式1,创建一个判断范围的`Assertor`

采用这种方式时,需要你提供一个实现了`et.Assertor`接口的类,您可以这样做:

```go
type Inrange struct {
	Min   int
	Max   int
	Value int
}

func (i *Inrange) Assert() error {
	if (i.Value >= i.Min) && (i.Value <= i.Max) {
		return nil
	}

	return fmt.Errorf("Expect in range [%v, %v], Actual: %v", i.Min, i.Max, i.Value)
}

```

然后,这样使用(注意需要先`import "github.com/tinyhubs/et/et"`):

```go
et.Assert(&Inrange{min, max, value}, t)
```

如果触发了断言失败,可以获得下面的结果,注意看第二行提示,这个就是我们Assertor返回的error哦:

```
et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example3_test.go:27
		Expect in range [1, 100], Actual: 320
```

`et.Assert`的用法好像跟`assert.Equal`其他的不一样对吧,但其实assert.Equal也是使用上面的Assertor类似的机制实现的.
这就涉及到下面的扩展方式2.

#### 方式2,自行封装assert函数

你可以在方式1的基础上再额外提供下面一个断言函数:

```go
func AssertInrange(t *testing.T, min int, max int, value int) {
	et.AssertInner(t, "", &Inrange{min, max, value}, 2)
}
```

然后,方式1中的实例代码可以改成下面这样:

```go
assertInrange(t, min, max, value)
```

输出结果如下,和方式1的输出结果相同:

```
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example3_test.go:40
		Expect in range [1, 100], Actual: 320
```

您可以在[et/examples/example3_test.go](et/examples/example3_test.go)查看这个完整的扩展的示例,您可以试着运行下diamante看看效果.

如果您觉得自己开发了一个很不错的扩展,请必要忘记分享给你周边的同事,如果觉得您的扩展可以帮到更多人,那么直接发一个pull request给我吧,或许我可以合并到et中去.

# et的设计和演化

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