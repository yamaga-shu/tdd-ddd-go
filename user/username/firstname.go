package username

import (
	"errors"
	"regexp"
)

type firstName struct {
	value string
}

func NewFirstName(val string) (*firstName, error) {
	// validate length
	if len(val) < 3 {
		return nil, errors.New("firstName must be at least 3 characters long")
	}

	// validate characters by regular expression
	validFirstName := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !validFirstName.MatchString(val) {
		return nil, errors.New("firstName can only contain alphabetic characters, numbers, or hyphens")
	}

	return &firstName{value: val}, nil
}
