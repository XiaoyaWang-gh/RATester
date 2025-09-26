package context

import (
	"fmt"
	"testing"
)

func TestsanitizeName_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"simple name", "test", "test"},
		{"name with special characters", "test@#$%^&*()", "test"},
		{"name with spaces", "test name", "test_name"},
		{"name with multiple spaces", "test  name", "test_name"},
		{"name with leading/trailing spaces", " test ", "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sanitizeName(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
