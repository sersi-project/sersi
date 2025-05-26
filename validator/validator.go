package validator

import (
	"fmt"
	"regexp"
)

func ValidateString(s string) error {
	if len(s) < 3 {
		return fmt.Errorf("name is too short: %s", s)
	}
	matched, err := regexp.MatchString("^[a-zA-Z0-9_-]+$", s)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("name is invalid: %s", s)
	}
	return nil
}
