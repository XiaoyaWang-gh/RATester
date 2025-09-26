package hugofs

import (
	"fmt"
	"testing"
)

func TestnormalizeFilename_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "empty string",
			filename: "",
			want:     "",
		},
		{
			name:     "normal string",
			filename: "test.txt",
			want:     "test.txt",
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

			if got := normalizeFilename(tt.filename); got != tt.want {
				t.Errorf("normalizeFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
