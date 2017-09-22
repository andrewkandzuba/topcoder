package datatypes


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitializeLane(t *testing.T) {
	header := Lane{key:30, forward:[]Lane{}}

	assert.Equal(t, 30, header.key, "Should be equal")
	assert.Equal(t, 0, len(header.forward), "Should be equal")

	forward := Lane{key:40, forward:[]Lane{}}
	header.forward = append(header.forward, forward)

	assert.Equal(t, 30, header.key, "Should be equal")
	assert.Equal(t, 1, len(header.forward), "Should be equal")
	assert.Equal(t, 40, header.forward[0].key, "Should be equal")
	assert.Equal(t, 0, len(header.forward[0].forward), "Should be equal")
}
