package functions

import "errors"

// ValidateToken checks if a token is valid or not
// it returns an error, if the given token is invalid
func ValidateToken(token string) error {
	if token == "" || len(token) != 32 {
		return errors.New("Please enter a valid token!")
	}

	return nil
}
