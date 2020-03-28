package trie

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIndex_Insert(t *testing.T) {

	index := NewIndex()
	index.Insert("abc")
	index.Insert("abb")
	index.Insert("dbs")
	index.Insert("aaa")
	index.Insert("zzz")

	assert.True(t, true)
}

func TestIndex_Search(t *testing.T) {
	index := NewIndex()
	index.Insert("abc")
	index.Insert("abb")
	index.Insert("dbs")
	index.Insert("aaa")
	index.Insert("zzz")

	assert.True(t, index.Search("abc"))
	assert.True(t, index.Search("zzz"))
	assert.False(t, index.Search("aa"))
	assert.False(t, index.Search("ab"))
	assert.False(t, index.Search("abcde"))
}

func TestIndex_SearchPrefix(t *testing.T) {
	index := NewIndex()
	index.Insert("abc")
	index.Insert("abb")
	index.Insert("dbs")
	index.Insert("aaa")
	index.Insert("zzz")

	assert.True(t, index.SearchPrefix("abc"))
	assert.True(t, index.SearchPrefix("zzz"))
	assert.True(t, index.SearchPrefix("aa"))
	assert.True(t, index.SearchPrefix("ab"))
	assert.False(t, index.SearchPrefix("abcde"))
	assert.False(t, index.SearchPrefix("ae"))
}

func BenchmarkIndex_Insert(b *testing.B) {
	input := "\"MARY\",\"PATRICIA\",\"LINDA\",\"BARBARA\",\"VINCENZO\",\"SHON\",\"LYNWOOD\",\"JERE\",\"HAI\""
	arr := strings.Split(input,",")
	for i:=0; i<len(arr); i++{
		arr[i]=strings.ToLower(strings.Trim(arr[i], "\""))
	}
	index := NewIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i:=0; i<len(arr); i++ {
			index.Insert(arr[i])
		}
	}
}


