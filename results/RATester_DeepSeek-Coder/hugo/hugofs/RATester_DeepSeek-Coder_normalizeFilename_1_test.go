package hugofs

import (
	"fmt"
	"runtime"
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestNormalizeFilename_1(t *testing.T) {
	t.Run("Test with empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := normalizeFilename("")
		if result != "" {
			t.Errorf("Expected '', got '%s'", result)
		}
	})

	t.Run("Test with non-empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := normalizeFilename("test.txt")
		if runtime.GOOS == "darwin" && result != norm.NFC.String("test.txt") {
			t.Errorf("Expected '%s', got '%s'", norm.NFC.String("test.txt"), result)
		} else if runtime.GOOS != "darwin" && result != "test.txt" {
			t.Errorf("Expected 'test.txt', got '%s'", result)
		}
	})
}
