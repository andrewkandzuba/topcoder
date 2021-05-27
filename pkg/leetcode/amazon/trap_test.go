package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trap(t *testing.T) {
	assert.EqualValues(t, 6, trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	assert.EqualValues(t, 9, trap([]int{4,2,0,3,2,5}))
}