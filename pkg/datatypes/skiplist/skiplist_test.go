package skiplist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipList_NewSkipList(t *testing.T) {
	sk := NewSkipList(4, 0.5)
	assert.Equal(t, uint(4), sk.maxLvl)
	assert.Equal(t, 0.5, sk.p)
	assert.Equal(t, uint(0), sk.level)
	assert.NotNil(t, sk.header)
	assert.Equal(t, -1, sk.header.key)
	assert.NotNil(t, sk.header.forward)
	assert.Equal(t, 5, len(sk.header.forward))
	for i := 0; i < len(sk.header.forward); i++ {
		assert.Nil(t, sk.header.forward[i])
	}
}

func TestSkipList_RandomLevel(t *testing.T) {
	sk := NewSkipList(4, 0.5)
	var count = 0
	for i := 0; i < 100; i++ {
		if sk.randomLevel() > 0 {
			count++
		}
	}
	assert.True(t, count > 0)
}

func TestSkipList_InsertFirstKeyIntoSkipList(t *testing.T) {
	sk := NewSkipList(4, 0.5)
	sk.Insert(2)

	sk.Print()

	assert.Equal(t,2, sk.header.forward[0].key)
}

func TestSkipList_InsertInsertAndPrint(t *testing.T) {
	sk := NewSkipList(3, 0.5)

	sk.Insert(3)
	sk.Insert(6)
	sk.Insert(7)
	sk.Insert(9)
	sk.Insert(12)
	sk.Insert(19)
	sk.Insert(17)
	sk.Insert(26)
	sk.Insert(21)
	sk.Insert(25)

	assert.Equal(t, 3, sk.header.forward[0].key)
	assert.Equal(t, 6, sk.header.forward[0].forward[0].key)
	assert.Equal(t, 7, sk.header.forward[0].forward[0].forward[0].key)
	assert.Equal(t, 9, sk.header.forward[0].forward[0].forward[0].forward[0].key)
	assert.Equal(t, 12, sk.header.forward[0].forward[0].forward[0].forward[0].forward[0].key)
	assert.Equal(t, 17, sk.header.forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].key)
	assert.Equal(t, 19, sk.header.forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].key)
	assert.Equal(t, 21, sk.header.forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].key)
	assert.Equal(t, 25, sk.header.forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].key)
	assert.Equal(t, 26, sk.header.forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].forward[0].key)

	sk.Print()
}

func TestSkipList_Search(t *testing.T)  {
	sk := NewSkipList(3, 0.5)

	assert.False(t, sk.Contains(3))

	sk.Insert(3)
	sk.Insert(6)
	sk.Insert(7)
	sk.Insert(9)
	sk.Insert(12)
	sk.Insert(19)
	sk.Insert(17)
	sk.Insert(26)
	sk.Insert(21)
	sk.Insert(25)

	assert.True(t, sk.Contains(3))
	assert.True(t, sk.Contains(26))
	assert.True(t, sk.Contains(17))
	assert.False(t, sk.Contains(5))
}