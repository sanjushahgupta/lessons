package domain

import (
	"fmt"
)

var (
	// ErrEmailExist returns an error if email exists in database
	ErrEmailExist = func(email string) error {
		return fmt.Errorf("email exists: %s", email)
	}
)
