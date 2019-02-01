package username

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestNewMember(t *testing.T) {

	PrintNewName("test")

	PrintNewName("test2")

	PrintNewName("")

	var shortName string
	for i := 0; i < 50; i++ {
		shortName += "a";
	}
	PrintNewName(shortName)

	var longName string
	for i := 0; i < 51; i++ {
		longName += "a";
	}
	PrintNewName(longName)
}
