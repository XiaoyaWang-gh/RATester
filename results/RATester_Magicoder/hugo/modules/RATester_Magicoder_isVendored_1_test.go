package modules

import (
	"fmt"
	"testing"
)

func TestisVendored_1(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
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

			c := &collector{}
			if got := c.isVendored(tt.args.dir); got != tt.want {
				t.Errorf("collector.isVendored() = %v, want %v", got, tt.want)
			}
		})
	}
}
