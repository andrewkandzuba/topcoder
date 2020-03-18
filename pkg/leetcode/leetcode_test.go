package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSums(t *testing.T) {
	var p = Problems{}
	assert.EqualValues(t, []int{0, 1}, p.twoSum([]int{2, 7, 11, 15}, 9))
	assert.EqualValues(t, []int{0, 6}, p.twoSum([]int{2,7,11,15,0,0,1}, 3))
	//assert.Equal(t, 0.0, p.findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 5.5, p.findMedianSortedArrays([]int{3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 9, 10}))
}
