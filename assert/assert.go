package assert

import (
	"github.com/tinyhubs/et/et"
	"testing"
)

// PassValue is used to check if exp equals to got.
func Equal(t *testing.T, exp, got interface{}) {
	et.AssertInner(t, "", &et.Equal{exp, got}, 2)
}

func Equali(t *testing.T, msg string, exp, got interface{}) {
	et.AssertInner(t, msg, &et.Equal{exp, got}, 2)
}

// NotEqual is used to check if exp is not equals to got
func NotEqual(t *testing.T, exp, got interface{}) {
	et.AssertInner(t, "", &et.NotEqual{exp, got}, 2)
}

// NotEqual is used to check if exp is not equals to got
func NotEquali(t *testing.T, msg string, exp, got interface{}) {
	et.AssertInner(t, msg, &et.NotEqual{exp, got}, 2)
}

// True is used to check the got be true.
func True(t *testing.T, got bool) {
	et.AssertInner(t, "", &et.True{got}, 2)
}

// True is used to check the got be true.
func Truei(t *testing.T, msg string, got bool) {
	et.AssertInner(t, msg, &et.True{got}, 2)
}

// False is used to check the got be false.
func False(t *testing.T, got bool) {
	et.AssertInner(t, "", &et.False{got}, 2)
}

// False is used to check the got be false.
func Falsei(t *testing.T, msg string, got bool) {
	et.AssertInner(t, msg, &et.False{got}, 2)
}

// Panic is used to check the fn should give a panic.
func Panic(t *testing.T, fn func()) {
	et.AssertInner(t, "", &et.Panic{fn}, 2)
}

// Panic is used to check the fn should give a panic.
func Panici(t *testing.T, msg string, fn func()) {
	et.AssertInner(t, msg, &et.Panic{fn}, 2)
}

// NoPanic is used to check the fn should not give a panic.
func NoPanic(t *testing.T, fn func()) {
	et.AssertInner(t, "", &et.NoPanic{fn}, 2)
}

// NoPanic is used to check the fn should not give a panic.
func NoPanici(t *testing.T, msg string, fn func()) {
	et.AssertInner(t, msg, &et.NoPanic{fn}, 2)
}

// Match is used to check the got is match to the regular expression of exp.
func Match(t *testing.T, regex string, got string) {
	et.AssertInner(t, "", &et.Match{regex, got}, 2)
}

// Match is used to check the got is match to the regular expression of exp.
func Matchi(t *testing.T, msg string, regex string, got string) {
	et.AssertInner(t, msg, &et.Match{regex, got}, 2)
}

func NotMatch(t *testing.T, regex string, got string) {
	et.AssertInner(t, "", &et.NotMatch{regex, got}, 2)
}

func NotMatchi(t *testing.T, msg string, regex string, got string) {
	et.AssertInner(t, msg, &et.NotMatch{regex, got}, 2)
}

func Nil(t *testing.T, got interface{}) {
	et.AssertInner(t, "", &et.Nil{got}, 2)
}

func Nili(t *testing.T, msg string, got interface{}) {
	et.AssertInner(t, msg, &et.Nil{got}, 2)
}

func NotNil(t *testing.T, got interface{}) {
	et.AssertInner(t, "", &et.NotNil{got}, 2)
}

func NotNili(t *testing.T, msg string, got interface{}) {
	et.AssertInner(t, msg, &et.NotNil{got}, 2)
}
