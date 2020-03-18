package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSums(t *testing.T) {
	assert.EqualValues(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.EqualValues(t, []int{0, 6}, twoSum([]int{2, 7, 11, 15, 0, 0, 1}, 3))
}

func Test_MedianOfTwoSortedArrays(t *testing.T) {
	assert.Equal(t, 1.5, findMedianSortedArrays([]int{1}, []int{2}))

	assert.Equal(t, 2.0, findMedianSortedArrays([]int{2}, []int{1, 3}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{5}, []int{1, 3}))
	assert.Equal(t, 7.0, findMedianSortedArrays([]int{7}, []int{4, 8}))

	assert.Equal(t, 4.5, findMedianSortedArrays([]int{1}, []int{2, 7, 12}))
	assert.Equal(t, 5.5, findMedianSortedArrays([]int{4}, []int{2, 7, 12}))
	assert.Equal(t, 6.0, findMedianSortedArrays([]int{5}, []int{2, 7, 12}))
	assert.Equal(t, 10.0, findMedianSortedArrays([]int{19}, []int{2, 7, 13}))

	assert.Equal(t, 7.0, findMedianSortedArrays([]int{1}, []int{2, 7, 12, 15}))
	assert.Equal(t, 7.0, findMedianSortedArrays([]int{4}, []int{2, 7, 12, 15}))
	assert.Equal(t, 9.0, findMedianSortedArrays([]int{9}, []int{2, 7, 12, 15}))
	assert.Equal(t, 12.0, findMedianSortedArrays([]int{13}, []int{2, 7, 12, 15}))
	assert.Equal(t, 12.0, findMedianSortedArrays([]int{18}, []int{2, 7, 12, 15}))

	assert.Equal(t, 5.0, findMedianSortedArrays([]int{1}, []int{2, 3, 7, 12, 15}))
	assert.Equal(t, 5.5, findMedianSortedArrays([]int{4}, []int{2, 3, 7, 12, 15}))
	assert.Equal(t, 10.5, findMedianSortedArrays([]int{9}, []int{2, 7, 12, 11, 15}))
	assert.Equal(t, 12.5, findMedianSortedArrays([]int{13}, []int{2, 7, 12, 14, 15}))
	assert.Equal(t, 13.5, findMedianSortedArrays([]int{18}, []int{2, 7, 12, 15, 17}))

	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{1, 3}, []int{5}))
	assert.Equal(t, 7.0, findMedianSortedArrays([]int{4, 8}, []int{7}))

	assert.Equal(t, 6.0, findMedianSortedArrays([]int{4, 8}, []int{5, 7}))
	assert.Equal(t, 4.0, findMedianSortedArrays([]int{1, 2}, []int{6, 7}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{1, 4}, []int{2, 7}))
	assert.Equal(t, 4.5, findMedianSortedArrays([]int{3, 6}, []int{2, 7}))
	assert.Equal(t, 8.0, findMedianSortedArrays([]int{9, 10}, []int{2, 7}))

	assert.Equal(t, 7.0, findMedianSortedArrays([]int{4, 8}, []int{5, 7, 11}))

	assert.Equal(t, 5.5, findMedianSortedArrays([]int{3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 9, 10}))

	assert.Equal(t, 3.0, findMedianSortedArrays([]int{}, []int{1, 2, 3, 9, 10}))
	assert.Equal(t, 1.0, findMedianSortedArrays([]int{}, []int{1}))
	assert.Equal(t, 1.0, findMedianSortedArrays([]int{1}, []int{}))
}
