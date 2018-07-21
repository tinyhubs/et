package assert

import (
	"github.com/tinyhubs/et/et"
	"testing"
)

// Equal is used to check if exp equals to got.
func Equal(t *testing.T, exp, got interface{}) {
	et.AssertInner(t, "", &et.Equal{exp, got}, 2)
}

// Equali is same with Equal but a need msg to express your intention.
func Equali(t *testing.T, msg string, exp, got interface{}) {
	et.AssertInner(t, msg, &et.Equal{exp, got}, 2)
}

// NotEqual is used to check if exp is not equals to got
func NotEqual(t *testing.T, exp, got interface{}) {
	et.AssertInner(t, "", &et.NotEqual{exp, got}, 2)
}

// NotEquali is same with NotEqual but a need msg to express your intention.
func NotEquali(t *testing.T, msg string, exp, got interface{}) {
	et.AssertInner(t, msg, &et.NotEqual{exp, got}, 2)
}

// True is used to check the got be true.
func True(t *testing.T, got bool) {
	et.AssertInner(t, "", &et.True{got}, 2)
}

// Truei is same with True but a need msg to express your intention.
func Truei(t *testing.T, msg string, got bool) {
	et.AssertInner(t, msg, &et.True{got}, 2)
}

// False is used to check the got be false.
func False(t *testing.T, got bool) {
	et.AssertInner(t, "", &et.False{got}, 2)
}

// Falsei is same with False but a need msg to express your intention.
func Falsei(t *testing.T, msg string, got bool) {
	et.AssertInner(t, msg, &et.False{got}, 2)
}

// Panic is used to check the fn should give a panic.
func Panic(t *testing.T, fn func()) {
	et.AssertInner(t, "", &et.Panic{fn}, 2)
}

// Panici is same with Panic but a need msg to express your intention.
func Panici(t *testing.T, msg string, fn func()) {
	et.AssertInner(t, msg, &et.Panic{fn}, 2)
}

// NoPanic is used to check the fn should not give a panic.
func NoPanic(t *testing.T, fn func()) {
	et.AssertInner(t, "", &et.NoPanic{fn}, 2)
}

// NoPanici is same with NoPanic but a need msg to express your intention.
func NoPanici(t *testing.T, msg string, fn func()) {
	et.AssertInner(t, msg, &et.NoPanic{fn}, 2)
}

// Match is used to check the got is match to the regular expression of exp.
func Match(t *testing.T, regex string, got string) {
	et.AssertInner(t, "", &et.Match{regex, got}, 2)
}

// Matchi is same with Match but a need msg to express your intention.
func Matchi(t *testing.T, msg string, regex string, got string) {
	et.AssertInner(t, msg, &et.Match{regex, got}, 2)
}

// NotMatch is used to check the got be not matched with exp.
func NotMatch(t *testing.T, regex string, got string) {
	et.AssertInner(t, "", &et.NotMatch{regex, got}, 2)
}

// NotMatchi is same with NotMatch but a need msg to express your intention.
func NotMatchi(t *testing.T, msg string, regex string, got string) {
	et.AssertInner(t, msg, &et.NotMatch{regex, got}, 2)
}

// Nil expect the got be nil.
func Nil(t *testing.T, got interface{}) {
	et.AssertInner(t, "", &et.Nil{got}, 2)
}

// Nili is same with NotMatch but a need msg to express your intention.
func Nili(t *testing.T, msg string, got interface{}) {
	et.AssertInner(t, msg, &et.Nil{got}, 2)
}

// NotNil expect the got be not nil.
func NotNil(t *testing.T, got interface{}) {
	et.AssertInner(t, "", &et.NotNil{got}, 2)
}

// NotNili is same with NotNil but a need msg to express your intention.
func NotNili(t *testing.T, msg string, got interface{}) {
	et.AssertInner(t, msg, &et.NotNil{got}, 2)
}
