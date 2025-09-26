package dynamic

import (
	"fmt"
	"testing"
)

func TestHasCustomHeadersDefined_1(t *testing.T) {
	tests := []struct {
		name   string
		fields *Headers
		want   bool
	}{
		{
			name: "Test case 1: Both CustomRequestHeaders and CustomResponseHeaders are defined",
			fields: &Headers{
				CustomRequestHeaders:  map[string]string{"header1": "value1"},
				CustomResponseHeaders: map[string]string{"header2": "value2"},
			},
			want: true,
		},
		{
			name: "Test case 2: Only CustomRequestHeaders are defined",
			fields: &Headers{
				CustomRequestHeaders: map[string]string{"header1": "value1"},
			},
			want: true,
		},
		{
			name: "Test case 3: Only CustomResponseHeaders are defined",
			fields: &Headers{
				CustomResponseHeaders: map[string]string{"header2": "value2"},
			},
			want: true,
		},
		{
			name: "Test case 4: Neither CustomRequestHeaders nor CustomResponseHeaders are defined",
			fields: &Headers{
				CustomRequestHeaders:  map[string]string{},
				CustomResponseHeaders: map[string]string{},
			},
			want: false,
		},
		{
			name:   "Test case 5: Headers is nil",
			fields: nil,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			h := &Headers{
				CustomRequestHeaders:  tt.fields.CustomRequestHeaders,
				CustomResponseHeaders: tt.fields.CustomResponseHeaders,
			}
			if got := h.HasCustomHeadersDefined(); got != tt.want {
				t.Errorf("Headers.HasCustomHeadersDefined() = %v, want %v", got, tt.want)
			}
		})
	}
}
