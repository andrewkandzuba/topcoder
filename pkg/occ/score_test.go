package occ

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintSorted(t *testing.T) {
	PrintSorted([]string{"ZAS", "BCD", "AAA", "AAA", "ABA","ABAC"})
}

func TestScore(t *testing.T) {
	score := Score([]string{"MARY","PATRICIA","LINDA","BARBARA","VINCENZO","SHON","LYNWOOD","JERE","HAI"})
	assert.Equal(t, 3194, score)
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Score([]string{"MARY","PATRICIA","LINDA","BARBARA","VINCENZO","SHON","LYNWOOD","JERE","HAI"})
	}
}
