package examples

import (
	"testing"
	"github.com/tinyhubs/et/assert"
)

func Test_1_Assert_Equal(t *testing.T) {
	assert.Equal(t, "123", "456")
}

func Test_1_Assert_Equali(t *testing.T) {
	assert.Equali(t, "Expect-the-values-is-equal", "123", "456")
}



