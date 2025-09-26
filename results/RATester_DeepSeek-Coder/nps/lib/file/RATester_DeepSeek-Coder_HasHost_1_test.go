package file

import (
	"fmt"
	"testing"
)

func TestHasHost_1(t *testing.T) {
	type args struct {
		h *Host
	}
	tests := []struct {
		name string
		s    *Client
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.HasHost(tt.args.h); got != tt.want {
				t.Errorf("Client.HasHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
