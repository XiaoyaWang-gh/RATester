package rst

import (
	"fmt"
	"testing"
)

func TestgetRstBinaryNameAndPath_1(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 string
	}{
		{
			name:  "Test case 1",
			want:  "rst2pdf",
			want1: "/usr/bin/rst2pdf",
		},
		{
			name:  "Test case 2",
			want:  "rst2html",
			want1: "/usr/bin/rst2html",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := getRstBinaryNameAndPath()
			if got != tt.want {
				t.Errorf("getRstBinaryNameAndPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRstBinaryNameAndPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
