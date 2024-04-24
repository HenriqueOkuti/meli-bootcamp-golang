package time_limit_test

import (
	"gotests03tarde/utils/time_limit"
	"os"
	"testing"
	"time"
)

func TestGetTimeLimit(t *testing.T) {
	t.Run("when TIMEOUT_FIBONACCI is not set", func(t *testing.T) {
		os.Unsetenv("TIMEOUT_FIBONACCI")

		expected := 10 * time.Second
		got := time_limit.GetTimeLimit()

		if got != expected {
			t.Errorf("GetTimeLimit() = %v; want %v", got, expected)
		}
	})

	t.Run("when TIMEOUT_FIBONACCI is set", func(t *testing.T) {
		os.Setenv("TIMEOUT_FIBONACCI", "5s")
		defer os.Unsetenv("TIMEOUT_FIBONACCI")

		expected := 5 * time.Second
		got := time_limit.GetTimeLimit()

		if got != expected {
			t.Errorf("GetTimeLimit() = %v; want %v", got, expected)
		}
	})
}
