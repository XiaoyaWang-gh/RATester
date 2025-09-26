package logs

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetLogger_1(t *testing.T) {
	t.Run("Test with no prefix", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		logger := GetLogger()
		if logger == nil {
			t.Errorf("Expected logger to be not nil")
		}
	})

	t.Run("Test with prefix", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		prefix := "test"
		logger := GetLogger(prefix)
		if logger == nil {
			t.Errorf("Expected logger to be not nil")
		}
		if logger.Prefix() != fmt.Sprintf(`[%s] `, strings.ToUpper(prefix)) {
			t.Errorf("Expected prefix to be %s, got %s", fmt.Sprintf(`[%s] `, strings.ToUpper(prefix)), logger.Prefix())
		}
	})

	t.Run("Test with multiple prefixes", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		prefixes := []string{"test1", "test2"}
		logger := GetLogger(prefixes...)
		if logger == nil {
			t.Errorf("Expected logger to be not nil")
		}
		if logger.Prefix() != fmt.Sprintf(`[%s] `, strings.ToUpper(prefixes[0])) {
			t.Errorf("Expected prefix to be %s, got %s", fmt.Sprintf(`[%s] `, strings.ToUpper(prefixes[0])), logger.Prefix())
		}
	})
}
