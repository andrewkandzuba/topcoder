package datatypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListNode(t *testing.T) {
	c := ListNode{v: 0}

	assert.Equal(t, (*ListNode)(nil), c.n, "Should be nil")
	assert.Equal(t, 0, c.v, "Should be equal")

	n := ListNode{v: 1, n: nil}
	c = ListNode{v: 0, n: &n}
	assert.Equal(t, &n, c.n, "Should be equal")
	assert.Equal(t, 0, c.v, "Should be equal")
	assert.Equal(t, 1, c.n.v, "Should be equal")
}
