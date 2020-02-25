package cmd

import (
	"fmt"
	"github.com/andrewkandzuba/topcoder/pkg/wiki"
	"io/ioutil"
)

func main() {

	var tempDir, _ = ioutil.TempDir("", "wiki")

	p1 := &wiki.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	_ = p1.Save(tempDir)
	p2 := wiki.LoadPage(tempDir, "TestPage")
	fmt.Println(string(p2.Body))
}
