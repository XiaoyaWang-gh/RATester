package logs

import (
	"fmt"
	"testing"
)

func TestColorByMethod_1(t *testing.T) {
	tests := []struct {
		name   string
		method string
		want   string
	}{
		{"Test1", "GET", "green"},
		{"Test2", "POST", "blue"},
		{"Test3", "PUT", "yellow"},
		{"Test4", "DELETE", "red"},
		{"Test5", "HEAD", "cyan"},
		{"Test6", "OPTIONS", "magenta"},
		{"Test7", "PATCH", "white"},
		{"Test8", "TRACE", "reset"},
		{"Test9", "CONNECT", "reset"},
		{"Test10", "INVALID", "reset"},
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
