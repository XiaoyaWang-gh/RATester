package proxy

import (
	"fmt"
	"testing"
)

func TestresetReqMethod_1(t *testing.T) {
	tests := []struct {
		name   string
		method string
		want   string
	}{
		{"Test Case 1", "GET", "GET"},
		{"Test Case 2", "POST", "POST"},
		{"Test Case 3", "ET", "GET"},
		{"Test Case 4", "OST", "POST"},
		{"Test Case 5", "PUT", "PUT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := resetReqMethod(tt.method); got != tt.want {
				t.Errorf("resetReqMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
