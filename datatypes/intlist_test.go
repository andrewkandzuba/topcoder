package datatypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntListNode(t *testing.T) {
	c := IntListNode{v: 0, n: nil}

	assert.Equal(t, nil, c.n, "Should be nil")
	assert.Equal(t, 0, c.v, "Should be equal")

	n := IntListNode{v: 1, n: nil}
	c = IntListNode{v: 0, n: n}
	assert.Equal(t, n, c.n, "Should be equal")
	assert.Equal(t, 0, c.v, "Should be equal")
	//assert.Equal(t, 1, c.n.v, "Should be equal")
}
