package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SearchInRotatedSortedArray(t *testing.T) {
	assert.EqualValues(t, searchInArray([]int{0}, 1), -1)
	assert.EqualValues(t, searchInArray([]int{1}, 1), 0)
	assert.EqualValues(t, searchInArray([]int{4,5,6,7,0,1,2}, 0), 4)
	assert.EqualValues(t, searchInArray([]int{4,5,6,7,0,1,2}, 3), -1)
}
