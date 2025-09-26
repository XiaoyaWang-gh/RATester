package provider

import (
	"fmt"
	"testing"
)

func TestNormalize_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"", ""},
		{"Hello World", "Hello-World"},
		{"hello world", "hello-world"},
		{"Hello_World", "Hello-World"},
		{"Hello.World", "Hello-World"},
		{"Hello World 123", "Hello-World-123"},
		{"Hello   World", "Hello-World"},
		{"Hello  World 123", "Hello-World-123"},
		{"Hello_World_123", "Hello-World-123"},
		{"Hello.World.123", "Hello-World-123"},
		{"Hello   World   123", "Hello-World-123"},
		{"Hello_World_123_", "Hello-World-123"},
		{"Hello.World.123.", "Hello-World-123"},
		{"Hello   World   123  ", "Hello-World-123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Normalize(tt.name); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
