# 概要

[![Build Status](https://travis-ci.org/tinyhubs/assert.svg?branch=master)](https://travis-ci.org/tinyhubs/assert)
[![GoDoc](https://godoc.org/github.com/tinyhubs/assert?status.svg)](https://godoc.org/github.com/tinyhubs/assert)
[![Language](https://img.shields.io/badge/language-go-lightgrey.svg)](https://github.com/tinyhubs/assert)
[![License](https://img.shields.io/badge/license-New%20BSD-yellow.svg?style=flat)](LICENSE)
[![codecov](https://codecov.io/gh/tinyhubs/assert/branch/master/graph/badge.svg)](https://codecov.io/gh/tinyhubs/assert)
[![goreport](https://www.goreportcard.com/badge/github.com/tinyhubs/assert)](https://www.goreportcard.com/report/github.com/tinyhubs/assert)

`et` 全名是`EasyTest`,它是一个有着丰富断言函数的,可扩展的单元测试辅助库.

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

`et`提供了assert和expect两套接口,所有在assert中存在的函数,except中也会存在.
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

这个两个用例的输出结果如下:

```text
	et-core.go:28:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:10
		Expect:123, Actual:456
	et-core.go:28:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:11
		Expect:333, Actual:7788
	et-core.go:28:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:12
		Expect:sina, Actual:google
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example2_test.go:16
		Expect:mtn, Actual:rnd
```

## 断言意图作为代码

测试代码也需要良好的可维护性,所以大家写测试代码的时候往往会在断言代码附近写上注释来表达断言的意图.如下所示:

```go
//  Check the email is xxx@timo.com
assert.Match(t, `^[a-zA-Z0-9-_]+@timo\.com$`, email)
```

但注释毕竟不是代码,难以得到开发人员的爱,很容易失去维护,所以如果能将断言的意图作为assert函数的一个参数,那么大家修改测试代码时就不会再忽略掉了.
这是一个很好的实践,`et`也支持这种用法.`et`的每个基本assert接口都支持一个`i`后缀的函数,使用这些`i`后缀的函数你可以在调用assert函数的时候,附上断言的意图.
如下所示,我们可以利用`et`的`i`系列函数将前面这个代码改成下面这样:

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

```text
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
AssertInrange(t, min, max, value)
```

输出结果如下,和方式1的输出结果相同:

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example3_test.go:40
		Expect in range [1, 100], Actual: 320
```

您可以在[`et/examples/example3_test.go`](https://github.com/tinyhubs/et/tree/master/examples/example3_test.go)查看这个完整的扩展的示例,您可以试着运行下diamante看看效果.

如果您觉得自己开发了一个很不错的扩展,请必要忘记分享给你周边的同事,如果觉得您的扩展可以帮到更多人,那么直接发一个pull request给我吧,或许我可以合并到et中去.

# `et`的设计

## 初衷:看着舒服少打字

`et`最初的设计初衷有两个,一个是我不想写那么多if,我希望一个断言逻辑一行搞定.另外一个原因是我不想写t参数.
综合来说就是想看起来舒服且码字要少.

举个例子,我当时就想这样写测试代码:

```go
assert.Equal(a, b)
```

很明显我是有java或者c语言背景的人.这里的assert其实是个对象,assert对象内部包含了一个t参数.我很快就做了个原型库.
基于这个库的一个完整的例子大概是这样:

```go
func TestAssert_Equal(t *testing.T) {
	assert := &Assert{T: t}
	assert.Equal("Expect-the-values-is-equal", "123", "456")
}

func TestAssert_NotEqual(t *testing.T) {
	assert := New(t)
	assert.NotEqual("Expect-the-values-is-not-equal", "123", "123")
}
```

## 支持except

接着,我在使用我自己做的[properties](https://github.com/tinyhubs/properties)库和[tinydom](https://github.com/tinyhubs/tinydom)库
写一个开发辅助工具的时候,我发现`except`对我也非常有用,这样我的用例如果有什么bug测试用例能够尽可能多地帮我检测出来.
但是,我总不能再new一个Except对象出来吧?所以我就想到是不是可以做成一次性New两个对象出来:

```go
func TestAssert_NotEqual(t *testing.T) {
	assert, except := New(t)
	except.NotEqual("Expect-the-values-is-not-equal", 333, 444)
	assert.NotEqual("Expect-the-values-is-not-equal", "123", "123")
}
```

这个办法初看起来也挺美好的,因为使用者可以根据需要自己决定New出的对象用哪一个,也就是你可以有下面三种情况自由组合:

```go
_, except := New(t)       //  只使用except
assert, _ := New(t)       //  只使用assert
assert, except := New(t)  //  同时使用except和assert
```

## 过犹不及

前面,为支持`expect`所带来的设计变化,看起来没大毛病.但是我总觉得有点疙里疙瘩的.因为个方案要求在每个`Test_xxx`函数入口的地方都要`New`一个`Assert`对象甚至额外的`Except`对象.
我在一个项目里面自己试用了一下,然后又感觉每次都要`New`这么个对象显得好蠢.

其实,我之所要`New`对象就是为了少写个`t`,但是光`_, assert := New(t)`这一句敲的字都相当于敲N个`t`了----有点过犹不及了.

由于google的设计里面,t就一个字符,所以在一个断言调用里面,t字符显得并不那么突出,所以我就放弃了少写t这个愿望----虽然t也不爽.
于是,我就重新设计了`et`库,最终代码可以直接这样写:

```go
assert.Equal(t, a, b)
```

这里`assert`是包名,为了支持`except`,我单独实现了一个`except`包,所以我们也可以这样写:

```go
except.Equal(t, a, b)
```

您或许会发现,except如果是另外一个包,那就必须写两个常常的`import`语句:

```go
import (
    "github.com/tinyhubs/et/assert"
    "github.com/tinyhubs/et/except"
)
```

这个我觉得没关系,因为现在的IDE都很智能,IDE自动会帮我写上这段import代码的.

## `i`系函数与测试代码的可维护性

这样过了一段时间,我自己遇到了一个bug,于是我试图搞明白为何我的UT用例没有覆盖到这个场景,然后我突然发现有一个用例我看不懂了.
原因是代码中的注释和这段注释下面的断言代码看起来有点矛盾,我不知道那个才是我最初的目的----到底是注释是对的呢,还是断言代码是对的呢?

总之,我发现我的单元测试代码中的注释慢慢出现年久失修的征兆了.我修改测试代码的时候往往不会关注代码上下文的注释.
注释在我的IDE的色彩配置里面甚至是灰色的,IDE都在暗示我忽略注释呢.

所以,我想是否有办法让我修改代码的时候能够对一些非常重要的断言做一些说明,但是又不以注释的形式体现?

于是,我想到了将注释信息填写到`assert`语句中的方法.最初,我修改了所有的`assert`函数,这样代码就写成这样了:

```go
assert.True(t, "获得Text对象成功", nil != text1)
assert.True(t, "获得Text对象成功", nil != text2)
assert.True(t, "获得Text对象成功", nil != text3)
assert.True(t, "全空白的Text不会被读取", nil == text4)
```

为此我同事修改了多有的测试代码中的assert或者except调用.

然而,我又发现一个新问题,assert函数里面,如果带上一个参数看起来确实要好很多,可维护性非常好.
但是,有时候写断言意图(下称"注解")的时间也比较麻烦,甚至很多语句干事情类似存在大量重复,比如上面这段代码里面前三个注解都一样.
所以,咱们干么这么愚蠢地写相同的注解呢?

于是,我就想是否有办法当我们需要写注解的时候就写上,但是不需要的时候不写也没问题,比如下面这些代码:

```go
assert.True(t, "获得Text对象成功", nil != text1)
assert.True(t, nil != text2)
assert.True(t, nil != text3)
assert.True(t, "全空白的Text不会被读取", nil == text4)
```

可惜go语言好像支持不了啊,go语言的函数的签名是固定的,像上面的`except.True`其第二个参数不能一会儿是字符串一会儿又是bool----
真怀恋C++的函数重载和模板机制.

不过go语言里面变参是支持的,所以我们可以将这个字符串放后面:

```go
assert.True(t, nil != text1, "获得Text对象成功")
assert.True(t, nil != text2, "")
assert.True(t, nil != text3, "")
assert.True(t, nil == text4, "全空白的Text不会被读取")
```

看起来也不错!而且,利用可变参数上面的代码可以写成这样:

```go
assert.True(t, nil != text1, "获得Text对象成功")
assert.True(t, nil != text2)
assert.True(t, nil != text3)
assert.True(t, nil == text4, "全空白的Text不会被读取")
```

OK,很完美!但是,没过多久我实现了`assert.Panic`和`assert.NoPanic`两个新函数,在写一个`assert.NoPanic`的调用样例的时候,又觉得不舒服了:

```go
assert.NoPanic(t, func() {
    throwPanic()
}, "Expect-the-func-do-not-throw-a-panic")
```

这个例子好丑啊,func()后面这个字符串就像是一只苗条的猫长了一条鳄鱼的粗壮尾巴.所以,我最终还是觉得写成下面这样,代码看起来会更舒服点:

```go
assert.NoPanic(t, "Expect-the-func-do-not-throw-a-panic",
func() {
    throwPanic()
})
```

将注解放在assert函数的最后作为变参还有一些不好的地方,因为函数原型最后都成为了变长的了,开发人员看到assert函数的原型对于理解这个assert的用法会产生疑惑.

再者,我当时正在考虑设计我的assert库可扩展能力,如果将注解放在函数的后面会导致大家扩展出来的新的断言函数的原型会跟我的assert库自带的assert函数的原型不一致,
这种不一致会加大大家的学习难度.

后来,我想到了go语言的fmt包以及c语言的库,他们有这样的用法:

```go
fmt.Printf("%s", mystr)
```

```cpp
printf("%s", mystr);
```

他们都是在基本核心单词的基础上增加简短的后缀来表明函数和同系列的函数的区别的.

我觉得我的assert库可以采用同样的机制:

```go
assert.TrueXX(t, "获得Text对象成功", nil != text1)
assert.True(t, nil != text2)
```

OK,那么这个XX应该设计为啥?

`f`? No,`f`是`format`的意思,跟注解不沾边.那么哪个单词可以表达注解?

`message`, `tips`, `notice`, `information` ...

以他们的首字母放在最后看起来是这样的:

```go
assert.Equalm(t, "Expect-the-values-is-equal", "123", "456")
assert.Equalt(t, "Expect-the-values-is-equal", "123", "456")
assert.Emptyn(t, "Expect-the-values-is-equal", "123", "456")
assert.Emptyi(t, "Expect-the-values-is-equal", "123", "456")
```

`m` `t` `n`者三个字符很容易融入到原始单词里面去,比如一言看上去似乎`qualm` `qualt` `mptyn`是一个单词.
好吧他们亲和力真强. `Printf`里面的`f`之所以亲和力没那么强是因为`f`通常作为单词的开头,很少有单词以f结尾,而且`f`在英文中看起来很瘦长,总是跟其他的字母不搭调.
而`m` `t` `n` 在单词里面见得很多,特别是以`t`和`n`为后缀的单词相当多.

看起来`Emptyi`是比较不错的选择,`i`系列函数就是这么诞生的.

## 扩展能力的设计

`et`最初只提供了`Equal`,`True`,`Panic`以及反逻辑的`NotEqual`,`False`,`NoPanic`这几个函数.但是,很快我发现我还需要增加判断是否为nil,
以及正则表达式匹配的断言函数.此时,我开始意识到,虽然`assert.True`,`assert.False`可以搞定一切,但是断言的提示信息并不友好.

下面这两个用例,断言逻辑其实一模一样,但是其断言错误提示却差别很大:

```go
func Test_output1(t *testing.T) {
	assert.Equal(t, 111, 222)
}

func Test_output2(t *testing.T) {
	assert.True(t, 111 == 222)
}
```

`Test_output1`的输出如下:

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example4_test.go:9
		Expect:111, Actual:222
```

`Test_output2`的输出如下:

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example4_test.go:13
		Expect:true, Actual:false
```

很明显`Test_output1`的输出对我们的帮助更大,`Test_output2`几乎没有提供额外的价值信息.
所以,总结来说,我们还是应该提供针对某种类型的断言的专用函数,而不是使用笼统的`assert.True`或者`assert.False`来凑活.

但对于一个程序库(`et`)的作者,我没法预见到其他人需要什么样的断言函数,我甚至都不知道自己未来需要什么断言函数.
所以,`et`库支持各种不同的断言检测逻辑是很有意义的.

所以,将原库的代码拆分成了三个部分:`et`,`assert`,`expect`. `et`是核心,`et-core.go`里面定义了统一的断言函数的调用形式,
封装了获取代码堆栈的逻辑.
并提供了一个`Assertor`接口,`et`默认支持的`Assertor`定义在`et-ext.go`文件中.`assert`包和`expect`包是两种基于`et`做的封装.
基本上前面的绝大多数例子中调用的断言函数都是来自这两个包.

最先被抽象出来的是`et.AssertInner`这是转给扩展断言库的作者使用的.它主要的作用就是使用统一的格式输出断言结果.同时,也是为了扩展库的
作者不必太过关心获取调用堆栈这些细节.
为此,`et.AssertInner`函数有一个额外的`callerSkip`参数,用于帮助`et`库确定需要跳过多少层堆栈.
如果我们提供的断言函数不是被封装很多层,那么`callerSkip`就永远是`2`.
对于写断言扩展的同学来说,确定你提供给用户的是那一层的接口是非常重要的,这可以有效减少学习成本.所以,我建议永远只在`et.AssertInner`之上
包装一层.

`et.Assert`,`et.Expect`者两个是随后抽象出来的.最初,我并不想提供这两个函数,但是我在自己写一些定制的断言函数的时候,我发现有些定制的
断言函数的使用频率并不是很高,通用化的意义也不大,所以我不向去做`AssertInRange`这样的二次封装.但是我又不能直接在断言代码中调用
`et.AssertInner`.所以,提供`et.Assert`,`et.Expect`,`et.Asserti`,`et.Expecti`这四个函数可以让我在懒惰一点:).

在扩展库的设计上,`et`力求平衡`使用感知`与`代码美观度`.

如下者几行代码,看起来还是有点差异,但是他们的共同特点就是`assert`总是与`Equal`靠近.

```go
assert.Equal(t, a, b)
assert.Equali(t, "message", a, b)
et.Assert(&Equal{a, b}, t)
AssertEqual(t, a, b)
```

再比如,`i`系函数里面,`message`总是在`t`的后面,即使是`et.AssertInner`,`et.Asserti`也不例外.

下面列出了`et`所遵循的设计约定:

1. 一个断言检测支持三种调用形式:

- `et.Assert`  按扩展库的方式调用
- `assert.XXX` 通常用到的断言函数
- `expect.XXX` except断言模式

2. 如果断言函数的第一个参数总是t,除了扩展模式

3. 带`i`后缀的函数都支持一个额外的字符串参数,该参数总是在`t`参数的后面.

4. `Assert`总是应该与断言的类型靠近

5. 断言函数或者`Assertor`的预期值总是在被检测值得前面


## 其他:断言库改名

`et`库最早其实叫`assert`,但是后来在支持扩展能力的时候,我发现必须这样写:

```go
assert.Assert(&InRange{a, b, v}, t)
```

很明显`assert.Assert`这种写法挺拉杂,为了最小化打字数量,就改名为`et`了,其含义是`Easy Test`.



