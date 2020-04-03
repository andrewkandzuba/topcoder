package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortAndCompress(t *testing.T) {
	assert.Equal(t, []int{0}, SortAndCompress([]int{0}))
	assert.Equal(t, []int{0,1}, SortAndCompress([]int{1,0}))
	assert.Equal(t, []int{0,1,2}, SortAndCompress([]int{2,0,1}))
	assert.Equal(t, []int{0, 1, 3, 4, 7}, SortAndCompress([]int{1, 0, 3, 7, 4, 3, 0}))
	assert.Equal(t, []int{0, 1}, SortAndCompress([]int{0,0,0,1,1,0,0,1,0,0,0}))
}