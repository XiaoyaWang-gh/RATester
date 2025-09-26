package forwardedheaders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValues_1(t *testing.T) {
	h := make(unsafeHeader)
	h.Set("Content-Type", "application/json")
	h.Set("Content-Length", "123")

	tests := []struct {
		name  string
		key   string
		value []string
	}{
		{"TestContentType", "Content-Type", []string{"application/json"}},
		{"TestContentLength", "Content-Length", []string{"123"}},
		{"TestEmpty", "Empty", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := h.Values(tt.key)
			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("Values() = %v, want %v", got, tt.value)
			}
		})
	}
}
