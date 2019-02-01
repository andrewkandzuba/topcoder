package datatypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestListNode(t *testing.T) {
	c := NewSkipListEntry(0, (*SkipListEntry)(nil), (*SkipListEntry)(nil))

	assert.Equal(t, (*SkipListEntry)(nil), c.next, "Should be nil")
	assert.Equal(t, 0, c.value, "Should be equal")

	n := SkipListEntry{1, nil, nil}
	c = &SkipListEntry{0, nil, nil}

	assert.Equal(t, &n, c.next, "Should be equal")
	assert.Equal(t, 0, c.value, "Should be equal")
	assert.Equal(t, 1, c.next.value, "Should be equal")
}
