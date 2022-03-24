package functions

import "errors"

// ValidateToken checks if a token is valid or not
// it returns an error, if the given token is invalid
func ValidateToken(token string) error {
	if len(token) != 32 {
		return errors.New("please enter a valid token")
	}

	return nil
}
