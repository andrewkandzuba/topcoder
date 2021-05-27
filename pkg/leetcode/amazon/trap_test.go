package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_trap(t *testing.T) {
	assert.EqualValues(t, 6, trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	assert.EqualValues(t, 9, trap([]int{4,2,0,3,2,5}))
	assert.EqualValues(t, 1, trap([]int{4,2,3}))
	assert.EqualValues(t, 0, trap([]int{2,3}))
	assert.EqualValues(t, 0, trap([]int{1}))
	assert.EqualValues(t, 0, trap([]int{0,0,0}))
	assert.EqualValues(t, 0, trap([]int{0,1,0}))
	assert.EqualValues(t, 1, trap([]int{1,0,1}))
	assert.EqualValues(t, 0, trap([]int{1,0,0}))
	assert.EqualValues(t, 0, trap([]int{0,0,1}))
}