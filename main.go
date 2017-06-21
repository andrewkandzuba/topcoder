package main

import u "topcoder/username"

func main() {

	u.PrintNewName("test")

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
	u.PrintNewName(longName)
}

