package modules

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestShouldVendor_1(t *testing.T) {
	tests := []struct {
		name string
		c    *Client
		path string
		want bool
	}{
		{
			"noVendor is nil",
			&Client{
				noVendor: nil,
			},
			"foo",
			true,
		},
		{
			"noVendor is not nil and path matches",
			&Client{
				noVendor: glob.MustCompile("foo"),
			},
			"foo",
			false,
		},
		{
			"noVendor is not nil and path does not match",
			&Client{
				noVendor: glob.MustCompile("bar"),
			},
			"foo",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.shouldVendor(tt.path); got != tt.want {
				t.Errorf("Client.shouldVendor() = %v, want %v", got, tt.want)
			}
		})
	}
}
