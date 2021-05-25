package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findKthLagest(t *testing.T) {
	assert.EqualValues(t, findKthLargest([]int{0}, 1), 0)
	assert.EqualValues(t, findKthLargest([]int{1}, 2), 1)
	assert.EqualValues(t, findKthLargest([]int{3,2,1,5,6,4}, 2), 5)
	assert.EqualValues(t, findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4), 4)
}

