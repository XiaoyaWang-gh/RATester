package commands

import (
	"fmt"
	"testing"
)

func TestHugo_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"success", false},
		{"error", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hugoBuilder{
				r: &rootCommand{},
				s: &serverCommand{},
			}

			got, err := c.hugo()
			if (err != nil) != tt.wantErr {
				t.Errorf("hugo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (got != nil) != tt.wantErr {
				t.Errorf("hugo() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
