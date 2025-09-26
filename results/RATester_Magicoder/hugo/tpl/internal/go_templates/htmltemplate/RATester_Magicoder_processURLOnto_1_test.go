package template

import (
	"fmt"
	"strings"
	"testing"
)

func TestprocessURLOnto_1(t *testing.T) {
	type args struct {
		s    string
		norm bool
		b    *strings.Builder
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

			if got := processURLOnto(tt.args.s, tt.args.norm, tt.args.b); got != tt.want {
				t.Errorf("processURLOnto() = %v, want %v", got, tt.want)
			}
		})
	}
}
