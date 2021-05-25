package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CountAndSay(t *testing.T) {
	assert.EqualValues(t, "1", countAndSay(1))
	assert.EqualValues(t, "1211", countAndSay(4))
	assert.EqualValues(t, "1211", countAndSay(30))
}
