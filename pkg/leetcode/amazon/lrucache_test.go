package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LRUCache_Case1(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}

	assert.EqualValues(t, 1, lRUCache.Get(1))    // return 1

	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}

	assert.EqualValues(t, -1, lRUCache.Get(2))    // returns -1 (not found)

	lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}

	assert.EqualValues(t, -1, lRUCache.Get(1))    // return -1 (not found)
	assert.EqualValues(t, 3, lRUCache.Get(3))    // return 3
	assert.EqualValues(t, 4, lRUCache.Get(4))    // return 4
}

func Test_LRUCache_Case2(t *testing.T) {
	lRUCache := Constructor(1)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}

	assert.EqualValues(t, -1, lRUCache.Get(1))    // return 1
	assert.EqualValues(t, 2, lRUCache.Get(2))    // return 1
}

func Test_LRUCache_Case3(t *testing.T) {
	//["LRUCache","get","put","get","put","put","get","get"]
	//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
    //[null,-1,null,-1,null,null,2,-1]
	lRUCache := Constructor(2)

	assert.EqualValues(t, -1, lRUCache.Get(2))

	lRUCache.Put(2, 6) // cache is {1=1}

	assert.EqualValues(t, -1, lRUCache.Get(1))

	lRUCache.Put(1, 5)
	lRUCache.Put(1, 2)

	assert.EqualValues(t, 2, lRUCache.Get(1))
	assert.EqualValues(t, 6, lRUCache.Get(2))
}

func Test_LRUCache_Case4(t *testing.T) {

   //["LRUCache","put","put","put","put","get","get"]
//	[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]
	lRUCache := Constructor(2)

	lRUCache.Put(2, 1)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 3)
	lRUCache.Put(4, 1)

	assert.EqualValues(t, -1, lRUCache.Get(1))
	assert.EqualValues(t, 3, lRUCache.Get(2))
}

func Test_LRUCache_Case5(t *testing.T) {

	//["LRUCache","put","put","put","put","get","get"]
	//	[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]
	lRUCache := Constructor(1)

	assert.EqualValues(t, -1, lRUCache.Get(1))

	lRUCache.Put(2, 1)
	lRUCache.Put(1, 1)

	assert.EqualValues(t, -1, lRUCache.Get(2))
	assert.EqualValues(t, 1, lRUCache.Get(1))
	assert.EqualValues(t, -1, lRUCache.Get(2))


	lRUCache.Put(2, 3)
	lRUCache.Put(4, 1)

	assert.EqualValues(t, 1, lRUCache.Get(4))
	assert.EqualValues(t, -1, lRUCache.Get(2))
}