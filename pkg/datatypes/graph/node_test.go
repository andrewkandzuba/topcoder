package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := NewNode(1)
	assert.NotNil(t, n)
	assert.Equal(t, 1, n.id)
	assert.NotNil(t, n.children)
	assert.Equal(t, 0, len(n.children))
}

func TestNode_AddChild(t *testing.T) {
	n := NewNode(1)
	n.AddChild(NewNode(2))
	assert.Equal(t, 1, len(n.children))
	assert.NotNil(t, n.children[0])
	assert.Equal(t, 1, n.children[0].id)
	assert.Equal(t, 0, len(n.children[0].children))
}