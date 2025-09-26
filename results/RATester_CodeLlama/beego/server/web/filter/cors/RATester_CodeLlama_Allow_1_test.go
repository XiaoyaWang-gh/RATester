package cors

import (
	"fmt"
	"testing"
)

func TestAllow_1(t *testing.T) {
	type args struct {
		opts *Options
	}
	tests := []struct {
		name string
		args args
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

			Allow(tt.args.opts)
		})
	}
}
