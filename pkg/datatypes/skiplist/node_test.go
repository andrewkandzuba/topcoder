package skiplist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializeNone(t *testing.T) {
	var node = NewNode(10, 2)
	assert.Equal(t, 10, node.key)
	assert.Equal(t, 3, len(node.forward))
	assert.Nil(t, node.forward[0])
	assert.Nil(t, node.forward[1])
	assert.Nil(t, node.forward[2])
}