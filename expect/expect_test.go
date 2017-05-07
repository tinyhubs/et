package expect

import (
	"testing"
)

func TestAssert_Equal(t *testing.T) {
	Equal(t, "123", "456")
}

func TestAssert_Equali(t *testing.T) {
	Equali(t, "Expect-the-values-is-equal", "123", "456")
}

func TestAssert_NotEqual(t *testing.T) {
	NotEqual(t, "123", "123")
}

func TestAssert_NotEquali(t *testing.T) {
	NotEquali(t, "Expect-the-values-is-not-equal", "123", "123")
}

func TestAssert_True(t *testing.T) {
	True(t, "123" == "456")
}

func TestAssert_Truei(t *testing.T) {
	Truei(t, "Expect-the-expresion-is-true", "123" == "456")
}

func TestAssert_False(t *testing.T) {
	False(t, "123" == "123")
}

func TestAssert_Falsei(t *testing.T) {
	Falsei(t, "Expect-the-expresion-is-false", "123" == "123")
}

func TestAssert_Panic(t *testing.T) {
	Panic(t, func() { /* Do nothing. */ })
}

func TestAssert_Panici(t *testing.T) {
	Panici(t, "Expect-the-func-throw-a-panic", func() {
		/* Do nothing. */
	})
}

func throwPanic() {
	panic(123)
}

func TestAssert_NoPanic(t *testing.T) {
	NoPanic(t, func() {
		throwPanic()
	})
}

func TestAssert_NoPanici(t *testing.T) {
	NoPanici(t, "Expect-the-func-do-not-throw-a-panic", func() {
		throwPanic()
	})
}
