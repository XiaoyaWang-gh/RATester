package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	h := make(unsafeHeader)
	h.Set("Content-Type", "application/json")

	tests := []struct {
		name string
		key  string
		want string
	}{
		{"TestGet", "Content-Type", "application/json"},
		{"TestGetEmpty", "Content-Length", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := h.Get(tt.key); got != tt.want {
				t.Errorf("unsafeHeader.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
