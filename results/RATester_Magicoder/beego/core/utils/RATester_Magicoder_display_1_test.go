package utils

import (
	"fmt"
	"testing"
)

func Testdisplay_1(t *testing.T) {
	tests := []struct {
		name      string
		displayed bool
		data      []interface{}
		want      string
	}{
		{
			name:      "Test Case 1",
			displayed: true,
			data:      []interface{}{"field name", "value"},
			want:      "[Debug] at Testdisplay() [/path/to/file.go:123]\n\n[Variables]\nfield name = value\n",
		},
		{
			name:      "Test Case 2",
			displayed: false,
			data:      []interface{}{"field want", "value"},
			want:      "[Debug] at Testdisplay() [/path/to/file.go:123]\n\n[Variables]\nfield want = value\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := display(tt.displayed, tt.data...); got != tt.want {
				t.Errorf("display() = %v, want %v", got, tt.want)
			}
		})
	}
}
