package encoding

import (
	"fmt"
	"testing"
)

func TestBase64Decode_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		decoded, err := ns.Base64Decode("QmFzZTY0")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if decoded != "Base64" {
			t.Errorf("expected 'Base64', got %v", decoded)
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Base64Decode(1234)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
