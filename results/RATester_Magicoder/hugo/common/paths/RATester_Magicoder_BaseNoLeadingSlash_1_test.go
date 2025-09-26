package paths

import (
	"fmt"
	"testing"
)

func TestBaseNoLeadingSlash_1(t *testing.T) {
	type args struct {
		p *Path
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := tt.args.p.BaseNoLeadingSlash(); got != tt.want {
				t.Errorf("BaseNoLeadingSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
