package app_test

import (
	app "go-server/app_test"
	"testing"
)

func TestTestable(t *testing.T) {
	t.Run("No failure", func(t *testing.T) {
		err := app.Testable()
		if err != nil {
			t.Error("Test failure at Testable")
		}
	})
}
