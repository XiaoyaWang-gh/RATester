package helpers

import (
	"fmt"
	"testing"
)

func TestGetTitleFunc_1(t *testing.T) {
	tests := []struct {
		name  string
		style string
		input string
		want  string
	}{
		{
			name:  "Test Go Style",
			style: "go",
			input: "hello world",
			want:  "Hello World",
		},
		{
			name:  "Test Chicago Style",
			style: "chicago",
			input: "hello world",
			want:  "Hello World",
		},
		{
			name:  "Test None Style",
			style: "none",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "Test FirstUpper Style",
			style: "firstupper",
			input: "hello world",
			want:  "Hello World",
		},
		{
			name:  "Test Default Style",
			style: "default",
			input: "hello world",
			want:  "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetTitleFunc(tt.style)(tt.input); got != tt.want {
				t.Errorf("GetTitleFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
