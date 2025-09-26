package web

import (
	"fmt"
	"testing"
)

func TestIsStaticCompress_1(t *testing.T) {
	BConfig.WebConfig.StaticExtensionsToGzip = []string{".jpg", ".png"}

	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Test Case 1", "test.jpg", true},
		{"Test Case 2", "test.png", true},
		{"Test Case 3", "test.txt", false},
		{"Test Case 4", "test.JPG", true},
		{"Test Case 5", "test.PNG", true},
		{"Test Case 6", "test.JPEG", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isStaticCompress(tt.args); got != tt.want {
				t.Errorf("isStaticCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
