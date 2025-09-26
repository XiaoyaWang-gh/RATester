package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestbodyAllowedForStatus_1(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{"100-199", 100, false},
		{"200-299", 200, true},
		{"300-399", 300, true},
		{"400-499", 400, true},
		{"500-599", 500, true},
		{"600-699", 600, true},
		{"700-799", 700, true},
		{"800-899", 800, true},
		{"900-999", 900, true},
		{"1000", 1000, true},
		{"http.StatusNoContent", http.StatusNoContent, false},
		{"http.StatusNotModified", http.StatusNotModified, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bodyAllowedForStatus(tt.status); got != tt.want {
				t.Errorf("bodyAllowedForStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
