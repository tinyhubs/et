package examples

import (
	"testing"
	"github.com/tinyhubs/et/expect"
	"github.com/tinyhubs/et/assert"
)

func Test_2_Expect_Equal(t *testing.T) {
	expect.Equal(t, "123", "456")
	expect.Equal(t, 333, 7788)
	expect.Equal(t, "sina", "google")
} //  stoped here

func Test_2_Assert_Equal(t *testing.T) {
	assert.Equal(t, "mtn", "rnd") //  stoped here
	assert.Equal(t, 444, 3366)
	assert.Equal(t, "timo", "bilibili")
}
