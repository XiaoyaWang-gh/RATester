package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestIs_1(t *testing.T) {
	t.Run("returns true when the target is a *FeatureNotAvailableError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := &FeatureNotAvailableError{}
		if !err.Is(err) {
			t.Errorf("expected Is to return true")
		}
	})

	t.Run("returns false when the target is not a *FeatureNotAvailableError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := &FeatureNotAvailableError{}
		otherErr := errors.New("other error")
		if err.Is(otherErr) {
			t.Errorf("expected Is to return false")
		}
	})
}
