package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isRobotBounded(t *testing.T) {
	assert.EqualValues(t, true, isRobotBounded("LL"))
	assert.EqualValues(t, true, isRobotBounded("RRRRLLLL"))
	assert.EqualValues(t, true, isRobotBounded("GGLLGG"))
	assert.EqualValues(t, false, isRobotBounded("GG"))
	assert.EqualValues(t, true, isRobotBounded("GL"))
	assert.EqualValues(t, true, isRobotBounded("RRRLLLGLLG"))
}