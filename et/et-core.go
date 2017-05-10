package et

import (
	"runtime"
	"testing"
)

type Assertor interface {
	Assert() error
}

func AssertInner(t *testing.T, message string, assertor Assertor, callerSkip int) {
	if err := assertor.Assert(); nil != err {
		_, file, line, _ := runtime.Caller(callerSkip)
		if "" == message {
			t.Errorf("\n%s:%d\n%s\n", file, line, err.Error())
		} else {
			t.Errorf("\n%s:%d\n%s\n%s\n", file, line, message, err.Error())
		}
		t.FailNow()
	}
}

func ExpectInner(t *testing.T, message string, assertor Assertor, callerSkip int) {
	if err := assertor.Assert(); nil != err {
		_, file, line, _ := runtime.Caller(callerSkip)
		if "" == message {
			t.Errorf("\n%s:%d\n%s\n", file, line, err.Error())
		} else {
			t.Errorf("\n%s:%d\n%s\n%s\n", file, line, message, err.Error())
		}
		t.Fail()
	}
}

func Assert(assertor Assertor, t *testing.T) {
	AssertInner(t, "", assertor, 2)
}


func Asserti(assertor Assertor, t *testing.T, message string) {
	AssertInner(t, message, assertor, 2)
}

func Expect(assertor Assertor, t *testing.T) {
	ExpectInner(t, "", assertor, 2)
}

func Expecti(assertor Assertor, t *testing.T, message string) {
	ExpectInner(t, message, assertor, 2)
}
