package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LadderLength(t *testing.T) {
	assert.EqualValues(t, 5, ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log","cog"}))
	assert.EqualValues(t, 0, ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log"}))
	assert.EqualValues(t, 5, ladderLength("hit", "cog", []string{"hat","hot","dot","dog","lot","log","cog"}))
}

func Test_differenceInCharacters(t *testing.T)  {
	assert.EqualValues(t, diffInOneChar("hit", "cog"), false)
	assert.EqualValues(t, diffInOneChar("hit", "hot"), true)
	assert.EqualValues(t, diffInOneChar("hit", "hit"), false)
}