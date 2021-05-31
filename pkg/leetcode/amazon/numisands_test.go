package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NumIslands_Case1(t *testing.T) {
	assert.EqualValues(t, 1, numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
}

func Test_NumIslands_Case2(t *testing.T) {
	assert.EqualValues(t, 3, numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}))
}

func Test_NumIslands_Case3(t *testing.T) {
	assert.EqualValues(t, 2, numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
	}))
}

func Test_NumIslands_Case4(t *testing.T) {
	assert.EqualValues(t, 1, numIslands([][]byte{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}))
}

func Test_NumIslands_Case5(t *testing.T) {
	assert.EqualValues(t, 1, numIslands([][]byte{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}))
}

func Test_NumIslands_Case6(t *testing.T) {
	assert.EqualValues(t, 1, numIslands([][]byte{
		{1},
	}))
}

func Test_NumIslands_Case7(t *testing.T) {
	assert.EqualValues(t, 1, numIslands([][]byte{
		{1},
		{1},
	}))
}