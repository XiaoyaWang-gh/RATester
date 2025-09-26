package encoding

import (
	"fmt"
	"testing"
)

func TestBase64Encode_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "SGVsbG8gd29ybGQ="
		result, err := ns.Base64Encode("Hello world")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Base64Encode(123)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
