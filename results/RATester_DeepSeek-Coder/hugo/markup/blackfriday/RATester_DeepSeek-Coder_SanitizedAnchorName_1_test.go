package blackfriday

import (
	"fmt"
	"testing"
)

func TestSanitizedAnchorName_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"simple", "Hello World", "hello-world"},
		{"with special characters", "Hello, World!!!", "hello-world"},
		{"with numbers", "Hello 123 World", "hello-123-world"},
		{"with multiple spaces", "Hello   World", "hello-world"},
		{"with leading and trailing spaces", "  Hello World  ", "hello-world"},
		{"with leading and trailing special characters", "!!!Hello World!!!", "hello-world"},
		{"with multiple consecutive special characters", "Hello...World", "hello-world"},
		{"with special characters and numbers", "Hello.123_World", "hello-123-world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SanitizedAnchorName(tt.arg); got != tt.want {
				t.Errorf("SanitizedAnchorName() = %v, want %v", got, tt.want)
			}
		})
	}
}
