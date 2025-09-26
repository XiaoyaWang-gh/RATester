package web

import (
	"fmt"
	"testing"
)

func TestURLFor_3(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		values   []interface{}
		want     string
	}{
		{
			name:     "Test Case 1",
			endpoint: "/test",
			values:   []interface{}{"field name"},
			want:     "/test/field name",
		},
		{
			name:     "Test Case 2",
			endpoint: "/test",
			values:   []interface{}{"field name", "field name 2"},
			want:     "/test/field name/field name 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := URLFor(tt.endpoint, tt.values...); got != tt.want {
				t.Errorf("URLFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
