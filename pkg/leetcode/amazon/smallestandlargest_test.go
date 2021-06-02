package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindSecondSmallestAndLargest_Case1(t *testing.T) {
	min, max := findSecondSmallestAndLargest([]int{12, 13, 1, 10, 34, 1})
	assert.EqualValues(t, min, 1)
	assert.EqualValues(t, max, 13)
}

func Test_FindSecondSmallestAndLargest_Case2(t *testing.T) {
	min, max := findSecondSmallestAndLargest([]int{2, 1, 5, 3})
	assert.EqualValues(t, min, 2)
	assert.EqualValues(t, max, 3)
}

func Test_FindSecondSmallestAndLargest_Case3(t *testing.T) {
	min, max := findSecondSmallestAndLargest([]int{1})
	assert.EqualValues(t, min, 1)
	assert.EqualValues(t, max, 1)
}

func Test_FindSecondSmallestAndLargest_Case4(t *testing.T) {
	min, max := findSecondSmallestAndLargest([]int{1, 1, 1, 1, 1, 0, 1, 1})
	assert.EqualValues(t, min, 1)
	assert.EqualValues(t, max, 1)
}