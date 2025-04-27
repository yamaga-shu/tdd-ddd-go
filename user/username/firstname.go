package username

import "errors"

type firstName struct {
	value string
}

func NewFirstName(val string) (*firstName, error) {
	if len(val) < 3 {
		return nil, errors.New("firstName must be at least 3 characters long")
	}

	return &firstName{value: val}, nil
}
