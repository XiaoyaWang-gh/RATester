package web

import (
	"fmt"
	"testing"
)

func TestgetUpperMethodString_1(t *testing.T) {
	cr := &ControllerRegister{}
	tests := []struct {
		name     string
		method   string
		expected string
	}{
		{"Test1", "GET", "GET"},
		{"Test2", "POST", "POST"},
		{"Test3", "PUT", "PUT"},
		{"Test4", "DELETE", "DELETE"},
		{"Test5", "OPTIONS", "OPTIONS"},
		{"Test6", "PATCH", "PATCH"},
		{"Test7", "*", "*"},
		{"Test8", "HEAD", "HEAD"},
		{"Test9", "INVALID", "INVALID"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := cr.getUpperMethodString(test.method)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
