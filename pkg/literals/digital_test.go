package literals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	assert.Equal(t, 0, (&Digital{"111"}).compare(Digital{"111"}), "Should be 0")
	assert.Equal(t, -1, (&Digital{"11"}).compare(Digital{"111"}), "Should be -1")
	assert.Equal(t, 1, (&Digital{"111"}).compare(Digital{"11"}), "Should be 1")
	assert.Equal(t, -1, (&Digital{"111"}).compare(Digital{"112"}), "Should be -1")
	assert.Equal(t, 1, (&Digital{"112"}).compare(Digital{"111"}), "Should be 1")
}