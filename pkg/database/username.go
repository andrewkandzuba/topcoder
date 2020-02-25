package username

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	UserNameInvalidEnerrorMessage      = "User name length is invalid"
	UserNameInvalidContentErrorMessage = "User name contains illegal symbols"
)

func PrintNewName(newName string) {

	r, err := NewMember(newName)
	if err != nil {
		fmt.Printf("\"%s\" is not available: \"%s\"\n", newName, err)
		return
	}
	fmt.Printf("\"%s\" is available\n", r)
}

func NewMember(newName string) (string, error) {

	length := len(newName)
	if length < 1 || length > 50 {
		return "", errors.New(UserNameInvalidEnerrorMessage)
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(newName) {
		return "", errors.New(UserNameInvalidContentErrorMessage)
	}

	return newName, nil
}