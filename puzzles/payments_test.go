package puzzles

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestFind(t *testing.T) {

	incoming := []int{100, 100, 225, 300, 473, 80}

	assert.Equal(t, find(incoming, 180), []int{100, 80})
	assert.Equal(t, find(incoming, 773), []int{300, 473})
	assert.Equal(t, find(incoming, 225), []int{225})
	assert.Equal(t, find(incoming, 325), []int{100, 225})
	assert.Equal(t, find(incoming, 335), []int{})
	assert.Equal(t, find(incoming, 573), []int{100, 473})
}

