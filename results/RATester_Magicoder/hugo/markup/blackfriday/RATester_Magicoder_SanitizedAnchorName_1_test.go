package blackfriday

import (
	"fmt"
	"testing"
)

func TestSanitizedAnchorName_1(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"simple", "Hello World", "hello-world"},
		{"with special characters", "Hello World!", "hello-world"},
		{"with numbers", "Hello 123 World", "hello-123-world"},
		{"with multiple spaces", "Hello   World", "hello-world"},
		{"with leading and trailing spaces", "  Hello World  ", "hello-world"},
		{"with leading and trailing special characters", "!Hello World!", "hello-world"},
		{"with leading and trailing numbers", "123Hello 123 World123", "hello-123-world"},
		{"with leading and trailing special characters and numbers", "!123Hello 123 World123!", "hello-123-world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SanitizedAnchorName(tt.text); got != tt.want {
				t.Errorf("SanitizedAnchorName() = %v, want %v", got, tt.want)
			}
		})
	}
}
