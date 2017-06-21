package username

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {

	assert.EqualValues(t, map[string]int{
		"MasterOfDisaster": 1,
		"DingBat" : 1,
		"Orpheus" : 1,
		"WolfMan" : 1,
		"MrKnowItAll" : 1,
	}, NewDatabase([]string{"MasterOfDisaster", "DingBat", "Orpheus", "WolfMan", "MrKnowItAll"}).names, "Should be equal");

	assert.NotEqual(t, map[string]int{}, NewDatabase([]string{"MasterOfDisaster", "DingBat", "Orpheus", "WolfMan", "MrKnowItAll"}).names, "Should not be equal");

	assert.NotEqual(t, map[string]int{
		"MasterOfDisaster": 1,
		"DingBat" : 1,
		"Orpheus" : 1,
		"WolfMan" : 1,
		"MrKnowItAll" : 1,
	}, NewDatabase([]string{"MasterOfDisaster", "MasterOfDisaster12", "DingBat4", "Orpheus", "WolfMan", "MrKnowItAll"}).names, "Should not be equal");
}
