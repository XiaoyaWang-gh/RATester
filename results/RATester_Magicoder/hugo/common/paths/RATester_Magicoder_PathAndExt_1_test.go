package paths

import (
	"fmt"
	"testing"
)

func TestPathAndExt_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		want1 string
	}{
		{
			name:  "Test case 1",
			input: "/path/to/file.txt",
			want:  "/path/to/file",
			want1: ".txt",
		},
		{
			name:  "Test case 2",
			input: "/path/to/file",
			want:  "/path/to/file",
			want1: "",
		},
		{
			name:  "Test case 3",
			input: "/path/to/file.with.multiple.dots.txt",
			want:  "/path/to/file.with.multiple.dots",
			want1: ".txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := PathAndExt(tt.input)
			if got != tt.want {
				t.Errorf("PathAndExt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PathAndExt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
