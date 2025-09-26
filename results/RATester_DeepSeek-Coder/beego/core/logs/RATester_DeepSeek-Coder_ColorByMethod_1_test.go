package logs

import (
	"fmt"
	"testing"
)

func TestColorByMethod_1(t *testing.T) {
	initColor()

	tests := []struct {
		name   string
		method string
		want   string
	}{
		{"GET", "GET", green},
		{"POST", "POST", blue},
		{"PUT", "PUT", yellow},
		{"DELETE", "DELETE", red},
		{"PATCH", "PATCH", cyan},
		{"HEAD", "HEAD", reset},
		{"OPTIONS", "OPTIONS", reset},
		{"UNKNOWN", "UNKNOWN", reset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ColorByMethod(tt.method); got != tt.want {
				t.Errorf("ColorByMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
