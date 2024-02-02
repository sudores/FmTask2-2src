package app

import (
	"errors"
	"os"
)

func Testable() error {
	failTest := os.Getenv("FAIL_TEST")
	if failTest == "1" {
		return errors.New("Unexpected error")
	}
	return nil
}
