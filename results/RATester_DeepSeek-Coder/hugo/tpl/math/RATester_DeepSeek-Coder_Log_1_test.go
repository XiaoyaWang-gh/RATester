package math

import (
	"fmt"
	"testing"
)

func TestLog_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		res, err := ns.Log(10)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if res != 2.302585092994046 {
			t.Errorf("expected 2.302585092994046, got %v", res)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Log("a")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
