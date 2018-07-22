# `et`的设计


## 初衷：看着舒服打字少

`et`最初的设计初衷有两个：一个是我不想写那么多 if，我希望一个断言逻辑一行搞定。另外一个原因是我不想写t参数。综合来说就是想看起来舒服且码字要少。

举个例子，我当时就想这样写测试代码：

```go
assert.Equal(a, b)
```

显然，我是有 Java 或者 C 语言背景的人。在我们眼里，这里的 `assert` 其实是个对象或者是类，assert 对象内部包含了一个 t 参数。我很快就做了个原型库。基于这个库的一个完整的例子大概是这样：

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

## 支持 except

接着，我在使用我自己做的 [properties](https://github.com/tinyhubs/properties) 库和 [tinydom](https://github.com/tinyhubs/tinydom) 库写一个开发辅助工具的时候，我发现`except`对我也非常有用，这样我的用例如果有什么 bug 测试用例能够尽可能多地帮我检测出来。但是，我总不能再 new 一个 Except 对象出来吧？所以，我就做成了让 New 一次性返回两个对象的方式。下面是个例子：

```go
func TestAssert_NotEqual(t *testing.T) {
	assert, except := New(t)
	except.NotEqual("Expect-the-values-is-not-equal", 333, 444)
	assert.NotEqual("Expect-the-values-is-not-equal", "123", "123")
}
```

这个主意看起来还挺不错。因为，得益于 Go 语言的语法特点，使用者可以根据需要自己决定 New 出的对象用哪一个，也就是你可以有下面三种情况自由组合：

```go
_, except := New(t)       //  只使用except
assert, _ := New(t)       //  只使用assert
assert, except := New(t)  //  同时使用except和assert
```

## 过犹不及

前面，为支持`expect`所带来的设计变化，看起来还行。但是，我总觉得有点疙里疙瘩的----因为个方案要求在每个`Test_xxx`函数入口的地方都要`New`一个`Assert`对象甚至额外的`Except`对象。我在一个项目里面自己试用了一下。当代码用例增多之后，就开始感觉每次都要`New`这么个对象显得好蠢。其实，我之所要`New`对象就是为了少写个`t`，但是光`_, assert := New(t)`这一句敲的字都相当于敲 N 个`t`了。很显然，这种设计有点"过犹不及"了。

由于 Go 语言标准库的设计里面，`t`就一个字符，所以在一个断言调用里面，`t`字符显得并不那么突出，所以我就放弃了少写`t`这个愿望----虽然t也不爽----并重新设计了`et`库的接口，最终代码可以直接这样写：

```go
assert.Equal(t, a, b)
```

这里`assert`是包名，为了支持`except`，我单独实现了一个`except`包，所以我们也可以这样写：

```go
except.Equal(t, a, b)
```

您或许会发现，except 如果是另外一个包，那就必须写两个常常的`import`语句（这个也要打好多字呢）：

```go
import (
    "github.com/tinyhubs/et/assert"
    "github.com/tinyhubs/et/except"
)
```

关于这一点，并不需要担心。现代的 IDE 都很智能，自动会帮我写上这段 import 代码的。

## 断言注解、i系函数测试和代码的可维护性

#### 让测试代码可维护性更好

过了一段时间，我在公司写的一个程序遇到了一个 bug。解决这个 bug 后，我试图搞明白为何我的 UT 用例没有覆盖到这个场景。在分析测试用例代码时，我忽然发现有一个用例我看不懂了。原因是代码中的注释和这段注释下面的断言代码看起来有点矛盾，我搞不清楚注释和代码到底哪个是对的?

这使我意识到，我的单元测试代码中的注释慢慢出现了"年久失修"的征兆了。我修改测试代码的时候，往往不会关注代码上下文的注释。在我的 IDE 的色彩配置里面，甚至将注释设置成灰色来避免注释干扰逻辑。所以，我想是否有办法让我修改代码的时候能够对一些非常重要的断言做一些说明，而且不能是以注释的形式体现这些说明？我很自然地想到给`assert`系列函数增加一个字符串参数的办法。最后，我修改了所有的`assert`函数的接口设计，这样代码就写成这样了：

```go
assert.True(t, "获得Text对象成功", nil != text1)
assert.True(t, "获得Text对象成功", nil != text2)
assert.True(t, "获得Text对象成功", nil != text3)
assert.True(t, "全空白的Text不会被读取", nil == text4)
```

后来，我在费力地修改我维护的几个代码仓库的测试用例的过程中，我又意识到一个新问题：虽然经过前面的改造，测试代码的可维护性确实好多了，然而很多时候写断言意图（下称"断言注解"）会存在较多重复的文字描述。比如上面这段代码，其前三个 assert 的断言注解都相同。重复性的文字显得很愚蠢。将断言注解变成可选应该是个好主意，比如，下面这些代码：

```go
assert.True(t, "获取Text对象成功（下同）", nil != text1)
assert.True(t, nil != text2)
assert.True(t, nil != text3)
assert.True(t, "全空白的Text不会被读取", nil == text4)
```

然而，Go 语言并不支持函数重载，这种代码会无法通过 Go 编译器的语法检查。好在 Go 支持变参函数，所以，我们可以将这个字符串放后面：

```go
assert.True(t, nil != text1, "获得Text对象成功")
assert.True(t, nil != text2, "")
assert.True(t, nil != text3, "")
assert.True(t, nil == text4, "全空白的Text不会被读取")
```

这个方案看起来也不错！而且，利用可变参数上面的代码可以进一步精简为下面这样：

```go
assert.True(t, nil != text1, "获得Text对象成功")
assert.True(t, nil != text2)
assert.True(t, nil != text3)
assert.True(t, nil == text4, "全空白的Text不会被读取")
```

OK，很完美！

没过多久，我实现了`assert.Panic`和`assert.NoPanic`两个新函数，在写一个`assert.NoPanic`的调用样例的时候，又觉得不舒服了：

```go
assert.NoPanic(t, func() {
    throwPanic()
}, "Expect-the-func-do-not-throw-a-panic")
```

这个代码看起来很丑陋----func()后面这个字符串就像是一只苗条的猫长了一条鳄鱼的粗壮尾巴。所以，我最终还是觉得写成下面这样，代码看起来会更舒服点：

```go
assert.NoPanic(t, "Expect-the-func-do-not-throw-a-panic",
func() {
    throwPanic()
})
```

将注解放在 assert 函数的最后作为变参还有一些不好的地方，因为函数原型最后都成为了变长的了，开发人员看到assert函数的原型对于理解这个 assert 的用法会产生疑惑。

再者，我当时正在考虑设计我的 assert 库可扩展能力，如果将注解放在函数的后面会导致大家扩展出来的新的断言函数的原型会跟我的 assert 库自带的assert函数的原型不一致，这种不一致会加大大家的学习难度。

如何解决这个问题呢？Go 语言的 fmt 包以及 C 语言的库给了我启发。他们有这样的用法：

```go
fmt.Printf("%s", mystr)
```

```cpp
printf("%s", mystr);
```

他们都是在基本核心单词（"print"）的基础上增加简短的后缀（"f"）来表明函数和同系列的函数的区别的。我觉得我的 assert 库可以采用同样的机制：

```go
assert.TrueXX(t, "获得Text对象成功", nil != text1)
assert.True(t, nil != text2)
```

OK，那么这个 XX 应该使用什么呢？----`f`? No，`f`是`format`的意思，跟断言注解不沾边。还有哪些单词可以表达断言注解的语义呢？

`message`， `tips`， `notice`， `information` ...

以他们的首字母放在最后看起来是这样的：

```go
assert.Equalm(t， "Expect-the-values-is-equal"， "123"， "456")
assert.Equalt(t， "Expect-the-values-is-equal"， "123"， "456")
assert.Emptyn(t， "Expect-the-values-is-equal"， "123"， "456")
assert.Emptyi(t， "Expect-the-values-is-equal"， "123"， "456")
```

`m` `t` `n`者三个字符很容易融入到原始单词里面去，比如一言看上去似乎`qualm` `qualt` `mptyn`是一个单词。这三个字母的亲和力很强，会出现在很多单词的结尾。`Printf`里面的`f`之所以亲和力没那么强。这主要是因为`f`通常作为单词的开头，很少有单词以f结尾。而且`f`在英文中看起来很瘦长，总是跟其他的字母不搭调。而`m` `t` `n` 在单词里面见得很多，特别是以`t`和`n`为后缀的单词相当多。

看起来字母`i`是比较不错的选择，这样，`i`系列函数就诞生了。

## 扩展能力的设计

`et`最初只提供了`Equal`，`True`，`Panic`以及反逻辑的`NotEqual`，`False`，`NoPanic`这几个函数。但是，很快我发现我还需要增加判断是否为 nil 以及正则表达式匹配的断言函数。我开始意识到，虽然`assert.True`，`assert.False`可以搞定一切，但是断言的提示信息并不友好。下面这两个用例，断言逻辑其实一模一样，但是其断言错误提示却差别很大：

```go
func Test_output1(t *testing.T) {
	assert.Equal(t, 111, 222)
}

func Test_output2(t *testing.T) {
	assert.True(t, 111 == 222)
}
```

`Test_output1`的输出如下：

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example4_test.go:9
		Expect:111, Actual:222
```

`Test_output2`的输出如下：

```text
	et-core.go:16:
		/Users/llj/mygithub/src/github.com/tinyhubs/et/examples/example4_test.go:13
		Expect:true, Actual:false
```

很明显`Test_output1`的输出可以让我们很清楚断言的两个参数的值，对我们的帮助更大，`Test_output2`的输出就缺少这些信息。所以，总结来说，我们还是应该提供针对某种类型的断言的专用函数，而不是使用笼统的`assert.True`或者`assert.False`来凑合。

作为程序库(`et`)的作者，我没法预见到其他人需要什么样的断言函数，我甚至都不知道自己未来需要什么断言函数。所以，`et`库支持各种不同的断言检测逻辑是很有意义的。于是，将原库的代码拆分成了三个部分：`et`，`assert`，`expect`。其中， `et`是核心，`et-core.go`里面定义了统一的断言函数的调用形式，封装了获取代码堆栈的逻辑。并提供了一个`Assertor`接口，`et`默认支持的`Assertor`定义在`et-ext.go`文件中。`assert`包和`expect`包是两种基于`et`做的封装。基本上前面的绝大多数例子中调用的断言函数都是来自这两个包。

最先被抽象出来的是`et.AssertInner`这是转给扩展断言库的作者使用的。它主要的作用就是使用统一的格式输出断言结果。同时，也是为了扩展库的作者不必太过关心获取调用堆栈这些细节。为此，`et.AssertInner`函数有一个额外的`callerSkip`参数，用于帮助`et`库确定需要跳过多少层堆栈。如果我们提供的断言函数不是被封装很多层，那么`callerSkip`就永远是`2`。对于写断言扩展的同学来说，确定你提供给用户的是那一层的接口是非常重要的，这可以有效减少学习成本。所以，我建议永远只在`et.AssertInner`之上包装一层。

`et.Assert`，`et.Expect`者两个是随后抽象出来的。最初，我并不想提供这两个函数，但是我在自己写一些定制的断言函数的时候，我发现有些定制的断言函数的使用频率并不是很高，通用化的意义也不大，所以我不计划去做`AssertInRange`这样的二次封装。但是我又不能直接在断言代码中调用`et.AssertInner`。所以，提供`et.Assert`，`et.Expect`，`et.Asserti`，`et.Expecti`这四个函数可以让我再懒惰一点:)。

在扩展库的设计上，`et`力求平衡`使用感知`与`代码美观度`。如下者几行代码，看起来还是有点差异，但是他们的共同特点就是`assert`总是与`Equal`靠近。

```go
assert.Equal(t, a, b)
assert.Equali(t, "message", a, b)
et.Assert(&Equal{a, b}, t)
AssertEqual(t, a, b)
```

再比如，`i`系函数里面，`message`总是在`t`的后面，即使是`et.AssertInner`，`et.Asserti`也不例外。

下面列出了`et`所遵循的设计约定：

1. 一个断言检测支持三种调用形式：

- `et.Assert`  按扩展库的方式调用
- `assert.XXX` 通常用到的断言函数
- `expect.XXX` except断言模式

2. 如果断言函数的第一个参数总是 t，除了扩展模式

3. 带`i`后缀的函数都支持一个额外的字符串参数，该参数总是在`t`参数的后面。

4. `Assert`总是应该与断言的类型靠近

5. 断言函数或者`Assertor`的预期值总是在被检测值的前面


## 其他：断言库改名

`et`库最早其实叫`assert`，但是后来在支持扩展能力的时候，我发现必须这样写：

```go
assert.Assert(&InRange{a, b, v}, t)
```

很明显`assert.Assert`这种写法挺拉杂的。为了最小化打字数量，就改名为`et`了，其含义是`Easy Test`。


