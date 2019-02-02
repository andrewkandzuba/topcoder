package wiki

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var tempDir, _ = ioutil.TempDir("", "example")

func TestPage(t *testing.T) {

	p := &Page{
		"Title",
		[]byte{'g', 'o', 'l', 'a', 'n', 'g'},
	}
	assert.EqualValues(t, "Title", p.Title, "Should be equal")
	assert.EqualValues(t, []byte{'g', 'o', 'l', 'a', 'n', 'g'}, p.Body, "Should be equal")
}

func TestWriteReadPage(t *testing.T) {
	title := "Title"
	path := tempDir + "/" + title + ".txt";

	deleteFile(path)
	v, err := isFileExists(path)

	assert.False(t, v, "File should not exist")
	assert.Nil(t, err, "No errors")

	p := &Page{
		"Title",
		[]byte{'g', 'o', 'l', 'a', 'n', 'g'},
	}

	err = p.Save(tempDir);
	assert.Nil(t, err, "No errors")

	v, err = isFileExists(path)

	assert.True(t, v, "File should exist")
	assert.Nil(t, err, "No errors")

	loaded := LoadPage(tempDir, title)

	assert.NotNil(t, loaded, "Should be loaded")
	assert.EqualValues(t, "Title", loaded.Title, "Should be equal")
	assert.EqualValues(t, []byte{'g', 'o', 'l', 'a', 'n', 'g'}, loaded.Body, "Should be equal")
}

func deleteFile(path string) {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("==> done deleting file")
}

func isFileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
