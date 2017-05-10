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
