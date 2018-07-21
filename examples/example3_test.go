package examples

import (
	"fmt"
	"testing"
	"github.com/tinyhubs/et/et"
)

type Inrange struct {
	Min   int
	Max   int
	Value int
}

func (i *Inrange) Assert() error {
	if (i.Value >= i.Min) && (i.Value <= i.Max) {
		return nil
	}

	return fmt.Errorf("expect in range [%v, %v], Actual: %v", i.Min, i.Max, i.Value)
}

func Test_mycase1(t *testing.T) {
	min := 1
	max := 100
	value := 320
	et.Assert(&Inrange{min, max, value}, t)
}

func AssertInrange(t *testing.T, min int, max int, value int) {
	et.AssertInner(t, "", &Inrange{min, max, value}, 2)
}

func Test_mycase2(t *testing.T) {
	min := 1
	max := 100
	value := 320
	AssertInrange(t, min, max, value)
}
