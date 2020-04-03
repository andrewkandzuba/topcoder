package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	assert.Equal(t, []int{0}, MergeSort([]int{0}))
	assert.Equal(t, []int{0,1}, MergeSort([]int{1,0}))
	assert.Equal(t, []int{0,1,2}, MergeSort([]int{2,0,1}))
	assert.Equal(t, []int{0, 0, 1, 3, 3, 4, 7}, MergeSort([]int{1, 0, 3, 7, 4, 3, 0}))
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{1, 0, 3, 7, 4, 3, 0})
	}
}
