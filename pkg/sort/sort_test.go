package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	assert.Equal(t, []int{0,1,2,3,3,4,7}, MergeSort([]int{1, 0,3, 7, 4, 3, 0}))
}

