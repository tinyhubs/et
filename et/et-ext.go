package et

import (
	"fmt"
	"reflect"
	"regexp"
)

type Equal struct {
	Expect interface{}
	Actual interface{}
}

func (r *Equal) Assert() error {
	if !reflect.DeepEqual(r.Expect, r.Actual) {
		return fmt.Errorf("Expect:%v, Actual:%v", r.Expect, r.Actual)
	}

	return nil
}

type NotEqual struct {
	Expect interface{}
	Actual interface{}
}

func (r *NotEqual) Assert() error {
	if reflect.DeepEqual(r.Expect, r.Actual) {
		return fmt.Errorf("Expect:%v, Actual:%v", r.Expect, r.Actual)
	}

	return nil
}

type True struct {
	Actual bool
}

func (r *True) Assert() error {
	if true != r.Actual {
		return fmt.Errorf("Expect:%v, Actual:%v", true, r.Actual)
	}

	return nil
}

type False struct {
	Actual bool
}

func (r *False) Assert() error {
	if false != r.Actual {
		return fmt.Errorf("Expect:%v, Actual:%v", false, r.Actual)
	}

	return nil
}

type Panic struct {
	F func()
}

func (r *Panic) Assert() (err error) {
	// 先对 err 赋值,占据一个位置
	err = fmt.Errorf("")

	// 如果fn抛出panic,那么逻辑会进入这里
	defer func() {
		recover()
		if nil == err {
			err = fmt.Errorf("Expect panic, but no panic catched")
		}
	}()

	r.F()

	// 如果程序的逻辑走到这里说明没有碰到任何panic
	err = nil
	return
}

type NoPanic struct {
	F func()
}

func (r *NoPanic) Assert() (err error) {
	// 先对 err 赋值,占据一个位置
	err = fmt.Errorf("")

	// 如果fn抛出panic,那么逻辑会进入这里
	defer func() {
		ret := recover()
		if nil != err {
			err = fmt.Errorf("Expect no panic, but panic catched:%v", ret)
		}
	}()

	r.F()

	err = nil
	return
}

type Match struct {
	Regexp string
	Actual string
}

func (r *Match) Assert() error {
	regex, _ := regexp.Compile(r.Regexp)
	if !regex.MatchString(r.Actual) {
		return fmt.Errorf("Expect match:`%s`, but actual `%s`", r.Regexp, r.Actual)
	}

	return nil
}

type NotMatch struct {
	Regexp string
	Actual string
}

func (r *NotMatch) Assert() error {
	regex, _ := regexp.Compile(r.Regexp)
	if regex.MatchString(r.Actual) {
		return fmt.Errorf("Expect not match:`%s`, but actual `%s`", r.Regexp, r.Actual)
	}

	return nil
}

type Nil struct {
	Actual interface{}
}

func (r *Nil) Assert() error {
	if nil != r.Actual {
		return fmt.Errorf("Expect nil, but actual not nil:%v", r.Actual)
	}

	return nil
}

type NotNil struct {
	Actual interface{}
}

func (r *NotNil) Assert() error {
	if nil != r.Actual {
		return fmt.Errorf("Expect not nil, but actual nil")
	}

	return nil
}
