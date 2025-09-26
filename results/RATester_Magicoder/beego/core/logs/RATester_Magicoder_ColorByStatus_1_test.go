package logs

import (
	"fmt"
	"testing"
)

func TestColorByStatus_1(t *testing.T) {
	initColor()

	tests := []struct {
		name string
		code int
		want string
	}{
		{"200", 200, "green"},
		{"299", 299, "green"},
		{"300", 300, "white"},
		{"399", 399, "white"},
		{"400", 400, "yellow"},
		{"499", 499, "yellow"},
		{"500", 500, "red"},
		{"599", 599, "red"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ColorByStatus(tt.code); got != tt.want {
				t.Errorf("ColorByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
