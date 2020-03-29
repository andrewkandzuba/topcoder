package trie

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIndex_Insert(t *testing.T) {

	index := NewIndex()
	index.Insert("ABC", 0)
	index.Insert("ABB", 1)
	index.Insert("DBS", 2)
	index.Insert("AAA", 3)
	index.Insert("ZZZ", 4)

	assert.True(t, true)
}

func TestIndex_findNode(t *testing.T) {
	arr := []string{"MARY","PATRICIA","LINDA","BARBARA","VINCENZO","SHON","LYNWOOD","JERE","HAI"}

	index := NewIndex()
	for i, a := range arr {
		index.Insert(a, i)
	}

	n := index.findNode("LINDA")
	assert.NotNil(t, n)
	assert.Equal(t, 40, n.Score)
}

func TestIndex_Search(t *testing.T) {
	index := NewIndex()
	index.Insert("ABC", 0)
	index.Insert("ABB", 1)
	index.Insert("DBS", 2)
	index.Insert("AAA", 3)
	index.Insert("ZZZ", 4)

	assert.True(t, index.Search("ABC"))
	assert.True(t, index.Search("ZZZ"))
	assert.False(t, index.Search("AA"))
	assert.False(t, index.Search("AB"))
	assert.False(t, index.Search("ABCDE"))
}

func TestIndex_SearchPrefix(t *testing.T) {
	index := NewIndex()
	index.Insert("ABC", 0)
	index.Insert("ABB", 1)
	index.Insert("DBS", 2)
	index.Insert("AAA",3)
	index.Insert("ZZZ", 4)

	assert.True(t, index.SearchPrefix("ABC"))
	assert.True(t, index.SearchPrefix("ZZZ"))
	assert.True(t, index.SearchPrefix("AA"))
	assert.True(t, index.SearchPrefix("AB"))
	assert.False(t, index.SearchPrefix("ABCDE"))
	assert.False(t, index.SearchPrefix("AE"))
}

func BenchmarkIndex_Insert(b *testing.B) {
	input := "\"MARY\",\"PATRICIA\",\"LINDA\",\"BARBARA\",\"VINCENZO\",\"SHON\",\"LYNWOOD\",\"JERE\",\"HAI\""
	arr := strings.Split(input,",")
	for i:=0; i<len(arr); i++{
		arr[i]=strings.Trim(arr[i], "\"")
	}
	index := NewIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i:=0; i<len(arr); i++ {
			index.Insert(arr[i], i)
		}
	}
}


