package utils

import (
	"fmt"
	"testing"
)

func TestSet_31(t *testing.T) {
	tests := []struct {
		name string
		f    StrTo
		v    string
		want StrTo
	}{
		{
			name: "Test with non-empty string",
			f:    "test",
			v:    "test1",
			want: "test1",
		},
		{
			name: "Test with empty string",
			f:    "test",
			v:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.f.Set(tt.v)
			if tt.f != tt.want {
				t.Errorf("got %v, want %v", tt.f, tt.want)
			}
		})
	}
}
