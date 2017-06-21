package username

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuildDB(t *testing.T) {

	assert.EqualValues(t, map[string]int{
		"MasterOfDisaster": 1,
		"DingBat" : 1,
		"Orpheus" : 1,
		"WolfMan" : 1,
		"MrKnowItAll" : 1,
	}, BuildDB([]string{"MasterOfDisaster", "DingBat", "Orpheus", "WolfMan", "MrKnowItAll"}), "Should be equal");

	assert.NotEqual(t, map[string]int{}, BuildDB([]string{"MasterOfDisaster", "DingBat", "Orpheus", "WolfMan", "MrKnowItAll"}), "Should not be equal");

	assert.NotEqual(t, map[string]int{
		"MasterOfDisaster": 1,
		"DingBat" : 1,
		"Orpheus" : 1,
		"WolfMan" : 1,
		"MrKnowItAll" : 1,
	}, BuildDB([]string{"MasterOfDisaster", "MasterOfDisaster12", "DingBat4", "Orpheus", "WolfMan", "MrKnowItAll"}), "Should not be equal");
}

func TestNewMember(t *testing.T) {
/*	u.PrintNewName("test")

	u.PrintNewName("test2")

	u.PrintNewName("")

	var shortName string
	for i := 0; i < 50; i++ {
		shortName += "a";
	}
	u.PrintNewName(shortName)

	var longName string
	for i := 0; i < 51; i++ {
		longName += "a";
	}
	u.PrintNewName(longName)*/
}
